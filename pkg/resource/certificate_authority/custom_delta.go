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
	desired *resource,
	latest *resource,
) {
	if ackcompare.IsNil(desired.ko.Spec.UsageMode) && ackcompare.IsNotNil(latest.ko.Spec.UsageMode) {
		desired.ko.Spec.UsageMode = latest.ko.Spec.UsageMode
	}

	if ackcompare.IsNil(desired.ko.Spec.KeyStorageSecurityStandard) && ackcompare.IsNotNil(latest.ko.Spec.KeyStorageSecurityStandard) {
		desired.ko.Spec.KeyStorageSecurityStandard = latest.ko.Spec.KeyStorageSecurityStandard
	}

	if ackcompare.IsNil(desired.ko.Spec.RevocationConfiguration) && ackcompare.IsNotNil(latest.ko.Spec.RevocationConfiguration) {
		desired.ko.Spec.RevocationConfiguration = &svcapitypes.RevocationConfiguration{}
	}

	if ackcompare.IsNotNil(desired.ko.Spec.RevocationConfiguration) && ackcompare.IsNotNil(latest.ko.Spec.RevocationConfiguration) {
		if ackcompare.IsNil(desired.ko.Spec.RevocationConfiguration.CRLConfiguration) && ackcompare.IsNotNil(latest.ko.Spec.RevocationConfiguration.CRLConfiguration) {
			desired.ko.Spec.RevocationConfiguration.CRLConfiguration = latest.ko.Spec.RevocationConfiguration.CRLConfiguration
		}
		if ackcompare.IsNil(desired.ko.Spec.RevocationConfiguration.OCSPConfiguration) && ackcompare.IsNotNil(latest.ko.Spec.RevocationConfiguration.OCSPConfiguration) {
			desired.ko.Spec.RevocationConfiguration.OCSPConfiguration = latest.ko.Spec.RevocationConfiguration.OCSPConfiguration
		}
	}
}
