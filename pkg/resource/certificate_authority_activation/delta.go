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

	if ackcompare.HasNilDifference(a.ko.Spec.Certificate, b.ko.Spec.Certificate) {
		delta.Add("Spec.Certificate", a.ko.Spec.Certificate, b.ko.Spec.Certificate)
	} else if a.ko.Spec.Certificate != nil && b.ko.Spec.Certificate != nil {
		if *a.ko.Spec.Certificate != *b.ko.Spec.Certificate {
			delta.Add("Spec.Certificate", a.ko.Spec.Certificate, b.ko.Spec.Certificate)
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
	if ackcompare.HasNilDifference(a.ko.Spec.CertificateChain, b.ko.Spec.CertificateChain) {
		delta.Add("Spec.CertificateChain", a.ko.Spec.CertificateChain, b.ko.Spec.CertificateChain)
	} else if a.ko.Spec.CertificateChain != nil && b.ko.Spec.CertificateChain != nil {
		if *a.ko.Spec.CertificateChain != *b.ko.Spec.CertificateChain {
			delta.Add("Spec.CertificateChain", a.ko.Spec.CertificateChain, b.ko.Spec.CertificateChain)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Status, b.ko.Spec.Status) {
		delta.Add("Spec.Status", a.ko.Spec.Status, b.ko.Spec.Status)
	} else if a.ko.Spec.Status != nil && b.ko.Spec.Status != nil {
		if *a.ko.Spec.Status != *b.ko.Spec.Status {
			delta.Add("Spec.Status", a.ko.Spec.Status, b.ko.Spec.Status)
		}
	}

	return delta
}
