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

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
)

func (rm *resourceManager) writeCertificateToSecret(
	ctx context.Context,
	certificate,
	resourceNamespace string,
	secretKeyReference *ackv1alpha1.SecretKeyReference,
) error {

	namespace := resourceNamespace
	if secretKeyReference.SecretReference.Namespace != "" {
		namespace = secretKeyReference.SecretReference.Namespace
	}

	err := rm.rr.WriteToSecret(ctx, certificate, namespace, secretKeyReference.SecretReference.Name, secretKeyReference.Key)
	rm.metrics.RecordAPICall("PATCH", "writeCertificateToSecret", err)
	if err != nil {
		return err
	}

	return nil
}
