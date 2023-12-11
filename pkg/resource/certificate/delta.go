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

	if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough, b.ko.Spec.APIPassthrough) {
		delta.Add("Spec.APIPassthrough", a.ko.Spec.APIPassthrough, b.ko.Spec.APIPassthrough)
	} else if a.ko.Spec.APIPassthrough != nil && b.ko.Spec.APIPassthrough != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Extensions, b.ko.Spec.APIPassthrough.Extensions) {
			delta.Add("Spec.APIPassthrough.Extensions", a.ko.Spec.APIPassthrough.Extensions, b.ko.Spec.APIPassthrough.Extensions)
		} else if a.ko.Spec.APIPassthrough.Extensions != nil && b.ko.Spec.APIPassthrough.Extensions != nil {
			if !reflect.DeepEqual(a.ko.Spec.APIPassthrough.Extensions.CertificatePolicies, b.ko.Spec.APIPassthrough.Extensions.CertificatePolicies) {
				delta.Add("Spec.APIPassthrough.Extensions.CertificatePolicies", a.ko.Spec.APIPassthrough.Extensions.CertificatePolicies, b.ko.Spec.APIPassthrough.Extensions.CertificatePolicies)
			}
			if !reflect.DeepEqual(a.ko.Spec.APIPassthrough.Extensions.CustomExtensions, b.ko.Spec.APIPassthrough.Extensions.CustomExtensions) {
				delta.Add("Spec.APIPassthrough.Extensions.CustomExtensions", a.ko.Spec.APIPassthrough.Extensions.CustomExtensions, b.ko.Spec.APIPassthrough.Extensions.CustomExtensions)
			}
			if !reflect.DeepEqual(a.ko.Spec.APIPassthrough.Extensions.ExtendedKeyUsage, b.ko.Spec.APIPassthrough.Extensions.ExtendedKeyUsage) {
				delta.Add("Spec.APIPassthrough.Extensions.ExtendedKeyUsage", a.ko.Spec.APIPassthrough.Extensions.ExtendedKeyUsage, b.ko.Spec.APIPassthrough.Extensions.ExtendedKeyUsage)
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Extensions.KeyUsage, b.ko.Spec.APIPassthrough.Extensions.KeyUsage) {
				delta.Add("Spec.APIPassthrough.Extensions.KeyUsage", a.ko.Spec.APIPassthrough.Extensions.KeyUsage, b.ko.Spec.APIPassthrough.Extensions.KeyUsage)
			} else if a.ko.Spec.APIPassthrough.Extensions.KeyUsage != nil && b.ko.Spec.APIPassthrough.Extensions.KeyUsage != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Extensions.KeyUsage.CRLSign, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.CRLSign) {
					delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.CRLSign", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.CRLSign, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.CRLSign)
				} else if a.ko.Spec.APIPassthrough.Extensions.KeyUsage.CRLSign != nil && b.ko.Spec.APIPassthrough.Extensions.KeyUsage.CRLSign != nil {
					if *a.ko.Spec.APIPassthrough.Extensions.KeyUsage.CRLSign != *b.ko.Spec.APIPassthrough.Extensions.KeyUsage.CRLSign {
						delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.CRLSign", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.CRLSign, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.CRLSign)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment) {
					delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment)
				} else if a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment != nil && b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment != nil {
					if *a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment != *b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment {
						delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DataEncipherment)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly) {
					delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly)
				} else if a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly != nil && b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly != nil {
					if *a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly != *b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly {
						delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DecipherOnly)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature) {
					delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature)
				} else if a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature != nil && b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature != nil {
					if *a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature != *b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature {
						delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.DigitalSignature)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly) {
					delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly)
				} else if a.ko.Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly != nil && b.ko.Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly != nil {
					if *a.ko.Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly != *b.ko.Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly {
						delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.EncipherOnly)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement) {
					delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement)
				} else if a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement != nil && b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement != nil {
					if *a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement != *b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement {
						delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyAgreement)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign) {
					delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign)
				} else if a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign != nil && b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign != nil {
					if *a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign != *b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign {
						delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyCertSign)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment) {
					delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment)
				} else if a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment != nil && b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment != nil {
					if *a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment != *b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment {
						delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.KeyEncipherment)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation) {
					delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation)
				} else if a.ko.Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation != nil && b.ko.Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation != nil {
					if *a.ko.Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation != *b.ko.Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation {
						delta.Add("Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation", a.ko.Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation, b.ko.Spec.APIPassthrough.Extensions.KeyUsage.NonRepudiation)
					}
				}
			}
			if !reflect.DeepEqual(a.ko.Spec.APIPassthrough.Extensions.SubjectAlternativeNames, b.ko.Spec.APIPassthrough.Extensions.SubjectAlternativeNames) {
				delta.Add("Spec.APIPassthrough.Extensions.SubjectAlternativeNames", a.ko.Spec.APIPassthrough.Extensions.SubjectAlternativeNames, b.ko.Spec.APIPassthrough.Extensions.SubjectAlternativeNames)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject, b.ko.Spec.APIPassthrough.Subject) {
			delta.Add("Spec.APIPassthrough.Subject", a.ko.Spec.APIPassthrough.Subject, b.ko.Spec.APIPassthrough.Subject)
		} else if a.ko.Spec.APIPassthrough.Subject != nil && b.ko.Spec.APIPassthrough.Subject != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.CommonName, b.ko.Spec.APIPassthrough.Subject.CommonName) {
				delta.Add("Spec.APIPassthrough.Subject.CommonName", a.ko.Spec.APIPassthrough.Subject.CommonName, b.ko.Spec.APIPassthrough.Subject.CommonName)
			} else if a.ko.Spec.APIPassthrough.Subject.CommonName != nil && b.ko.Spec.APIPassthrough.Subject.CommonName != nil {
				if *a.ko.Spec.APIPassthrough.Subject.CommonName != *b.ko.Spec.APIPassthrough.Subject.CommonName {
					delta.Add("Spec.APIPassthrough.Subject.CommonName", a.ko.Spec.APIPassthrough.Subject.CommonName, b.ko.Spec.APIPassthrough.Subject.CommonName)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.Country, b.ko.Spec.APIPassthrough.Subject.Country) {
				delta.Add("Spec.APIPassthrough.Subject.Country", a.ko.Spec.APIPassthrough.Subject.Country, b.ko.Spec.APIPassthrough.Subject.Country)
			} else if a.ko.Spec.APIPassthrough.Subject.Country != nil && b.ko.Spec.APIPassthrough.Subject.Country != nil {
				if *a.ko.Spec.APIPassthrough.Subject.Country != *b.ko.Spec.APIPassthrough.Subject.Country {
					delta.Add("Spec.APIPassthrough.Subject.Country", a.ko.Spec.APIPassthrough.Subject.Country, b.ko.Spec.APIPassthrough.Subject.Country)
				}
			}
			if !reflect.DeepEqual(a.ko.Spec.APIPassthrough.Subject.CustomAttributes, b.ko.Spec.APIPassthrough.Subject.CustomAttributes) {
				delta.Add("Spec.APIPassthrough.Subject.CustomAttributes", a.ko.Spec.APIPassthrough.Subject.CustomAttributes, b.ko.Spec.APIPassthrough.Subject.CustomAttributes)
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.DistinguishedNameQualifier, b.ko.Spec.APIPassthrough.Subject.DistinguishedNameQualifier) {
				delta.Add("Spec.APIPassthrough.Subject.DistinguishedNameQualifier", a.ko.Spec.APIPassthrough.Subject.DistinguishedNameQualifier, b.ko.Spec.APIPassthrough.Subject.DistinguishedNameQualifier)
			} else if a.ko.Spec.APIPassthrough.Subject.DistinguishedNameQualifier != nil && b.ko.Spec.APIPassthrough.Subject.DistinguishedNameQualifier != nil {
				if *a.ko.Spec.APIPassthrough.Subject.DistinguishedNameQualifier != *b.ko.Spec.APIPassthrough.Subject.DistinguishedNameQualifier {
					delta.Add("Spec.APIPassthrough.Subject.DistinguishedNameQualifier", a.ko.Spec.APIPassthrough.Subject.DistinguishedNameQualifier, b.ko.Spec.APIPassthrough.Subject.DistinguishedNameQualifier)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.GenerationQualifier, b.ko.Spec.APIPassthrough.Subject.GenerationQualifier) {
				delta.Add("Spec.APIPassthrough.Subject.GenerationQualifier", a.ko.Spec.APIPassthrough.Subject.GenerationQualifier, b.ko.Spec.APIPassthrough.Subject.GenerationQualifier)
			} else if a.ko.Spec.APIPassthrough.Subject.GenerationQualifier != nil && b.ko.Spec.APIPassthrough.Subject.GenerationQualifier != nil {
				if *a.ko.Spec.APIPassthrough.Subject.GenerationQualifier != *b.ko.Spec.APIPassthrough.Subject.GenerationQualifier {
					delta.Add("Spec.APIPassthrough.Subject.GenerationQualifier", a.ko.Spec.APIPassthrough.Subject.GenerationQualifier, b.ko.Spec.APIPassthrough.Subject.GenerationQualifier)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.GivenName, b.ko.Spec.APIPassthrough.Subject.GivenName) {
				delta.Add("Spec.APIPassthrough.Subject.GivenName", a.ko.Spec.APIPassthrough.Subject.GivenName, b.ko.Spec.APIPassthrough.Subject.GivenName)
			} else if a.ko.Spec.APIPassthrough.Subject.GivenName != nil && b.ko.Spec.APIPassthrough.Subject.GivenName != nil {
				if *a.ko.Spec.APIPassthrough.Subject.GivenName != *b.ko.Spec.APIPassthrough.Subject.GivenName {
					delta.Add("Spec.APIPassthrough.Subject.GivenName", a.ko.Spec.APIPassthrough.Subject.GivenName, b.ko.Spec.APIPassthrough.Subject.GivenName)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.Initials, b.ko.Spec.APIPassthrough.Subject.Initials) {
				delta.Add("Spec.APIPassthrough.Subject.Initials", a.ko.Spec.APIPassthrough.Subject.Initials, b.ko.Spec.APIPassthrough.Subject.Initials)
			} else if a.ko.Spec.APIPassthrough.Subject.Initials != nil && b.ko.Spec.APIPassthrough.Subject.Initials != nil {
				if *a.ko.Spec.APIPassthrough.Subject.Initials != *b.ko.Spec.APIPassthrough.Subject.Initials {
					delta.Add("Spec.APIPassthrough.Subject.Initials", a.ko.Spec.APIPassthrough.Subject.Initials, b.ko.Spec.APIPassthrough.Subject.Initials)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.Locality, b.ko.Spec.APIPassthrough.Subject.Locality) {
				delta.Add("Spec.APIPassthrough.Subject.Locality", a.ko.Spec.APIPassthrough.Subject.Locality, b.ko.Spec.APIPassthrough.Subject.Locality)
			} else if a.ko.Spec.APIPassthrough.Subject.Locality != nil && b.ko.Spec.APIPassthrough.Subject.Locality != nil {
				if *a.ko.Spec.APIPassthrough.Subject.Locality != *b.ko.Spec.APIPassthrough.Subject.Locality {
					delta.Add("Spec.APIPassthrough.Subject.Locality", a.ko.Spec.APIPassthrough.Subject.Locality, b.ko.Spec.APIPassthrough.Subject.Locality)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.Organization, b.ko.Spec.APIPassthrough.Subject.Organization) {
				delta.Add("Spec.APIPassthrough.Subject.Organization", a.ko.Spec.APIPassthrough.Subject.Organization, b.ko.Spec.APIPassthrough.Subject.Organization)
			} else if a.ko.Spec.APIPassthrough.Subject.Organization != nil && b.ko.Spec.APIPassthrough.Subject.Organization != nil {
				if *a.ko.Spec.APIPassthrough.Subject.Organization != *b.ko.Spec.APIPassthrough.Subject.Organization {
					delta.Add("Spec.APIPassthrough.Subject.Organization", a.ko.Spec.APIPassthrough.Subject.Organization, b.ko.Spec.APIPassthrough.Subject.Organization)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.OrganizationalUnit, b.ko.Spec.APIPassthrough.Subject.OrganizationalUnit) {
				delta.Add("Spec.APIPassthrough.Subject.OrganizationalUnit", a.ko.Spec.APIPassthrough.Subject.OrganizationalUnit, b.ko.Spec.APIPassthrough.Subject.OrganizationalUnit)
			} else if a.ko.Spec.APIPassthrough.Subject.OrganizationalUnit != nil && b.ko.Spec.APIPassthrough.Subject.OrganizationalUnit != nil {
				if *a.ko.Spec.APIPassthrough.Subject.OrganizationalUnit != *b.ko.Spec.APIPassthrough.Subject.OrganizationalUnit {
					delta.Add("Spec.APIPassthrough.Subject.OrganizationalUnit", a.ko.Spec.APIPassthrough.Subject.OrganizationalUnit, b.ko.Spec.APIPassthrough.Subject.OrganizationalUnit)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.Pseudonym, b.ko.Spec.APIPassthrough.Subject.Pseudonym) {
				delta.Add("Spec.APIPassthrough.Subject.Pseudonym", a.ko.Spec.APIPassthrough.Subject.Pseudonym, b.ko.Spec.APIPassthrough.Subject.Pseudonym)
			} else if a.ko.Spec.APIPassthrough.Subject.Pseudonym != nil && b.ko.Spec.APIPassthrough.Subject.Pseudonym != nil {
				if *a.ko.Spec.APIPassthrough.Subject.Pseudonym != *b.ko.Spec.APIPassthrough.Subject.Pseudonym {
					delta.Add("Spec.APIPassthrough.Subject.Pseudonym", a.ko.Spec.APIPassthrough.Subject.Pseudonym, b.ko.Spec.APIPassthrough.Subject.Pseudonym)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.SerialNumber, b.ko.Spec.APIPassthrough.Subject.SerialNumber) {
				delta.Add("Spec.APIPassthrough.Subject.SerialNumber", a.ko.Spec.APIPassthrough.Subject.SerialNumber, b.ko.Spec.APIPassthrough.Subject.SerialNumber)
			} else if a.ko.Spec.APIPassthrough.Subject.SerialNumber != nil && b.ko.Spec.APIPassthrough.Subject.SerialNumber != nil {
				if *a.ko.Spec.APIPassthrough.Subject.SerialNumber != *b.ko.Spec.APIPassthrough.Subject.SerialNumber {
					delta.Add("Spec.APIPassthrough.Subject.SerialNumber", a.ko.Spec.APIPassthrough.Subject.SerialNumber, b.ko.Spec.APIPassthrough.Subject.SerialNumber)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.State, b.ko.Spec.APIPassthrough.Subject.State) {
				delta.Add("Spec.APIPassthrough.Subject.State", a.ko.Spec.APIPassthrough.Subject.State, b.ko.Spec.APIPassthrough.Subject.State)
			} else if a.ko.Spec.APIPassthrough.Subject.State != nil && b.ko.Spec.APIPassthrough.Subject.State != nil {
				if *a.ko.Spec.APIPassthrough.Subject.State != *b.ko.Spec.APIPassthrough.Subject.State {
					delta.Add("Spec.APIPassthrough.Subject.State", a.ko.Spec.APIPassthrough.Subject.State, b.ko.Spec.APIPassthrough.Subject.State)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.Surname, b.ko.Spec.APIPassthrough.Subject.Surname) {
				delta.Add("Spec.APIPassthrough.Subject.Surname", a.ko.Spec.APIPassthrough.Subject.Surname, b.ko.Spec.APIPassthrough.Subject.Surname)
			} else if a.ko.Spec.APIPassthrough.Subject.Surname != nil && b.ko.Spec.APIPassthrough.Subject.Surname != nil {
				if *a.ko.Spec.APIPassthrough.Subject.Surname != *b.ko.Spec.APIPassthrough.Subject.Surname {
					delta.Add("Spec.APIPassthrough.Subject.Surname", a.ko.Spec.APIPassthrough.Subject.Surname, b.ko.Spec.APIPassthrough.Subject.Surname)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.APIPassthrough.Subject.Title, b.ko.Spec.APIPassthrough.Subject.Title) {
				delta.Add("Spec.APIPassthrough.Subject.Title", a.ko.Spec.APIPassthrough.Subject.Title, b.ko.Spec.APIPassthrough.Subject.Title)
			} else if a.ko.Spec.APIPassthrough.Subject.Title != nil && b.ko.Spec.APIPassthrough.Subject.Title != nil {
				if *a.ko.Spec.APIPassthrough.Subject.Title != *b.ko.Spec.APIPassthrough.Subject.Title {
					delta.Add("Spec.APIPassthrough.Subject.Title", a.ko.Spec.APIPassthrough.Subject.Title, b.ko.Spec.APIPassthrough.Subject.Title)
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CertificateAuthorityARN, b.ko.Spec.CertificateAuthorityARN) {
		delta.Add("Spec.CertificateAuthorityARN", a.ko.Spec.CertificateAuthorityARN, b.ko.Spec.CertificateAuthorityARN)
	} else if a.ko.Spec.CertificateAuthorityARN != nil && b.ko.Spec.CertificateAuthorityARN != nil {
		if *a.ko.Spec.CertificateAuthorityARN != *b.ko.Spec.CertificateAuthorityARN {
			delta.Add("Spec.CertificateAuthorityARN", a.ko.Spec.CertificateAuthorityARN, b.ko.Spec.CertificateAuthorityARN)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.CertificateAuthorityRef, b.ko.Spec.CertificateAuthorityRef) {
		delta.Add("Spec.CertificateAuthorityRef", a.ko.Spec.CertificateAuthorityRef, b.ko.Spec.CertificateAuthorityRef)
	}
	if !bytes.Equal(a.ko.Spec.CSR, b.ko.Spec.CSR) {
		delta.Add("Spec.CSR", a.ko.Spec.CSR, b.ko.Spec.CSR)
	}
	if !reflect.DeepEqual(a.ko.Spec.CSRRef, b.ko.Spec.CSRRef) {
		delta.Add("Spec.CSRRef", a.ko.Spec.CSRRef, b.ko.Spec.CSRRef)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.IdempotencyToken, b.ko.Spec.IdempotencyToken) {
		delta.Add("Spec.IdempotencyToken", a.ko.Spec.IdempotencyToken, b.ko.Spec.IdempotencyToken)
	} else if a.ko.Spec.IdempotencyToken != nil && b.ko.Spec.IdempotencyToken != nil {
		if *a.ko.Spec.IdempotencyToken != *b.ko.Spec.IdempotencyToken {
			delta.Add("Spec.IdempotencyToken", a.ko.Spec.IdempotencyToken, b.ko.Spec.IdempotencyToken)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.SigningAlgorithm, b.ko.Spec.SigningAlgorithm) {
		delta.Add("Spec.SigningAlgorithm", a.ko.Spec.SigningAlgorithm, b.ko.Spec.SigningAlgorithm)
	} else if a.ko.Spec.SigningAlgorithm != nil && b.ko.Spec.SigningAlgorithm != nil {
		if *a.ko.Spec.SigningAlgorithm != *b.ko.Spec.SigningAlgorithm {
			delta.Add("Spec.SigningAlgorithm", a.ko.Spec.SigningAlgorithm, b.ko.Spec.SigningAlgorithm)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TemplateARN, b.ko.Spec.TemplateARN) {
		delta.Add("Spec.TemplateARN", a.ko.Spec.TemplateARN, b.ko.Spec.TemplateARN)
	} else if a.ko.Spec.TemplateARN != nil && b.ko.Spec.TemplateARN != nil {
		if *a.ko.Spec.TemplateARN != *b.ko.Spec.TemplateARN {
			delta.Add("Spec.TemplateARN", a.ko.Spec.TemplateARN, b.ko.Spec.TemplateARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Validity, b.ko.Spec.Validity) {
		delta.Add("Spec.Validity", a.ko.Spec.Validity, b.ko.Spec.Validity)
	} else if a.ko.Spec.Validity != nil && b.ko.Spec.Validity != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.Validity.Type, b.ko.Spec.Validity.Type) {
			delta.Add("Spec.Validity.Type", a.ko.Spec.Validity.Type, b.ko.Spec.Validity.Type)
		} else if a.ko.Spec.Validity.Type != nil && b.ko.Spec.Validity.Type != nil {
			if *a.ko.Spec.Validity.Type != *b.ko.Spec.Validity.Type {
				delta.Add("Spec.Validity.Type", a.ko.Spec.Validity.Type, b.ko.Spec.Validity.Type)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.Validity.Value, b.ko.Spec.Validity.Value) {
			delta.Add("Spec.Validity.Value", a.ko.Spec.Validity.Value, b.ko.Spec.Validity.Value)
		} else if a.ko.Spec.Validity.Value != nil && b.ko.Spec.Validity.Value != nil {
			if *a.ko.Spec.Validity.Value != *b.ko.Spec.Validity.Value {
				delta.Add("Spec.Validity.Value", a.ko.Spec.Validity.Value, b.ko.Spec.Validity.Value)
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ValidityNotBefore, b.ko.Spec.ValidityNotBefore) {
		delta.Add("Spec.ValidityNotBefore", a.ko.Spec.ValidityNotBefore, b.ko.Spec.ValidityNotBefore)
	} else if a.ko.Spec.ValidityNotBefore != nil && b.ko.Spec.ValidityNotBefore != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.ValidityNotBefore.Type, b.ko.Spec.ValidityNotBefore.Type) {
			delta.Add("Spec.ValidityNotBefore.Type", a.ko.Spec.ValidityNotBefore.Type, b.ko.Spec.ValidityNotBefore.Type)
		} else if a.ko.Spec.ValidityNotBefore.Type != nil && b.ko.Spec.ValidityNotBefore.Type != nil {
			if *a.ko.Spec.ValidityNotBefore.Type != *b.ko.Spec.ValidityNotBefore.Type {
				delta.Add("Spec.ValidityNotBefore.Type", a.ko.Spec.ValidityNotBefore.Type, b.ko.Spec.ValidityNotBefore.Type)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.ValidityNotBefore.Value, b.ko.Spec.ValidityNotBefore.Value) {
			delta.Add("Spec.ValidityNotBefore.Value", a.ko.Spec.ValidityNotBefore.Value, b.ko.Spec.ValidityNotBefore.Value)
		} else if a.ko.Spec.ValidityNotBefore.Value != nil && b.ko.Spec.ValidityNotBefore.Value != nil {
			if *a.ko.Spec.ValidityNotBefore.Value != *b.ko.Spec.ValidityNotBefore.Value {
				delta.Add("Spec.ValidityNotBefore.Value", a.ko.Spec.ValidityNotBefore.Value, b.ko.Spec.ValidityNotBefore.Value)
			}
		}
	}

	return delta
}