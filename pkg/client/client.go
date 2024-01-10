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

package client

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

var (
	dynClient *dynamic.DynamicClient
)

func GetDynamicClient() (client *dynamic.DynamicClient, err error) {
	if dynClient == nil {
		config, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}

		dynClient, err = dynamic.NewForConfig(config)
		if err != nil {
			return nil, err
		}
	}
	return dynClient, nil
}

func GetSecretsClient(
	namespace string,
) (client v1.SecretInterface, err error) {

	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	secretsClient := clientset.CoreV1().Secrets(namespace)
	return secretsClient, nil
}
