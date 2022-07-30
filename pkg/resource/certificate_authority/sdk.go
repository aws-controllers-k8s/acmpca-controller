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

package certificate_authority

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
	_ = &svcapitypes.CertificateAuthority{}
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

	var resp *svcsdk.DescribeCertificateAuthorityOutput
	resp, err = rm.sdkapi.DescribeCertificateAuthorityWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeCertificateAuthority", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.CertificateAuthority.Arn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.CertificateAuthority.Arn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.CertificateAuthority.CertificateAuthorityConfiguration != nil {
		f1 := &svcapitypes.CertificateAuthorityConfiguration{}
		if resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions != nil {
			f1f0 := &svcapitypes.CSRExtensions{}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage != nil {
				f1f0f0 := &svcapitypes.KeyUsage{}
				if resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.CRLSign != nil {
					f1f0f0.CRLSign = resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.CRLSign
				}
				if resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.DataEncipherment != nil {
					f1f0f0.DataEncipherment = resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.DataEncipherment
				}
				if resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.DecipherOnly != nil {
					f1f0f0.DecipherOnly = resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.DecipherOnly
				}
				if resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.DigitalSignature != nil {
					f1f0f0.DigitalSignature = resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.DigitalSignature
				}
				if resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.EncipherOnly != nil {
					f1f0f0.EncipherOnly = resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.EncipherOnly
				}
				if resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.KeyAgreement != nil {
					f1f0f0.KeyAgreement = resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.KeyAgreement
				}
				if resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.KeyCertSign != nil {
					f1f0f0.KeyCertSign = resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.KeyCertSign
				}
				if resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.KeyEncipherment != nil {
					f1f0f0.KeyEncipherment = resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.KeyEncipherment
				}
				if resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.NonRepudiation != nil {
					f1f0f0.NonRepudiation = resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.KeyUsage.NonRepudiation
				}
				f1f0.KeyUsage = f1f0f0
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.SubjectInformationAccess != nil {
				f1f0f1 := []*svcapitypes.AccessDescription{}
				for _, f1f0f1iter := range resp.CertificateAuthority.CertificateAuthorityConfiguration.CsrExtensions.SubjectInformationAccess {
					f1f0f1elem := &svcapitypes.AccessDescription{}
					if f1f0f1iter.AccessLocation != nil {
						f1f0f1elemf0 := &svcapitypes.GeneralName{}
						if f1f0f1iter.AccessLocation.DirectoryName != nil {
							f1f0f1elemf0f0 := &svcapitypes.ASN1Subject{}
							if f1f0f1iter.AccessLocation.DirectoryName.CommonName != nil {
								f1f0f1elemf0f0.CommonName = f1f0f1iter.AccessLocation.DirectoryName.CommonName
							}
							if f1f0f1iter.AccessLocation.DirectoryName.Country != nil {
								f1f0f1elemf0f0.Country = f1f0f1iter.AccessLocation.DirectoryName.Country
							}
							if f1f0f1iter.AccessLocation.DirectoryName.DistinguishedNameQualifier != nil {
								f1f0f1elemf0f0.DistinguishedNameQualifier = f1f0f1iter.AccessLocation.DirectoryName.DistinguishedNameQualifier
							}
							if f1f0f1iter.AccessLocation.DirectoryName.GenerationQualifier != nil {
								f1f0f1elemf0f0.GenerationQualifier = f1f0f1iter.AccessLocation.DirectoryName.GenerationQualifier
							}
							if f1f0f1iter.AccessLocation.DirectoryName.GivenName != nil {
								f1f0f1elemf0f0.GivenName = f1f0f1iter.AccessLocation.DirectoryName.GivenName
							}
							if f1f0f1iter.AccessLocation.DirectoryName.Initials != nil {
								f1f0f1elemf0f0.Initials = f1f0f1iter.AccessLocation.DirectoryName.Initials
							}
							if f1f0f1iter.AccessLocation.DirectoryName.Locality != nil {
								f1f0f1elemf0f0.Locality = f1f0f1iter.AccessLocation.DirectoryName.Locality
							}
							if f1f0f1iter.AccessLocation.DirectoryName.Organization != nil {
								f1f0f1elemf0f0.Organization = f1f0f1iter.AccessLocation.DirectoryName.Organization
							}
							if f1f0f1iter.AccessLocation.DirectoryName.OrganizationalUnit != nil {
								f1f0f1elemf0f0.OrganizationalUnit = f1f0f1iter.AccessLocation.DirectoryName.OrganizationalUnit
							}
							if f1f0f1iter.AccessLocation.DirectoryName.Pseudonym != nil {
								f1f0f1elemf0f0.Pseudonym = f1f0f1iter.AccessLocation.DirectoryName.Pseudonym
							}
							if f1f0f1iter.AccessLocation.DirectoryName.SerialNumber != nil {
								f1f0f1elemf0f0.SerialNumber = f1f0f1iter.AccessLocation.DirectoryName.SerialNumber
							}
							if f1f0f1iter.AccessLocation.DirectoryName.State != nil {
								f1f0f1elemf0f0.State = f1f0f1iter.AccessLocation.DirectoryName.State
							}
							if f1f0f1iter.AccessLocation.DirectoryName.Surname != nil {
								f1f0f1elemf0f0.Surname = f1f0f1iter.AccessLocation.DirectoryName.Surname
							}
							if f1f0f1iter.AccessLocation.DirectoryName.Title != nil {
								f1f0f1elemf0f0.Title = f1f0f1iter.AccessLocation.DirectoryName.Title
							}
							f1f0f1elemf0.DirectoryName = f1f0f1elemf0f0
						}
						if f1f0f1iter.AccessLocation.DnsName != nil {
							f1f0f1elemf0.DNSName = f1f0f1iter.AccessLocation.DnsName
						}
						if f1f0f1iter.AccessLocation.EdiPartyName != nil {
							f1f0f1elemf0f2 := &svcapitypes.EDIPartyName{}
							if f1f0f1iter.AccessLocation.EdiPartyName.NameAssigner != nil {
								f1f0f1elemf0f2.NameAssigner = f1f0f1iter.AccessLocation.EdiPartyName.NameAssigner
							}
							if f1f0f1iter.AccessLocation.EdiPartyName.PartyName != nil {
								f1f0f1elemf0f2.PartyName = f1f0f1iter.AccessLocation.EdiPartyName.PartyName
							}
							f1f0f1elemf0.EDIPartyName = f1f0f1elemf0f2
						}
						if f1f0f1iter.AccessLocation.IpAddress != nil {
							f1f0f1elemf0.IPAddress = f1f0f1iter.AccessLocation.IpAddress
						}
						if f1f0f1iter.AccessLocation.OtherName != nil {
							f1f0f1elemf0f4 := &svcapitypes.OtherName{}
							if f1f0f1iter.AccessLocation.OtherName.TypeId != nil {
								f1f0f1elemf0f4.TypeID = f1f0f1iter.AccessLocation.OtherName.TypeId
							}
							if f1f0f1iter.AccessLocation.OtherName.Value != nil {
								f1f0f1elemf0f4.Value = f1f0f1iter.AccessLocation.OtherName.Value
							}
							f1f0f1elemf0.OtherName = f1f0f1elemf0f4
						}
						if f1f0f1iter.AccessLocation.RegisteredId != nil {
							f1f0f1elemf0.RegisteredID = f1f0f1iter.AccessLocation.RegisteredId
						}
						if f1f0f1iter.AccessLocation.Rfc822Name != nil {
							f1f0f1elemf0.RFC822Name = f1f0f1iter.AccessLocation.Rfc822Name
						}
						if f1f0f1iter.AccessLocation.UniformResourceIdentifier != nil {
							f1f0f1elemf0.UniformResourceIdentifier = f1f0f1iter.AccessLocation.UniformResourceIdentifier
						}
						f1f0f1elem.AccessLocation = f1f0f1elemf0
					}
					if f1f0f1iter.AccessMethod != nil {
						f1f0f1elemf1 := &svcapitypes.AccessMethod{}
						if f1f0f1iter.AccessMethod.AccessMethodType != nil {
							f1f0f1elemf1.AccessMethodType = f1f0f1iter.AccessMethod.AccessMethodType
						}
						if f1f0f1iter.AccessMethod.CustomObjectIdentifier != nil {
							f1f0f1elemf1.CustomObjectIdentifier = f1f0f1iter.AccessMethod.CustomObjectIdentifier
						}
						f1f0f1elem.AccessMethod = f1f0f1elemf1
					}
					f1f0f1 = append(f1f0f1, f1f0f1elem)
				}
				f1f0.SubjectInformationAccess = f1f0f1
			}
			f1.CSRExtensions = f1f0
		}
		if resp.CertificateAuthority.CertificateAuthorityConfiguration.KeyAlgorithm != nil {
			f1.KeyAlgorithm = resp.CertificateAuthority.CertificateAuthorityConfiguration.KeyAlgorithm
		}
		if resp.CertificateAuthority.CertificateAuthorityConfiguration.SigningAlgorithm != nil {
			f1.SigningAlgorithm = resp.CertificateAuthority.CertificateAuthorityConfiguration.SigningAlgorithm
		}
		if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject != nil {
			f1f3 := &svcapitypes.ASN1Subject{}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.CommonName != nil {
				f1f3.CommonName = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.CommonName
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Country != nil {
				f1f3.Country = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Country
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier != nil {
				f1f3.DistinguishedNameQualifier = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.GenerationQualifier != nil {
				f1f3.GenerationQualifier = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.GenerationQualifier
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.GivenName != nil {
				f1f3.GivenName = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.GivenName
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Initials != nil {
				f1f3.Initials = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Initials
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Locality != nil {
				f1f3.Locality = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Locality
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Organization != nil {
				f1f3.Organization = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Organization
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.OrganizationalUnit != nil {
				f1f3.OrganizationalUnit = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.OrganizationalUnit
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Pseudonym != nil {
				f1f3.Pseudonym = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Pseudonym
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.SerialNumber != nil {
				f1f3.SerialNumber = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.SerialNumber
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.State != nil {
				f1f3.State = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.State
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Surname != nil {
				f1f3.Surname = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Surname
			}
			if resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Title != nil {
				f1f3.Title = resp.CertificateAuthority.CertificateAuthorityConfiguration.Subject.Title
			}
			f1.Subject = f1f3
		}
		ko.Spec.CertificateAuthorityConfiguration = f1
	} else {
		ko.Spec.CertificateAuthorityConfiguration = nil
	}
	if resp.CertificateAuthority.KeyStorageSecurityStandard != nil {
		ko.Spec.KeyStorageSecurityStandard = resp.CertificateAuthority.KeyStorageSecurityStandard
	} else {
		ko.Spec.KeyStorageSecurityStandard = nil
	}
	if resp.CertificateAuthority.RevocationConfiguration != nil {
		f10 := &svcapitypes.RevocationConfiguration{}
		if resp.CertificateAuthority.RevocationConfiguration.CrlConfiguration != nil {
			f10f0 := &svcapitypes.CRLConfiguration{}
			if resp.CertificateAuthority.RevocationConfiguration.CrlConfiguration.CustomCname != nil {
				f10f0.CustomCNAME = resp.CertificateAuthority.RevocationConfiguration.CrlConfiguration.CustomCname
			}
			if resp.CertificateAuthority.RevocationConfiguration.CrlConfiguration.Enabled != nil {
				f10f0.Enabled = resp.CertificateAuthority.RevocationConfiguration.CrlConfiguration.Enabled
			}
			if resp.CertificateAuthority.RevocationConfiguration.CrlConfiguration.ExpirationInDays != nil {
				f10f0.ExpirationInDays = resp.CertificateAuthority.RevocationConfiguration.CrlConfiguration.ExpirationInDays
			}
			if resp.CertificateAuthority.RevocationConfiguration.CrlConfiguration.S3BucketName != nil {
				f10f0.S3BucketName = resp.CertificateAuthority.RevocationConfiguration.CrlConfiguration.S3BucketName
			}
			if resp.CertificateAuthority.RevocationConfiguration.CrlConfiguration.S3ObjectAcl != nil {
				f10f0.S3ObjectACL = resp.CertificateAuthority.RevocationConfiguration.CrlConfiguration.S3ObjectAcl
			}
			f10.CRLConfiguration = f10f0
		}
		if resp.CertificateAuthority.RevocationConfiguration.OcspConfiguration != nil {
			f10f1 := &svcapitypes.OCSPConfiguration{}
			if resp.CertificateAuthority.RevocationConfiguration.OcspConfiguration.Enabled != nil {
				f10f1.Enabled = resp.CertificateAuthority.RevocationConfiguration.OcspConfiguration.Enabled
			}
			if resp.CertificateAuthority.RevocationConfiguration.OcspConfiguration.OcspCustomCname != nil {
				f10f1.OCSPCustomCNAME = resp.CertificateAuthority.RevocationConfiguration.OcspConfiguration.OcspCustomCname
			}
			f10.OCSPConfiguration = f10f1
		}
		ko.Spec.RevocationConfiguration = f10
	} else {
		ko.Spec.RevocationConfiguration = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return (r.ko.Status.ACKResourceMetadata == nil || r.ko.Status.ACKResourceMetadata.ARN == nil)

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeCertificateAuthorityInput, error) {
	res := &svcsdk.DescribeCertificateAuthorityInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetCertificateAuthorityArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
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

	var resp *svcsdk.CreateCertificateAuthorityOutput
	_ = resp
	resp, err = rm.sdkapi.CreateCertificateAuthorityWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateCertificateAuthority", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.CertificateAuthorityArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.CertificateAuthorityArn)
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
) (*svcsdk.CreateCertificateAuthorityInput, error) {
	res := &svcsdk.CreateCertificateAuthorityInput{}

	if r.ko.Spec.CertificateAuthorityConfiguration != nil {
		f0 := &svcsdk.CertificateAuthorityConfiguration{}
		if r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions != nil {
			f0f0 := &svcsdk.CsrExtensions{}
			if r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage != nil {
				f0f0f0 := &svcsdk.KeyUsage{}
				if r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign != nil {
					f0f0f0.SetCRLSign(*r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign)
				}
				if r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment != nil {
					f0f0f0.SetDataEncipherment(*r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment)
				}
				if r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly != nil {
					f0f0f0.SetDecipherOnly(*r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly)
				}
				if r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature != nil {
					f0f0f0.SetDigitalSignature(*r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature)
				}
				if r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly != nil {
					f0f0f0.SetEncipherOnly(*r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly)
				}
				if r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement != nil {
					f0f0f0.SetKeyAgreement(*r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement)
				}
				if r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign != nil {
					f0f0f0.SetKeyCertSign(*r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign)
				}
				if r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment != nil {
					f0f0f0.SetKeyEncipherment(*r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment)
				}
				if r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation != nil {
					f0f0f0.SetNonRepudiation(*r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation)
				}
				f0f0.SetKeyUsage(f0f0f0)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess != nil {
				f0f0f1 := []*svcsdk.AccessDescription{}
				for _, f0f0f1iter := range r.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess {
					f0f0f1elem := &svcsdk.AccessDescription{}
					if f0f0f1iter.AccessLocation != nil {
						f0f0f1elemf0 := &svcsdk.GeneralName{}
						if f0f0f1iter.AccessLocation.DirectoryName != nil {
							f0f0f1elemf0f0 := &svcsdk.ASN1Subject{}
							if f0f0f1iter.AccessLocation.DirectoryName.CommonName != nil {
								f0f0f1elemf0f0.SetCommonName(*f0f0f1iter.AccessLocation.DirectoryName.CommonName)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.Country != nil {
								f0f0f1elemf0f0.SetCountry(*f0f0f1iter.AccessLocation.DirectoryName.Country)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.DistinguishedNameQualifier != nil {
								f0f0f1elemf0f0.SetDistinguishedNameQualifier(*f0f0f1iter.AccessLocation.DirectoryName.DistinguishedNameQualifier)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.GenerationQualifier != nil {
								f0f0f1elemf0f0.SetGenerationQualifier(*f0f0f1iter.AccessLocation.DirectoryName.GenerationQualifier)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.GivenName != nil {
								f0f0f1elemf0f0.SetGivenName(*f0f0f1iter.AccessLocation.DirectoryName.GivenName)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.Initials != nil {
								f0f0f1elemf0f0.SetInitials(*f0f0f1iter.AccessLocation.DirectoryName.Initials)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.Locality != nil {
								f0f0f1elemf0f0.SetLocality(*f0f0f1iter.AccessLocation.DirectoryName.Locality)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.Organization != nil {
								f0f0f1elemf0f0.SetOrganization(*f0f0f1iter.AccessLocation.DirectoryName.Organization)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.OrganizationalUnit != nil {
								f0f0f1elemf0f0.SetOrganizationalUnit(*f0f0f1iter.AccessLocation.DirectoryName.OrganizationalUnit)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.Pseudonym != nil {
								f0f0f1elemf0f0.SetPseudonym(*f0f0f1iter.AccessLocation.DirectoryName.Pseudonym)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.SerialNumber != nil {
								f0f0f1elemf0f0.SetSerialNumber(*f0f0f1iter.AccessLocation.DirectoryName.SerialNumber)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.State != nil {
								f0f0f1elemf0f0.SetState(*f0f0f1iter.AccessLocation.DirectoryName.State)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.Surname != nil {
								f0f0f1elemf0f0.SetSurname(*f0f0f1iter.AccessLocation.DirectoryName.Surname)
							}
							if f0f0f1iter.AccessLocation.DirectoryName.Title != nil {
								f0f0f1elemf0f0.SetTitle(*f0f0f1iter.AccessLocation.DirectoryName.Title)
							}
							f0f0f1elemf0.SetDirectoryName(f0f0f1elemf0f0)
						}
						if f0f0f1iter.AccessLocation.DNSName != nil {
							f0f0f1elemf0.SetDnsName(*f0f0f1iter.AccessLocation.DNSName)
						}
						if f0f0f1iter.AccessLocation.EDIPartyName != nil {
							f0f0f1elemf0f2 := &svcsdk.EdiPartyName{}
							if f0f0f1iter.AccessLocation.EDIPartyName.NameAssigner != nil {
								f0f0f1elemf0f2.SetNameAssigner(*f0f0f1iter.AccessLocation.EDIPartyName.NameAssigner)
							}
							if f0f0f1iter.AccessLocation.EDIPartyName.PartyName != nil {
								f0f0f1elemf0f2.SetPartyName(*f0f0f1iter.AccessLocation.EDIPartyName.PartyName)
							}
							f0f0f1elemf0.SetEdiPartyName(f0f0f1elemf0f2)
						}
						if f0f0f1iter.AccessLocation.IPAddress != nil {
							f0f0f1elemf0.SetIpAddress(*f0f0f1iter.AccessLocation.IPAddress)
						}
						if f0f0f1iter.AccessLocation.OtherName != nil {
							f0f0f1elemf0f4 := &svcsdk.OtherName{}
							if f0f0f1iter.AccessLocation.OtherName.TypeID != nil {
								f0f0f1elemf0f4.SetTypeId(*f0f0f1iter.AccessLocation.OtherName.TypeID)
							}
							if f0f0f1iter.AccessLocation.OtherName.Value != nil {
								f0f0f1elemf0f4.SetValue(*f0f0f1iter.AccessLocation.OtherName.Value)
							}
							f0f0f1elemf0.SetOtherName(f0f0f1elemf0f4)
						}
						if f0f0f1iter.AccessLocation.RegisteredID != nil {
							f0f0f1elemf0.SetRegisteredId(*f0f0f1iter.AccessLocation.RegisteredID)
						}
						if f0f0f1iter.AccessLocation.RFC822Name != nil {
							f0f0f1elemf0.SetRfc822Name(*f0f0f1iter.AccessLocation.RFC822Name)
						}
						if f0f0f1iter.AccessLocation.UniformResourceIdentifier != nil {
							f0f0f1elemf0.SetUniformResourceIdentifier(*f0f0f1iter.AccessLocation.UniformResourceIdentifier)
						}
						f0f0f1elem.SetAccessLocation(f0f0f1elemf0)
					}
					if f0f0f1iter.AccessMethod != nil {
						f0f0f1elemf1 := &svcsdk.AccessMethod{}
						if f0f0f1iter.AccessMethod.AccessMethodType != nil {
							f0f0f1elemf1.SetAccessMethodType(*f0f0f1iter.AccessMethod.AccessMethodType)
						}
						if f0f0f1iter.AccessMethod.CustomObjectIdentifier != nil {
							f0f0f1elemf1.SetCustomObjectIdentifier(*f0f0f1iter.AccessMethod.CustomObjectIdentifier)
						}
						f0f0f1elem.SetAccessMethod(f0f0f1elemf1)
					}
					f0f0f1 = append(f0f0f1, f0f0f1elem)
				}
				f0f0.SetSubjectInformationAccess(f0f0f1)
			}
			f0.SetCsrExtensions(f0f0)
		}
		if r.ko.Spec.CertificateAuthorityConfiguration.KeyAlgorithm != nil {
			f0.SetKeyAlgorithm(*r.ko.Spec.CertificateAuthorityConfiguration.KeyAlgorithm)
		}
		if r.ko.Spec.CertificateAuthorityConfiguration.SigningAlgorithm != nil {
			f0.SetSigningAlgorithm(*r.ko.Spec.CertificateAuthorityConfiguration.SigningAlgorithm)
		}
		if r.ko.Spec.CertificateAuthorityConfiguration.Subject != nil {
			f0f3 := &svcsdk.ASN1Subject{}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.CommonName != nil {
				f0f3.SetCommonName(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.CommonName)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.Country != nil {
				f0f3.SetCountry(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.Country)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier != nil {
				f0f3.SetDistinguishedNameQualifier(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier != nil {
				f0f3.SetGenerationQualifier(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.GivenName != nil {
				f0f3.SetGivenName(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.GivenName)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.Initials != nil {
				f0f3.SetInitials(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.Initials)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.Locality != nil {
				f0f3.SetLocality(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.Locality)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.Organization != nil {
				f0f3.SetOrganization(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.Organization)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit != nil {
				f0f3.SetOrganizationalUnit(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.Pseudonym != nil {
				f0f3.SetPseudonym(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.Pseudonym)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.SerialNumber != nil {
				f0f3.SetSerialNumber(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.SerialNumber)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.State != nil {
				f0f3.SetState(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.State)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.Surname != nil {
				f0f3.SetSurname(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.Surname)
			}
			if r.ko.Spec.CertificateAuthorityConfiguration.Subject.Title != nil {
				f0f3.SetTitle(*r.ko.Spec.CertificateAuthorityConfiguration.Subject.Title)
			}
			f0.SetSubject(f0f3)
		}
		res.SetCertificateAuthorityConfiguration(f0)
	}
	if r.ko.Spec.CertificateAuthorityType != nil {
		res.SetCertificateAuthorityType(*r.ko.Spec.CertificateAuthorityType)
	}
	if r.ko.Spec.KeyStorageSecurityStandard != nil {
		res.SetKeyStorageSecurityStandard(*r.ko.Spec.KeyStorageSecurityStandard)
	}
	if r.ko.Spec.RevocationConfiguration != nil {
		f3 := &svcsdk.RevocationConfiguration{}
		if r.ko.Spec.RevocationConfiguration.CRLConfiguration != nil {
			f3f0 := &svcsdk.CrlConfiguration{}
			if r.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME != nil {
				f3f0.SetCustomCname(*r.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME)
			}
			if r.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled != nil {
				f3f0.SetEnabled(*r.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled)
			}
			if r.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays != nil {
				f3f0.SetExpirationInDays(*r.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays)
			}
			if r.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName != nil {
				f3f0.SetS3BucketName(*r.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName)
			}
			if r.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL != nil {
				f3f0.SetS3ObjectAcl(*r.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL)
			}
			f3.SetCrlConfiguration(f3f0)
		}
		if r.ko.Spec.RevocationConfiguration.OCSPConfiguration != nil {
			f3f1 := &svcsdk.OcspConfiguration{}
			if r.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled != nil {
				f3f1.SetEnabled(*r.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled)
			}
			if r.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME != nil {
				f3f1.SetOcspCustomCname(*r.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME)
			}
			f3.SetOcspConfiguration(f3f1)
		}
		res.SetRevocationConfiguration(f3)
	}
	if r.ko.Spec.Tags != nil {
		f4 := []*svcsdk.Tag{}
		for _, f4iter := range r.ko.Spec.Tags {
			f4elem := &svcsdk.Tag{}
			if f4iter.Key != nil {
				f4elem.SetKey(*f4iter.Key)
			}
			if f4iter.Value != nil {
				f4elem.SetValue(*f4iter.Value)
			}
			f4 = append(f4, f4elem)
		}
		res.SetTags(f4)
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
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newUpdateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
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

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.UpdateCertificateAuthorityInput, error) {
	res := &svcsdk.UpdateCertificateAuthorityInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetCertificateAuthorityArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}
	if r.ko.Spec.RevocationConfiguration != nil {
		f1 := &svcsdk.RevocationConfiguration{}
		if r.ko.Spec.RevocationConfiguration.CRLConfiguration != nil {
			f1f0 := &svcsdk.CrlConfiguration{}
			if r.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME != nil {
				f1f0.SetCustomCname(*r.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME)
			}
			if r.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled != nil {
				f1f0.SetEnabled(*r.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled)
			}
			if r.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays != nil {
				f1f0.SetExpirationInDays(*r.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays)
			}
			if r.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName != nil {
				f1f0.SetS3BucketName(*r.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName)
			}
			if r.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL != nil {
				f1f0.SetS3ObjectAcl(*r.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL)
			}
			f1.SetCrlConfiguration(f1f0)
		}
		if r.ko.Spec.RevocationConfiguration.OCSPConfiguration != nil {
			f1f1 := &svcsdk.OcspConfiguration{}
			if r.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled != nil {
				f1f1.SetEnabled(*r.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled)
			}
			if r.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME != nil {
				f1f1.SetOcspCustomCname(*r.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME)
			}
			f1.SetOcspConfiguration(f1f1)
		}
		res.SetRevocationConfiguration(f1)
	}

	return res, nil
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
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteCertificateAuthorityOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteCertificateAuthorityWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteCertificateAuthority", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteCertificateAuthorityInput, error) {
	res := &svcsdk.DeleteCertificateAuthorityInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetCertificateAuthorityArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.CertificateAuthority,
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
		"ValidationException":
		return true
	default:
		return false
	}
}
