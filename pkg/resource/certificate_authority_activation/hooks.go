// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package certificate_authority_activation

import (
	"context"
	"fmt"
	"sync"

	client "github.com/aws-controllers-k8s/acmpca-controller/pkg/client"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	svcsdk "github.com/aws/aws-sdk-go/service/acmpca"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	mu sync.Mutex
)

func (rm *resourceManager) customFindCertificateAuthorityActivation(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.customFindCertificateAuthorityActivation")
	defer func() {
		exit(err)
	}()

	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if r.ko.Spec.CertificateAuthorityARN == nil {
		return nil, ackerr.NotFound
	}

	// lock runtime
	mu.Lock()
	defer mu.Unlock()

	// List all the CertificateAuthorityActivations
	dynClient, err := client.GetDynamicClient()
	if err != nil {
		return nil, err
	}

	var caActivationResource = schema.GroupVersionResource{Group: "acmpca.services.k8s.aws", Version: "v1alpha1", Resource: "certificateauthorityactivations"}
	list, err := dynClient.Resource(caActivationResource).Namespace("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	numFound := 0

	for _, item := range list.Items {

		certificateAuthorityARN, found, err := unstructured.NestedString(item.UnstructuredContent(), "spec", "certificateAuthorityARN")
		if err != nil {
			return nil, err
		}

		if !found {
			return nil, fmt.Errorf("certificateAuthorityARN field not found on CertificateAuthorityActivation spec")
		}

		if certificateAuthorityARN == *r.ko.Spec.CertificateAuthorityARN {
			numFound++
			if numFound > 1 {
				status, found, err := unstructured.NestedString(item.Object, "status", "status")
				if err != nil {
					return nil, err
				}

				if !found {
					return nil, fmt.Errorf("status field not found on CertificateAuthorityActivation status")
				}

				if status == svcsdk.CertificateAuthorityStatusActive {
					return nil, ackerr.Terminal
				}
			}
		}
	}

	input := &svcsdk.DescribeCertificateAuthorityInput{}
	input.CertificateAuthorityArn = r.ko.Spec.CertificateAuthorityARN

	var resp *svcsdk.DescribeCertificateAuthorityOutput
	resp, err = rm.sdkapi.DescribeCertificateAuthorityWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeCertificateAuthority", err)
	if err != nil {
		return nil, err
	}

	r.ko.Status.Status = resp.CertificateAuthority.Status

	if numFound == 1 {
		if *r.ko.Status.Status == svcsdk.CertificateAuthorityStatusCreating || *r.ko.Status.Status == svcsdk.CertificateAuthorityStatusPendingCertificate {
			return nil, ackerr.NotFound
		}
	}

	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}
