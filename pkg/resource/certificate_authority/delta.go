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
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	customSetDefaults(a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration, b.ko.Spec.CertificateAuthorityConfiguration) {
		delta.Add("Spec.CertificateAuthorityConfiguration", a.ko.Spec.CertificateAuthorityConfiguration, b.ko.Spec.CertificateAuthorityConfiguration)
	} else if a.ko.Spec.CertificateAuthorityConfiguration != nil && b.ko.Spec.CertificateAuthorityConfiguration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions) {
			delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions)
		} else if a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions != nil && b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage) {
				delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage != nil && b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign) {
					delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign)
				} else if a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign != nil && b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign != nil {
					if *a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign != *b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign {
						delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.CRLSign)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment) {
					delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment)
				} else if a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment != nil && b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment != nil {
					if *a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment != *b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment {
						delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DataEncipherment)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly) {
					delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly)
				} else if a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly != nil && b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly != nil {
					if *a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly != *b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly {
						delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DecipherOnly)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature) {
					delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature)
				} else if a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature != nil && b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature != nil {
					if *a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature != *b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature {
						delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.DigitalSignature)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly) {
					delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly)
				} else if a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly != nil && b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly != nil {
					if *a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly != *b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly {
						delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.EncipherOnly)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement) {
					delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement)
				} else if a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement != nil && b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement != nil {
					if *a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement != *b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement {
						delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyAgreement)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign) {
					delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign)
				} else if a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign != nil && b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign != nil {
					if *a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign != *b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign {
						delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyCertSign)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment) {
					delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment)
				} else if a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment != nil && b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment != nil {
					if *a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment != *b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment {
						delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.KeyEncipherment)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation) {
					delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation)
				} else if a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation != nil && b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation != nil {
					if *a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation != *b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation {
						delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.KeyUsage.NonRepudiation)
					}
				}
			}
			if len(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess) != len(b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess) {
				delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess)
			} else if len(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess) > 0 {
				if !reflect.DeepEqual(a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess) {
					delta.Add("Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess", a.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess, b.ko.Spec.CertificateAuthorityConfiguration.CSRExtensions.SubjectInformationAccess)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.KeyAlgorithm, b.ko.Spec.CertificateAuthorityConfiguration.KeyAlgorithm) {
			delta.Add("Spec.CertificateAuthorityConfiguration.KeyAlgorithm", a.ko.Spec.CertificateAuthorityConfiguration.KeyAlgorithm, b.ko.Spec.CertificateAuthorityConfiguration.KeyAlgorithm)
		} else if a.ko.Spec.CertificateAuthorityConfiguration.KeyAlgorithm != nil && b.ko.Spec.CertificateAuthorityConfiguration.KeyAlgorithm != nil {
			if *a.ko.Spec.CertificateAuthorityConfiguration.KeyAlgorithm != *b.ko.Spec.CertificateAuthorityConfiguration.KeyAlgorithm {
				delta.Add("Spec.CertificateAuthorityConfiguration.KeyAlgorithm", a.ko.Spec.CertificateAuthorityConfiguration.KeyAlgorithm, b.ko.Spec.CertificateAuthorityConfiguration.KeyAlgorithm)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.SigningAlgorithm, b.ko.Spec.CertificateAuthorityConfiguration.SigningAlgorithm) {
			delta.Add("Spec.CertificateAuthorityConfiguration.SigningAlgorithm", a.ko.Spec.CertificateAuthorityConfiguration.SigningAlgorithm, b.ko.Spec.CertificateAuthorityConfiguration.SigningAlgorithm)
		} else if a.ko.Spec.CertificateAuthorityConfiguration.SigningAlgorithm != nil && b.ko.Spec.CertificateAuthorityConfiguration.SigningAlgorithm != nil {
			if *a.ko.Spec.CertificateAuthorityConfiguration.SigningAlgorithm != *b.ko.Spec.CertificateAuthorityConfiguration.SigningAlgorithm {
				delta.Add("Spec.CertificateAuthorityConfiguration.SigningAlgorithm", a.ko.Spec.CertificateAuthorityConfiguration.SigningAlgorithm, b.ko.Spec.CertificateAuthorityConfiguration.SigningAlgorithm)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject, b.ko.Spec.CertificateAuthorityConfiguration.Subject) {
			delta.Add("Spec.CertificateAuthorityConfiguration.Subject", a.ko.Spec.CertificateAuthorityConfiguration.Subject, b.ko.Spec.CertificateAuthorityConfiguration.Subject)
		} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.CommonName, b.ko.Spec.CertificateAuthorityConfiguration.Subject.CommonName) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.CommonName", a.ko.Spec.CertificateAuthorityConfiguration.Subject.CommonName, b.ko.Spec.CertificateAuthorityConfiguration.Subject.CommonName)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.CommonName != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.CommonName != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.CommonName != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.CommonName {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.CommonName", a.ko.Spec.CertificateAuthorityConfiguration.Subject.CommonName, b.ko.Spec.CertificateAuthorityConfiguration.Subject.CommonName)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.Country, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Country) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Country", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Country, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Country)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.Country != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.Country != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.Country != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.Country {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Country", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Country, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Country)
				}
			}
			if len(a.ko.Spec.CertificateAuthorityConfiguration.Subject.CustomAttributes) != len(b.ko.Spec.CertificateAuthorityConfiguration.Subject.CustomAttributes) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.CustomAttributes", a.ko.Spec.CertificateAuthorityConfiguration.Subject.CustomAttributes, b.ko.Spec.CertificateAuthorityConfiguration.Subject.CustomAttributes)
			} else if len(a.ko.Spec.CertificateAuthorityConfiguration.Subject.CustomAttributes) > 0 {
				if !reflect.DeepEqual(a.ko.Spec.CertificateAuthorityConfiguration.Subject.CustomAttributes, b.ko.Spec.CertificateAuthorityConfiguration.Subject.CustomAttributes) {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.CustomAttributes", a.ko.Spec.CertificateAuthorityConfiguration.Subject.CustomAttributes, b.ko.Spec.CertificateAuthorityConfiguration.Subject.CustomAttributes)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier, b.ko.Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier", a.ko.Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier, b.ko.Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier", a.ko.Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier, b.ko.Spec.CertificateAuthorityConfiguration.Subject.DistinguishedNameQualifier)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier, b.ko.Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier", a.ko.Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier, b.ko.Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier", a.ko.Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier, b.ko.Spec.CertificateAuthorityConfiguration.Subject.GenerationQualifier)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.GivenName, b.ko.Spec.CertificateAuthorityConfiguration.Subject.GivenName) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.GivenName", a.ko.Spec.CertificateAuthorityConfiguration.Subject.GivenName, b.ko.Spec.CertificateAuthorityConfiguration.Subject.GivenName)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.GivenName != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.GivenName != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.GivenName != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.GivenName {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.GivenName", a.ko.Spec.CertificateAuthorityConfiguration.Subject.GivenName, b.ko.Spec.CertificateAuthorityConfiguration.Subject.GivenName)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.Initials, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Initials) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Initials", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Initials, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Initials)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.Initials != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.Initials != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.Initials != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.Initials {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Initials", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Initials, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Initials)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.Locality, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Locality) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Locality", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Locality, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Locality)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.Locality != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.Locality != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.Locality != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.Locality {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Locality", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Locality, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Locality)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.Organization, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Organization) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Organization", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Organization, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Organization)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.Organization != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.Organization != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.Organization != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.Organization {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Organization", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Organization, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Organization)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit, b.ko.Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit", a.ko.Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit, b.ko.Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit", a.ko.Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit, b.ko.Spec.CertificateAuthorityConfiguration.Subject.OrganizationalUnit)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.Pseudonym, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Pseudonym) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Pseudonym", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Pseudonym, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Pseudonym)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.Pseudonym != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.Pseudonym != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.Pseudonym != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.Pseudonym {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Pseudonym", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Pseudonym, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Pseudonym)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.SerialNumber, b.ko.Spec.CertificateAuthorityConfiguration.Subject.SerialNumber) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.SerialNumber", a.ko.Spec.CertificateAuthorityConfiguration.Subject.SerialNumber, b.ko.Spec.CertificateAuthorityConfiguration.Subject.SerialNumber)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.SerialNumber != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.SerialNumber != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.SerialNumber != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.SerialNumber {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.SerialNumber", a.ko.Spec.CertificateAuthorityConfiguration.Subject.SerialNumber, b.ko.Spec.CertificateAuthorityConfiguration.Subject.SerialNumber)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.State, b.ko.Spec.CertificateAuthorityConfiguration.Subject.State) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.State", a.ko.Spec.CertificateAuthorityConfiguration.Subject.State, b.ko.Spec.CertificateAuthorityConfiguration.Subject.State)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.State != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.State != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.State != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.State {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.State", a.ko.Spec.CertificateAuthorityConfiguration.Subject.State, b.ko.Spec.CertificateAuthorityConfiguration.Subject.State)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.Surname, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Surname) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Surname", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Surname, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Surname)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.Surname != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.Surname != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.Surname != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.Surname {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Surname", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Surname, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Surname)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityConfiguration.Subject.Title, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Title) {
				delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Title", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Title, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Title)
			} else if a.ko.Spec.CertificateAuthorityConfiguration.Subject.Title != nil && b.ko.Spec.CertificateAuthorityConfiguration.Subject.Title != nil {
				if *a.ko.Spec.CertificateAuthorityConfiguration.Subject.Title != *b.ko.Spec.CertificateAuthorityConfiguration.Subject.Title {
					delta.Add("Spec.CertificateAuthorityConfiguration.Subject.Title", a.ko.Spec.CertificateAuthorityConfiguration.Subject.Title, b.ko.Spec.CertificateAuthorityConfiguration.Subject.Title)
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.KeyStorageSecurityStandard, b.ko.Spec.KeyStorageSecurityStandard) {
		delta.Add("Spec.KeyStorageSecurityStandard", a.ko.Spec.KeyStorageSecurityStandard, b.ko.Spec.KeyStorageSecurityStandard)
	} else if a.ko.Spec.KeyStorageSecurityStandard != nil && b.ko.Spec.KeyStorageSecurityStandard != nil {
		if *a.ko.Spec.KeyStorageSecurityStandard != *b.ko.Spec.KeyStorageSecurityStandard {
			delta.Add("Spec.KeyStorageSecurityStandard", a.ko.Spec.KeyStorageSecurityStandard, b.ko.Spec.KeyStorageSecurityStandard)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RevocationConfiguration, b.ko.Spec.RevocationConfiguration) {
		delta.Add("Spec.RevocationConfiguration", a.ko.Spec.RevocationConfiguration, b.ko.Spec.RevocationConfiguration)
	} else if a.ko.Spec.RevocationConfiguration != nil && b.ko.Spec.RevocationConfiguration != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.RevocationConfiguration.CRLConfiguration, b.ko.Spec.RevocationConfiguration.CRLConfiguration) {
			delta.Add("Spec.RevocationConfiguration.CRLConfiguration", a.ko.Spec.RevocationConfiguration.CRLConfiguration, b.ko.Spec.RevocationConfiguration.CRLConfiguration)
		} else if a.ko.Spec.RevocationConfiguration.CRLConfiguration != nil && b.ko.Spec.RevocationConfiguration.CRLConfiguration != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME, b.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME) {
				delta.Add("Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME", a.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME, b.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME)
			} else if a.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME != nil && b.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME != nil {
				if *a.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME != *b.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME {
					delta.Add("Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME", a.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME, b.ko.Spec.RevocationConfiguration.CRLConfiguration.CustomCNAME)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled, b.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled) {
				delta.Add("Spec.RevocationConfiguration.CRLConfiguration.Enabled", a.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled, b.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled)
			} else if a.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled != nil && b.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled != nil {
				if *a.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled != *b.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled {
					delta.Add("Spec.RevocationConfiguration.CRLConfiguration.Enabled", a.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled, b.ko.Spec.RevocationConfiguration.CRLConfiguration.Enabled)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays, b.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays) {
				delta.Add("Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays", a.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays, b.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays)
			} else if a.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays != nil && b.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays != nil {
				if *a.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays != *b.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays {
					delta.Add("Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays", a.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays, b.ko.Spec.RevocationConfiguration.CRLConfiguration.ExpirationInDays)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName, b.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName) {
				delta.Add("Spec.RevocationConfiguration.CRLConfiguration.S3BucketName", a.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName, b.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName)
			} else if a.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName != nil && b.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName != nil {
				if *a.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName != *b.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName {
					delta.Add("Spec.RevocationConfiguration.CRLConfiguration.S3BucketName", a.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName, b.ko.Spec.RevocationConfiguration.CRLConfiguration.S3BucketName)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL, b.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL) {
				delta.Add("Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL", a.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL, b.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL)
			} else if a.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL != nil && b.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL != nil {
				if *a.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL != *b.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL {
					delta.Add("Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL", a.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL, b.ko.Spec.RevocationConfiguration.CRLConfiguration.S3ObjectACL)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.RevocationConfiguration.OCSPConfiguration, b.ko.Spec.RevocationConfiguration.OCSPConfiguration) {
			delta.Add("Spec.RevocationConfiguration.OCSPConfiguration", a.ko.Spec.RevocationConfiguration.OCSPConfiguration, b.ko.Spec.RevocationConfiguration.OCSPConfiguration)
		} else if a.ko.Spec.RevocationConfiguration.OCSPConfiguration != nil && b.ko.Spec.RevocationConfiguration.OCSPConfiguration != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled, b.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled) {
				delta.Add("Spec.RevocationConfiguration.OCSPConfiguration.Enabled", a.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled, b.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled)
			} else if a.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled != nil && b.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled != nil {
				if *a.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled != *b.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled {
					delta.Add("Spec.RevocationConfiguration.OCSPConfiguration.Enabled", a.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled, b.ko.Spec.RevocationConfiguration.OCSPConfiguration.Enabled)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME, b.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME) {
				delta.Add("Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME", a.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME, b.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME)
			} else if a.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME != nil && b.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME != nil {
				if *a.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME != *b.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME {
					delta.Add("Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME", a.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME, b.ko.Spec.RevocationConfiguration.OCSPConfiguration.OCSPCustomCNAME)
				}
			}
		}
	}
	desiredACKTags, _ := convertToOrderedACKTags(a.ko.Spec.Tags)
	latestACKTags, _ := convertToOrderedACKTags(b.ko.Spec.Tags)
	if !ackcompare.MapStringStringEqual(desiredACKTags, latestACKTags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Type, b.ko.Spec.Type) {
		delta.Add("Spec.Type", a.ko.Spec.Type, b.ko.Spec.Type)
	} else if a.ko.Spec.Type != nil && b.ko.Spec.Type != nil {
		if *a.ko.Spec.Type != *b.ko.Spec.Type {
			delta.Add("Spec.Type", a.ko.Spec.Type, b.ko.Spec.Type)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.UsageMode, b.ko.Spec.UsageMode) {
		delta.Add("Spec.UsageMode", a.ko.Spec.UsageMode, b.ko.Spec.UsageMode)
	} else if a.ko.Spec.UsageMode != nil && b.ko.Spec.UsageMode != nil {
		if *a.ko.Spec.UsageMode != *b.ko.Spec.UsageMode {
			delta.Add("Spec.UsageMode", a.ko.Spec.UsageMode, b.ko.Spec.UsageMode)
		}
	}

	return delta
}
