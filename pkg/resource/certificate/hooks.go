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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func (rm *resourceManager) writeCertificateToSecret(
	ctx context.Context,
	certificate *string,
	annotations map[string]string,
) (err error) {

	namespace, in_map := annotations["ack.aws.k8s.io/secret-namespace"]
	if !in_map {
		namespace = "default"
	}

	name, in_map := annotations["ack.aws.k8s.io/secret-name"]
	if !in_map {
		return ackerr.SecretNotFound
	}

	key, in_map := annotations["ack.aws.k8s.io/secret-key"]
	if !in_map {
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

	// Retrieve secret
	secret, err := secretsClient.Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return ackerr.SecretNotFound
	}

	// Update field
	if secret.Data == nil {
		secret.Data = map[string][]byte{}
	}
	secret.Data[key] = []byte(*certificate)

	_, err = secretsClient.Update(ctx, secret, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}
