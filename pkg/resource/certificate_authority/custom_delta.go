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
	"github.com/aws/aws-sdk-go/aws"
)

func customSetDefaults(
	a *resource,
	b *resource,
) {
	// Default value of UsageMode is GENERAL_PURPOSE
	defaultUsageMode := aws.String("GENERAL_PURPOSE")

	if ackcompare.IsNil(a.ko.Spec.UsageMode) && ackcompare.IsNotNil(b.ko.Spec.UsageMode) {
		a.ko.Spec.UsageMode = defaultUsageMode
	}

	// Default value of KeyStorageSecurityStandard is FIPS_140_2_LEVEL_3_OR_HIGHER
	defaultKeyStorageSecurityStandard := aws.String("FIPS_140_2_LEVEL_3_OR_HIGHER")

	if ackcompare.IsNil(a.ko.Spec.KeyStorageSecurityStandard) && ackcompare.IsNotNil(b.ko.Spec.KeyStorageSecurityStandard) {
		a.ko.Spec.KeyStorageSecurityStandard = defaultKeyStorageSecurityStandard
	}

	// Default value of RevocationConfiguration.CrlConfiguration.Enabled is false
	defaultCrlConfigurationEnabled := aws.Bool(false)

	// Default value of RevocationConfiguration.OcspConfiguration.Enabled is false
	defaultOcspConfigurationEnabled := aws.Bool(false)

	if ackcompare.IsNil(a.ko.Spec.RevocationConfiguration) && ackcompare.IsNotNil(b.ko.Spec.RevocationConfiguration) {
		revocationConfiguration := &svcapitypes.RevocationConfiguration{}
		revocationConfiguration.CRLConfiguration = &svcapitypes.CRLConfiguration{}
		revocationConfiguration.CRLConfiguration.Enabled = defaultCrlConfigurationEnabled

		revocationConfiguration.OCSPConfiguration = &svcapitypes.OCSPConfiguration{}
		revocationConfiguration.OCSPConfiguration.Enabled = defaultOcspConfigurationEnabled

		a.ko.Spec.RevocationConfiguration = revocationConfiguration
	}
}
