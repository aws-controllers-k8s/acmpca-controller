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
	"encoding/json"

	client "github.com/aws-controllers-k8s/acmpca-controller/pkg/client"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (rm *resourceManager) writeCertificateToSecret(
	ctx context.Context,
	certificate *string,
	r *resource,
) (err error) {

	annotations := r.ko.ObjectMeta.GetAnnotations()

	namespace, found := annotations["acmpca.services.k8s.aws/output-secret-namespace"]
	if !found {
		namespace = r.MetaObject().GetNamespace()
	}

	name, found := annotations["acmpca.services.k8s.aws/output-secret-name"]
	if !found {
		return ackerr.SecretNotFound
	}

	key, found := annotations["acmpca.services.k8s.aws/output-secret-key"]
	if !found {
		key = "certificate"
	}

	secretsClient, err := client.GetSecretsClient(namespace)
	if err != nil {
		return err
	}

	secret := corev1.Secret{
		Data: map[string][]byte{
			key: []byte(*certificate),
		},
	}

	payloadBytes, err := json.Marshal(secret)
	if err != nil {
		return err
	}

	_, err = secretsClient.Patch(ctx, name, types.StrategicMergePatchType, payloadBytes, metav1.PatchOptions{})
	rm.metrics.RecordAPICall("PATCH", "writeCertificateToSecret", err)
	if err != nil {
		return err
	}

	return nil
}
