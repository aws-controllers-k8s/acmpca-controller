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
	"encoding/json"
	"fmt"
	"sync"

	client "github.com/aws-controllers-k8s/acmpca-controller/pkg/client"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	svcsdk "github.com/aws/aws-sdk-go/service/acmpca"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
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

	if r.ko.Spec.CertificateAuthorityARN == nil && r.ko.Spec.CertificateAuthorityRef.From.Name == nil {
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

	if r.ko.Spec.CertificateAuthorityARN == nil && r.ko.Spec.CertificateAuthorityRef != nil {
		var caResource = schema.GroupVersionResource{Group: "acmpca.services.k8s.aws", Version: "v1alpha1", Resource: "certificateauthorities"}
		ca, err := dynClient.Resource(caResource).Namespace(r.MetaObject().GetNamespace()).Get(ctx, *r.ko.Spec.CertificateAuthorityRef.From.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		certificateAuthorityARN, found, err := unstructured.NestedString(ca.UnstructuredContent(), "status", "ackResourceMetadata", "arn")
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, fmt.Errorf("arn field not found on CertificateAuthority status")
		}
		r.ko.Spec.CertificateAuthorityARN = &certificateAuthorityARN
	}

	var caActivationResource = schema.GroupVersionResource{Group: "acmpca.services.k8s.aws", Version: "v1alpha1", Resource: "certificateauthorityactivations"}
	list, err := dynClient.Resource(caActivationResource).Namespace(r.MetaObject().GetNamespace()).List(ctx, metav1.ListOptions{})
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
			certificateAuthorityRef, found, err := unstructured.NestedString(item.UnstructuredContent(), "spec", "certificateAuthorityRef", "from", "name")
			if err != nil {
				return nil, err
			}
			if !found {
				return nil, fmt.Errorf("certificateAuthorityARN or certificateAuthorityRef field not found on CertificateAuthorityActivation spec")
			}
			var caResource = schema.GroupVersionResource{Group: "acmpca.services.k8s.aws", Version: "v1alpha1", Resource: "certificateauthorities"}
			ca, err := dynClient.Resource(caResource).Namespace(r.MetaObject().GetNamespace()).Get(ctx, certificateAuthorityRef, metav1.GetOptions{})
			if err != nil {
				return nil, err
			}
			certificateAuthorityARN, found, err = unstructured.NestedString(ca.UnstructuredContent(), "status", "ackResourceMetadata", "arn")
			if err != nil {
				return nil, err
			}
			if !found {
				return nil, fmt.Errorf("arn field not found on CertificateAuthority status")
			}
		}

		if certificateAuthorityARN == *r.ko.Spec.CertificateAuthorityARN {
			numFound++
			if numFound > 1 {
				status, found, err := unstructured.NestedString(item.Object, "spec", "status")
				if err != nil {
					return nil, err
				}

				if !found {
					return nil, fmt.Errorf("status field not found on CertificateAuthorityActivation spec")
				}

				if status == svcsdk.CertificateAuthorityStatusActive {
					return nil, fmt.Errorf("CertificateAuthorityActivation resource already exists for this CertificateAuthority")
				}
			}
		}
	}

	if numFound == 0 {
		return nil, ackerr.NotFound
	}

	input := &svcsdk.DescribeCertificateAuthorityInput{}
	input.CertificateAuthorityArn = r.ko.Spec.CertificateAuthorityARN

	var resp *svcsdk.DescribeCertificateAuthorityOutput
	resp, err = rm.sdkapi.DescribeCertificateAuthorityWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeCertificateAuthority", err)
	if err != nil {
		return nil, err
	}

	ko := r.ko.DeepCopy()

	if resp.CertificateAuthority.Status != nil {
		ko.Spec.Status = resp.CertificateAuthority.Status
	} else {
		ko.Spec.Status = nil
	}

	if numFound == 1 {
		if ko.Spec.Status == nil || *ko.Spec.Status == svcsdk.CertificateAuthorityStatusCreating || *ko.Spec.Status == svcsdk.CertificateAuthorityStatusPendingCertificate {
			return nil, ackerr.NotFound
		}
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

func (rm *resourceManager) customUpdateCertificateAuthorityActivation(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()

	input := &svcsdk.UpdateCertificateAuthorityInput{}

	if desired.ko.Spec.CertificateAuthorityARN != nil {
		input.SetCertificateAuthorityArn(*desired.ko.Spec.CertificateAuthorityARN)
	}

	if desired.ko.Spec.Status != nil && (*desired.ko.Spec.Status == svcsdk.CertificateAuthorityStatusActive || *desired.ko.Spec.Status == svcsdk.CertificateAuthorityStatusDisabled) {
		input.SetStatus(*desired.ko.Spec.Status)
	}

	var resp *svcsdk.UpdateCertificateAuthorityOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateCertificateAuthorityWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateCertificateAuthority", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

func (rm *resourceManager) writeCertificateChainToSecret(
	ctx context.Context,
	certificate string,
	certificateChain string,
	objectMeta metav1.ObjectMeta,
) (err error) {

	annotations := objectMeta.GetAnnotations()

	namespace, found := annotations["acmpca.services.k8s.aws/chain-secret-namespace"]
	if !found {
		namespace = objectMeta.GetNamespace()
	}

	name, found := annotations["acmpca.services.k8s.aws/chain-secret-name"]
	if !found {
		return ackerr.SecretNotFound
	}

	key, found := annotations["acmpca.services.k8s.aws/chain-secret-key"]
	if !found {
		key = "certificateChain"
	}

	completeCertificateChain := certificate

	if certificateChain != "" {
		completeCertificateChain = certificate + "\n" + certificateChain
	}

	secretsClient, err := client.GetSecretsClient(namespace)
	if err != nil {
		return err
	}

	secret := corev1.Secret{
		Data: map[string][]byte{
			key: []byte(completeCertificateChain),
		},
	}

	payloadBytes, err := json.Marshal(secret)
	if err != nil {
		return err
	}

	_, err = secretsClient.Patch(ctx, name, types.StrategicMergePatchType, payloadBytes, metav1.PatchOptions{})
	rm.metrics.RecordAPICall("PATCH", "writeCertificateChainToSecret", err)
	if err != nil {
		return err
	}

	return nil
}

func (rm *resourceManager) customDeleteCertificateAuthorityActivation(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {

	input := &svcsdk.DescribeCertificateAuthorityInput{}
	input.CertificateAuthorityArn = r.ko.Spec.CertificateAuthorityARN

	var resp *svcsdk.DescribeCertificateAuthorityOutput
	resp, err = rm.sdkapi.DescribeCertificateAuthorityWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeCertificateAuthority", err)
	if err != nil {
		return nil, err
	}

	if resp.CertificateAuthority.Status != nil && *resp.CertificateAuthority.Status != svcsdk.CertificateAuthorityStatusDeleted && *resp.CertificateAuthority.Status != svcsdk.CertificateAuthorityStatusDisabled {
		input := &svcsdk.UpdateCertificateAuthorityInput{}

		if r.ko.Spec.CertificateAuthorityARN != nil {
			input.SetCertificateAuthorityArn(*r.ko.Spec.CertificateAuthorityARN)
		}

		input.SetStatus(svcsdk.CertificateAuthorityStatusDisabled)

		var resp *svcsdk.UpdateCertificateAuthorityOutput
		_ = resp
		resp, err = rm.sdkapi.UpdateCertificateAuthorityWithContext(ctx, input)
		rm.metrics.RecordAPICall("UPDATE", "UpdateCertificateAuthority", err)
		if err != nil {
			return nil, err
		}
	}

	return r, nil
}
