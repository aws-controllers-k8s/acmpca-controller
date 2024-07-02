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

package certificate_authority

import (
	svcapitypes "github.com/aws-controllers-k8s/acmpca-controller/apis/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

func customSetDefaults(
	a *resource,
	b *resource,
) {
	if ackcompare.IsNil(a.ko.Spec.UsageMode) && ackcompare.IsNotNil(b.ko.Spec.UsageMode) {
		a.ko.Spec.UsageMode = b.ko.Spec.UsageMode
	}

	if ackcompare.IsNil(a.ko.Spec.KeyStorageSecurityStandard) && ackcompare.IsNotNil(b.ko.Spec.KeyStorageSecurityStandard) {
		a.ko.Spec.KeyStorageSecurityStandard = b.ko.Spec.KeyStorageSecurityStandard
	}

	if ackcompare.IsNil(a.ko.Spec.RevocationConfiguration) && ackcompare.IsNotNil(b.ko.Spec.RevocationConfiguration) {
		a.ko.Spec.RevocationConfiguration = &svcapitypes.RevocationConfiguration{}
	}

	if ackcompare.IsNotNil(a.ko.Spec.RevocationConfiguration) && ackcompare.IsNotNil(b.ko.Spec.RevocationConfiguration) {
		if ackcompare.IsNil(a.ko.Spec.RevocationConfiguration.CRLConfiguration) && ackcompare.IsNotNil(b.ko.Spec.RevocationConfiguration.CRLConfiguration) {
			a.ko.Spec.RevocationConfiguration.CRLConfiguration = b.ko.Spec.RevocationConfiguration.CRLConfiguration
		}
		if ackcompare.IsNil(a.ko.Spec.RevocationConfiguration.OCSPConfiguration) && ackcompare.IsNotNil(b.ko.Spec.RevocationConfiguration.OCSPConfiguration) {
			a.ko.Spec.RevocationConfiguration.OCSPConfiguration = b.ko.Spec.RevocationConfiguration.OCSPConfiguration
		}
	}
}
