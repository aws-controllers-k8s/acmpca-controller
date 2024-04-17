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

package certificate

import (
	"context"

	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	svcsdk "github.com/aws/aws-sdk-go/service/acmpca"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (rm *resourceManager) writeCertificateToSecret(
	ctx context.Context,
	resourceARN string,
	caARN string,
	objectMeta metav1.ObjectMeta,
) (err error) {

	input := &svcsdk.GetCertificateInput{}
	input.CertificateArn = &resourceARN
	input.CertificateAuthorityArn = &caARN
	resp, err := rm.sdkapi.GetCertificateWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetCertificate", err)
	if err != nil {
		return err
	}

	annotations := objectMeta.GetAnnotations()

	namespace, found := annotations["acmpca.services.k8s.aws/certificate-secret-namespace"]
	if !found {
		namespace = objectMeta.GetNamespace()
	}

	name, found := annotations["acmpca.services.k8s.aws/certificate-secret-name"]
	if !found {
		return ackerr.SecretNotFound
	}

	key, found := annotations["acmpca.services.k8s.aws/certificate-secret-key"]
	if !found {
		key = "certificate"
	}

	err = rm.rr.WriteToSecret(ctx, *resp.Certificate, namespace, name, key)
	rm.metrics.RecordAPICall("PATCH", "writeCertificateToSecret", err)
	if err != nil {
		return err
	}

	return nil
}
