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

	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func (rm *resourceManager) writeCertificateToSecret(
	ctx context.Context,
	certificate *string,
	annotations map[string]string,
) (err error) {

	namespace, found := annotations["ack.aws.k8s.io/secret-namespace"]
	if !found {
		namespace = "default"
	}

	name, found := annotations["ack.aws.k8s.io/secret-name"]
	if !found {
		return ackerr.SecretNotFound
	}

	key, found := annotations["ack.aws.k8s.io/secret-key"]
	if !found {
		return ackerr.SecretNotFound
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	secretsClient := clientset.CoreV1().Secrets(namespace)

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
