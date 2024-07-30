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

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	svcsdk "github.com/aws/aws-sdk-go/service/acmpca"
)

func (rm *resourceManager) writeCertificateChainToSecret(
	ctx context.Context,
	certificate string,
	certificateChain string,
	resourceNamespace string,
	secretKeyReference *ackv1alpha1.SecretKeyReference,
) (err error) {

	namespace := resourceNamespace
	if secretKeyReference.SecretReference.Namespace != "" {
		namespace = secretKeyReference.SecretReference.Namespace
	}

	completeCertificateChain := certificate
	if certificateChain != "" {
		completeCertificateChain = certificate + "\n" + certificateChain
	}

	err = rm.rr.WriteToSecret(ctx, completeCertificateChain, namespace, secretKeyReference.SecretReference.Name, secretKeyReference.Key)
	rm.metrics.RecordAPICall("PATCH", "writeCertificateChainToSecret", err)
	if err != nil {
		return err
	}

	return nil
}

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

	if ko.Spec.Status == nil || *ko.Spec.Status == svcsdk.CertificateAuthorityStatusCreating || *ko.Spec.Status == svcsdk.CertificateAuthorityStatusPendingCertificate {
		return nil, ackerr.NotFound
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

	if desired.ko.Spec.Status != nil {
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
