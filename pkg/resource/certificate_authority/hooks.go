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
	"context"

	svcsdk "github.com/aws/aws-sdk-go/service/acmpca"
)

func (rm *resourceManager) getCertificateSigningRequest(
	ctx context.Context,
	resourceARN string,
) (*string, error) {
	input := &svcsdk.GetCertificateAuthorityCsrInput{}
	input.CertificateAuthorityArn = &resourceARN
	resp, err := rm.sdkapi.GetCertificateAuthorityCsrWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetCertificateAuthorityCsr", err)
	if err != nil {
		return nil, err
	}
	csr := resp.Csr
	return csr, err
}
