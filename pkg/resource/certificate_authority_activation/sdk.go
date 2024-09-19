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

// Code generated by ack-generate. DO NOT EDIT.

package certificate_authority_activation

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/acmpca"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/acmpca-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ACMPCA{}
	_ = &svcapitypes.CertificateAuthorityActivation{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	return rm.customFindCertificateAuthorityActivation(ctx, r)
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}
	certificateSecret := ""
	certificateChainSecret := ""
	if desired.ko.Spec.Certificate != nil {
		certificateSecret, err = rm.rr.SecretValueFromReference(ctx, desired.ko.Spec.Certificate)
		if err != nil {
			return nil, ackrequeue.Needed(err)
		}
		if certificateSecret != "" {
			input.SetCertificate([]byte(certificateSecret))
		}
	}
	if desired.ko.Spec.CertificateChain != nil {
		certificateChainSecret, err = rm.rr.SecretValueFromReference(ctx, desired.ko.Spec.CertificateChain)
		if err != nil {
			return nil, ackrequeue.Needed(err)
		}
		if certificateChainSecret != "" {
			input.SetCertificateChain([]byte(certificateChainSecret))
		}
	}

	var resp *svcsdk.ImportCertificateAuthorityCertificateOutput
	_ = resp
	resp, err = rm.sdkapi.ImportCertificateAuthorityCertificateWithContext(ctx, input)
	if err != nil {
		input := &svcsdk.DescribeCertificateAuthorityInput{}
		input.CertificateAuthorityArn = desired.ko.Spec.CertificateAuthorityARN

		var describeResp *svcsdk.DescribeCertificateAuthorityOutput
		describeResp, describeErr := rm.sdkapi.DescribeCertificateAuthorityWithContext(ctx, input)
		rm.metrics.RecordAPICall("READ_ONE", "DescribeCertificateAuthority", err)
		if describeErr != nil {
			return desired, ackrequeue.NeededAfter(describeErr, ackrequeue.DefaultRequeueAfterDuration)
		}

		if *describeResp.CertificateAuthority.Status != svcsdk.CertificateAuthorityStatusFailed && *describeResp.CertificateAuthority.Status != svcsdk.CertificateAuthorityStatusDeleted && *describeResp.CertificateAuthority.Status != svcsdk.CertificateAuthorityStatusDisabled {
			return desired, ackrequeue.NeededAfter(err, ackrequeue.DefaultRequeueAfterDuration)
		}

		return nil, err
	}
	rm.metrics.RecordAPICall("CREATE", "ImportCertificateAuthorityCertificate", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	rm.setStatusDefaults(ko)
	if certificateSecret != "" {
		err = rm.writeCertificateChainToSecret(ctx, certificateSecret, certificateChainSecret, desired.ko.GetNamespace(), desired.ko.Spec.CompleteCertificateChainOutput)
		if err != nil {
			return nil, err
		}
	}

	if desired.ko.Spec.Status != nil && *desired.ko.Spec.Status == svcsdk.CertificateAuthorityStatusDisabled {
		updateInput := &svcsdk.UpdateCertificateAuthorityInput{}

		updateInput.SetStatus(*desired.ko.Spec.Status)

		if desired.ko.Spec.CertificateAuthorityARN != nil {
			updateInput.SetCertificateAuthorityArn(*desired.ko.Spec.CertificateAuthorityARN)
		}

		_, err = rm.sdkapi.UpdateCertificateAuthorityWithContext(ctx, updateInput)
		rm.metrics.RecordAPICall("UPDATE", "UpdateCertificateAuthority", err)
		if err != nil {
			return nil, err
		}
	}
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.ImportCertificateAuthorityCertificateInput, error) {
	res := &svcsdk.ImportCertificateAuthorityCertificateInput{}

	if r.ko.Spec.CertificateAuthorityARN != nil {
		res.SetCertificateAuthorityArn(*r.ko.Spec.CertificateAuthorityARN)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	return rm.customUpdateCertificateAuthorityActivation(ctx, desired, latest, delta)
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	return rm.customDeleteCertificateAuthorityActivation(ctx, r)
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.CertificateAuthorityActivation,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "InvalidAction",
		"InvalidParameterCombination",
		"InvalidParameterValue",
		"InvalidQueryParameter",
		"MissingParameter",
		"ValidationError",
		"ValidationException",
		"CertificateMismatchException",
		"InvalidArnException",
		"InvalidRequestException",
		"InvalidStateException",
		"MalformedCertificateException",
		"RequestFailedException":
		return true
	default:
		return false
	}
}

// getImmutableFieldChanges returns list of immutable fields from the
func (rm *resourceManager) getImmutableFieldChanges(
	delta *ackcompare.Delta,
) []string {
	var fields []string
	if delta.DifferentAt("Spec.Certificate") {
		fields = append(fields, "Certificate")
	}
	if delta.DifferentAt("Spec.CertificateAuthorityARN") {
		fields = append(fields, "CertificateAuthorityARN")
	}
	if delta.DifferentAt("Spec.CertificateChain") {
		fields = append(fields, "CertificateChain")
	}
	if delta.DifferentAt("Spec.CompleteCertificateChainOutput") {
		fields = append(fields, "CompleteCertificateChainOutput")
	}

	return fields
}
