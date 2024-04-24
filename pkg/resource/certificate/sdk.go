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

package certificate

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
	_ = &svcapitypes.Certificate{}
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
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.GetCertificateOutput
	resp, err = rm.sdkapi.GetCertificateWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetCertificate", err)
	if err != nil {
		if reqErr, ok := ackerr.AWSRequestFailure(err); ok && reqErr.StatusCode() == 404 {
			return nil, ackerr.NotFound
		}
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	rm.setStatusDefaults(ko)
	err = rm.writeCertificateToSecret(ctx, *resp.Certificate, r.ko.ObjectMeta)
	if err != nil && strings.HasPrefix(err.Error(), "RequestInProgressException") {
		return &resource{ko}, ackrequeue.NeededAfter(err, ackrequeue.DefaultRequeueAfterDuration)
	}
	if err != nil {
		return nil, err
	}
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.CertificateAuthorityARN == nil || (r.ko.Status.ACKResourceMetadata == nil || r.ko.Status.ACKResourceMetadata.ARN == nil)

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetCertificateInput, error) {
	res := &svcsdk.GetCertificateInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetCertificateArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}
	if r.ko.Spec.CertificateAuthorityARN != nil {
		res.SetCertificateAuthorityArn(*r.ko.Spec.CertificateAuthorityARN)
	}

	return res, nil
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
	if desired.ko.Spec.CertificateSigningRequest != nil {
		input.SetCsr([]byte(*desired.ko.Spec.CertificateSigningRequest))
	}

	var resp *svcsdk.IssueCertificateOutput
	_ = resp
	resp, err = rm.sdkapi.IssueCertificateWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "IssueCertificate", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.CertificateArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.CertificateArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.IssueCertificateInput, error) {
	res := &svcsdk.IssueCertificateInput{}

	if r.ko.Spec.APIPassthrough != nil {
		f0 := &svcsdk.ApiPassthrough{}
		if r.ko.Spec.APIPassthrough.Extensions != nil {
			f0f0 := &svcsdk.Extensions{}
			if r.ko.Spec.APIPassthrough.Extensions.CertificatePolicies != nil {
				f0f0f0 := []*svcsdk.PolicyInformation{}
				for _, f0f0f0iter := range r.ko.Spec.APIPassthrough.Extensions.CertificatePolicies {
					f0f0f0elem := &svcsdk.PolicyInformation{}
					if f0f0f0iter.CertPolicyID != nil {
						f0f0f0elem.SetCertPolicyId(*f0f0f0iter.CertPolicyID)
					}
					if f0f0f0iter.PolicyQualifiers != nil {
						f0f0f0elemf1 := []*svcsdk.PolicyQualifierInfo{}
						for _, f0f0f0elemf1iter := range f0f0f0iter.PolicyQualifiers {
							f0f0f0elemf1elem := &svcsdk.PolicyQualifierInfo{}
							if f0f0f0elemf1iter.PolicyQualifierID != nil {
								f0f0f0elemf1elem.SetPolicyQualifierId(*f0f0f0elemf1iter.PolicyQualifierID)
							}
							if f0f0f0elemf1iter.Qualifier != nil {
								f0f0f0elemf1elemf1 := &svcsdk.Qualifier{}
								if f0f0f0elemf1iter.Qualifier.CPSURI != nil {
									f0f0f0elemf1elemf1.SetCpsUri(*f0f0f0elemf1iter.Qualifier.CPSURI)
								}
								f0f0f0elemf1elem.SetQualifier(f0f0f0elemf1elemf1)
							}
							f0f0f0elemf1 = append(f0f0f0elemf1, f0f0f0elemf1elem)
						}
						f0f0f0elem.SetPolicyQualifiers(f0f0f0elemf1)
					}
					f0f0f0 = append(f0f0f0, f0f0f0elem)
				}
				f0f0.SetCertificatePolicies(f0f0f0)
			}
			if r.ko.Spec.APIPassthrough.Extensions.CustomExtensions != nil {
				f0f0f1 := []*svcsdk.CustomExtension{}
				for _, f0f0f1iter := range r.ko.Spec.APIPassthrough.Extensions.CustomExtensions {
					f0f0f1elem := &svcsdk.CustomExtension{}
					if f0f0f1iter.Critical != nil {
						f0f0f1elem.SetCritical(*f0f0f1iter.Critical)
					}
					if f0f0f1iter.ObjectIdentifier != nil {
						f0f0f1elem.SetObjectIdentifier(*f0f0f1iter.ObjectIdentifier)
					}
					if f0f0f1iter.Value != nil {
						f0f0f1elem.SetValue(*f0f0f1iter.Value)
					}
					f0f0f1 = append(f0f0f1, f0f0f1elem)
				}
				f0f0.SetCustomExtensions(f0f0f1)
			}
			if r.ko.Spec.APIPassthrough.Extensions.ExtendedKeyUsage != nil {
				f0f0f2 := []*svcsdk.ExtendedKeyUsage{}
				for _, f0f0f2iter := range r.ko.Spec.APIPassthrough.Extensions.ExtendedKeyUsage {
					f0f0f2elem := &svcsdk.ExtendedKeyUsage{}
					if f0f0f2iter.ExtendedKeyUsageObjectIdentifier != nil {
						f0f0f2elem.SetExtendedKeyUsageObjectIdentifier(*f0f0f2iter.ExtendedKeyUsageObjectIdentifier)
					}
					if f0f0f2iter.ExtendedKeyUsageType != nil {
						f0f0f2elem.SetExtendedKeyUsageType(*f0f0f2iter.ExtendedKeyUsageType)
					}
					f0f0f2 = append(f0f0f2, f0f0f2elem)
				}
				f0f0.SetExtendedKeyUsage(f0f0f2)
			}
			if r.ko.Spec.APIPassthrough.Extensions.KeyUsage != nil {
				f0f0f3 := &svcsdk.KeyUsage{}
				if r.ko.Spec.APIPassthrough.Extensions.KeyUsage.CRLSign != nil {
					f0f0f3.SetCRLSign(*r.ko.Spec.APIPassthrough.Extensions.KeyUsage.CRLSign)
				}
				if r.ko.Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment != nil {
					f0f0f3.SetDataEncipherment(*r.ko.Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment)
				}
				if r.ko.Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly != nil {
					f0f0f3.SetDecipherOnly(*r.ko.Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly)
				}
				if r.ko.Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature != nil {
					f0f0f3.SetDigitalSignature(*r.ko.Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature)
				}
				if r.ko.Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly != nil {
					f0f0f3.SetEncipherOnly(*r.ko.Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly)
				}
				if r.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement != nil {
					f0f0f3.SetKeyAgreement(*r.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement)
				}
				if r.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign != nil {
					f0f0f3.SetKeyCertSign(*r.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign)
				}
				if r.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment != nil {
					f0f0f3.SetKeyEncipherment(*r.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment)
				}
				if r.ko.Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation != nil {
					f0f0f3.SetNonRepudiation(*r.ko.Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation)
				}
				f0f0.SetKeyUsage(f0f0f3)
			}
			if r.ko.Spec.APIPassthrough.Extensions.SubjectAlternativeNames != nil {
				f0f0f4 := []*svcsdk.GeneralName{}
				for _, f0f0f4iter := range r.ko.Spec.APIPassthrough.Extensions.SubjectAlternativeNames {
					f0f0f4elem := &svcsdk.GeneralName{}
					if f0f0f4iter.DirectoryName != nil {
						f0f0f4elemf0 := &svcsdk.ASN1Subject{}
						if f0f0f4iter.DirectoryName.CommonName != nil {
							f0f0f4elemf0.SetCommonName(*f0f0f4iter.DirectoryName.CommonName)
						}
						if f0f0f4iter.DirectoryName.Country != nil {
							f0f0f4elemf0.SetCountry(*f0f0f4iter.DirectoryName.Country)
						}
						if f0f0f4iter.DirectoryName.CustomAttributes != nil {
							f0f0f4elemf0f2 := []*svcsdk.CustomAttribute{}
							for _, f0f0f4elemf0f2iter := range f0f0f4iter.DirectoryName.CustomAttributes {
								f0f0f4elemf0f2elem := &svcsdk.CustomAttribute{}
								if f0f0f4elemf0f2iter.ObjectIdentifier != nil {
									f0f0f4elemf0f2elem.SetObjectIdentifier(*f0f0f4elemf0f2iter.ObjectIdentifier)
								}
								if f0f0f4elemf0f2iter.Value != nil {
									f0f0f4elemf0f2elem.SetValue(*f0f0f4elemf0f2iter.Value)
								}
								f0f0f4elemf0f2 = append(f0f0f4elemf0f2, f0f0f4elemf0f2elem)
							}
							f0f0f4elemf0.SetCustomAttributes(f0f0f4elemf0f2)
						}
						if f0f0f4iter.DirectoryName.DistinguishedNameQualifier != nil {
							f0f0f4elemf0.SetDistinguishedNameQualifier(*f0f0f4iter.DirectoryName.DistinguishedNameQualifier)
						}
						if f0f0f4iter.DirectoryName.GenerationQualifier != nil {
							f0f0f4elemf0.SetGenerationQualifier(*f0f0f4iter.DirectoryName.GenerationQualifier)
						}
						if f0f0f4iter.DirectoryName.GivenName != nil {
							f0f0f4elemf0.SetGivenName(*f0f0f4iter.DirectoryName.GivenName)
						}
						if f0f0f4iter.DirectoryName.Initials != nil {
							f0f0f4elemf0.SetInitials(*f0f0f4iter.DirectoryName.Initials)
						}
						if f0f0f4iter.DirectoryName.Locality != nil {
							f0f0f4elemf0.SetLocality(*f0f0f4iter.DirectoryName.Locality)
						}
						if f0f0f4iter.DirectoryName.Organization != nil {
							f0f0f4elemf0.SetOrganization(*f0f0f4iter.DirectoryName.Organization)
						}
						if f0f0f4iter.DirectoryName.OrganizationalUnit != nil {
							f0f0f4elemf0.SetOrganizationalUnit(*f0f0f4iter.DirectoryName.OrganizationalUnit)
						}
						if f0f0f4iter.DirectoryName.Pseudonym != nil {
							f0f0f4elemf0.SetPseudonym(*f0f0f4iter.DirectoryName.Pseudonym)
						}
						if f0f0f4iter.DirectoryName.SerialNumber != nil {
							f0f0f4elemf0.SetSerialNumber(*f0f0f4iter.DirectoryName.SerialNumber)
						}
						if f0f0f4iter.DirectoryName.State != nil {
							f0f0f4elemf0.SetState(*f0f0f4iter.DirectoryName.State)
						}
						if f0f0f4iter.DirectoryName.Surname != nil {
							f0f0f4elemf0.SetSurname(*f0f0f4iter.DirectoryName.Surname)
						}
						if f0f0f4iter.DirectoryName.Title != nil {
							f0f0f4elemf0.SetTitle(*f0f0f4iter.DirectoryName.Title)
						}
						f0f0f4elem.SetDirectoryName(f0f0f4elemf0)
					}
					if f0f0f4iter.DNSName != nil {
						f0f0f4elem.SetDnsName(*f0f0f4iter.DNSName)
					}
					if f0f0f4iter.EDIPartyName != nil {
						f0f0f4elemf2 := &svcsdk.EdiPartyName{}
						if f0f0f4iter.EDIPartyName.NameAssigner != nil {
							f0f0f4elemf2.SetNameAssigner(*f0f0f4iter.EDIPartyName.NameAssigner)
						}
						if f0f0f4iter.EDIPartyName.PartyName != nil {
							f0f0f4elemf2.SetPartyName(*f0f0f4iter.EDIPartyName.PartyName)
						}
						f0f0f4elem.SetEdiPartyName(f0f0f4elemf2)
					}
					if f0f0f4iter.IPAddress != nil {
						f0f0f4elem.SetIpAddress(*f0f0f4iter.IPAddress)
					}
					if f0f0f4iter.OtherName != nil {
						f0f0f4elemf4 := &svcsdk.OtherName{}
						if f0f0f4iter.OtherName.TypeID != nil {
							f0f0f4elemf4.SetTypeId(*f0f0f4iter.OtherName.TypeID)
						}
						if f0f0f4iter.OtherName.Value != nil {
							f0f0f4elemf4.SetValue(*f0f0f4iter.OtherName.Value)
						}
						f0f0f4elem.SetOtherName(f0f0f4elemf4)
					}
					if f0f0f4iter.RegisteredID != nil {
						f0f0f4elem.SetRegisteredId(*f0f0f4iter.RegisteredID)
					}
					if f0f0f4iter.RFC822Name != nil {
						f0f0f4elem.SetRfc822Name(*f0f0f4iter.RFC822Name)
					}
					if f0f0f4iter.UniformResourceIdentifier != nil {
						f0f0f4elem.SetUniformResourceIdentifier(*f0f0f4iter.UniformResourceIdentifier)
					}
					f0f0f4 = append(f0f0f4, f0f0f4elem)
				}
				f0f0.SetSubjectAlternativeNames(f0f0f4)
			}
			f0.SetExtensions(f0f0)
		}
		if r.ko.Spec.APIPassthrough.Subject != nil {
			f0f1 := &svcsdk.ASN1Subject{}
			if r.ko.Spec.APIPassthrough.Subject.CommonName != nil {
				f0f1.SetCommonName(*r.ko.Spec.APIPassthrough.Subject.CommonName)
			}
			if r.ko.Spec.APIPassthrough.Subject.Country != nil {
				f0f1.SetCountry(*r.ko.Spec.APIPassthrough.Subject.Country)
			}
			if r.ko.Spec.APIPassthrough.Subject.CustomAttributes != nil {
				f0f1f2 := []*svcsdk.CustomAttribute{}
				for _, f0f1f2iter := range r.ko.Spec.APIPassthrough.Subject.CustomAttributes {
					f0f1f2elem := &svcsdk.CustomAttribute{}
					if f0f1f2iter.ObjectIdentifier != nil {
						f0f1f2elem.SetObjectIdentifier(*f0f1f2iter.ObjectIdentifier)
					}
					if f0f1f2iter.Value != nil {
						f0f1f2elem.SetValue(*f0f1f2iter.Value)
					}
					f0f1f2 = append(f0f1f2, f0f1f2elem)
				}
				f0f1.SetCustomAttributes(f0f1f2)
			}
			if r.ko.Spec.APIPassthrough.Subject.DistinguishedNameQualifier != nil {
				f0f1.SetDistinguishedNameQualifier(*r.ko.Spec.APIPassthrough.Subject.DistinguishedNameQualifier)
			}
			if r.ko.Spec.APIPassthrough.Subject.GenerationQualifier != nil {
				f0f1.SetGenerationQualifier(*r.ko.Spec.APIPassthrough.Subject.GenerationQualifier)
			}
			if r.ko.Spec.APIPassthrough.Subject.GivenName != nil {
				f0f1.SetGivenName(*r.ko.Spec.APIPassthrough.Subject.GivenName)
			}
			if r.ko.Spec.APIPassthrough.Subject.Initials != nil {
				f0f1.SetInitials(*r.ko.Spec.APIPassthrough.Subject.Initials)
			}
			if r.ko.Spec.APIPassthrough.Subject.Locality != nil {
				f0f1.SetLocality(*r.ko.Spec.APIPassthrough.Subject.Locality)
			}
			if r.ko.Spec.APIPassthrough.Subject.Organization != nil {
				f0f1.SetOrganization(*r.ko.Spec.APIPassthrough.Subject.Organization)
			}
			if r.ko.Spec.APIPassthrough.Subject.OrganizationalUnit != nil {
				f0f1.SetOrganizationalUnit(*r.ko.Spec.APIPassthrough.Subject.OrganizationalUnit)
			}
			if r.ko.Spec.APIPassthrough.Subject.Pseudonym != nil {
				f0f1.SetPseudonym(*r.ko.Spec.APIPassthrough.Subject.Pseudonym)
			}
			if r.ko.Spec.APIPassthrough.Subject.SerialNumber != nil {
				f0f1.SetSerialNumber(*r.ko.Spec.APIPassthrough.Subject.SerialNumber)
			}
			if r.ko.Spec.APIPassthrough.Subject.State != nil {
				f0f1.SetState(*r.ko.Spec.APIPassthrough.Subject.State)
			}
			if r.ko.Spec.APIPassthrough.Subject.Surname != nil {
				f0f1.SetSurname(*r.ko.Spec.APIPassthrough.Subject.Surname)
			}
			if r.ko.Spec.APIPassthrough.Subject.Title != nil {
				f0f1.SetTitle(*r.ko.Spec.APIPassthrough.Subject.Title)
			}
			f0.SetSubject(f0f1)
		}
		res.SetApiPassthrough(f0)
	}
	if r.ko.Spec.CertificateAuthorityARN != nil {
		res.SetCertificateAuthorityArn(*r.ko.Spec.CertificateAuthorityARN)
	}
	if r.ko.Spec.SigningAlgorithm != nil {
		res.SetSigningAlgorithm(*r.ko.Spec.SigningAlgorithm)
	}
	if r.ko.Spec.TemplateARN != nil {
		res.SetTemplateArn(*r.ko.Spec.TemplateARN)
	}
	if r.ko.Spec.Validity != nil {
		f4 := &svcsdk.Validity{}
		if r.ko.Spec.Validity.Type != nil {
			f4.SetType(*r.ko.Spec.Validity.Type)
		}
		if r.ko.Spec.Validity.Value != nil {
			f4.SetValue(*r.ko.Spec.Validity.Value)
		}
		res.SetValidity(f4)
	}
	if r.ko.Spec.ValidityNotBefore != nil {
		f5 := &svcsdk.Validity{}
		if r.ko.Spec.ValidityNotBefore.Type != nil {
			f5.SetType(*r.ko.Spec.ValidityNotBefore.Type)
		}
		if r.ko.Spec.ValidityNotBefore.Value != nil {
			f5.SetValue(*r.ko.Spec.ValidityNotBefore.Value)
		}
		res.SetValidityNotBefore(f5)
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
	return nil, ackerr.NewTerminalError(ackerr.NotImplemented)
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
	// TODO(jaypipes): Figure this out...
	return nil, nil

}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Certificate,
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
		"InvalidArgsException",
		"InvalidArnException",
		"InvalidStateException",
		"LimitExceededException",
		"MalformedCSRException":
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
	if delta.DifferentAt("Spec.APIPassthrough") {
		fields = append(fields, "APIPassthrough")
	}
	if delta.DifferentAt("Spec.CertificateAuthorityARN") {
		fields = append(fields, "CertificateAuthorityARN")
	}
	if delta.DifferentAt("Spec.CertificateSigningRequest") {
		fields = append(fields, "CertificateSigningRequest")
	}
	if delta.DifferentAt("Spec.SigningAlgorithm") {
		fields = append(fields, "SigningAlgorithm")
	}
	if delta.DifferentAt("Spec.TemplateARN") {
		fields = append(fields, "TemplateARN")
	}
	if delta.DifferentAt("Spec.Validity") {
		fields = append(fields, "Validity")
	}
	if delta.DifferentAt("Spec.ValidityNotBefore") {
		fields = append(fields, "ValidityNotBefore")
	}

	return fields
}
