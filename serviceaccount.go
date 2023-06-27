// Copyright 2023 Cover Whale Insurance Solutions Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kopts

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceAccount is a Kubernetes service account
type ServiceAccount struct {
	corev1.ServiceAccount
}

type ServiceAccountOpt func(*ServiceAccount)

// NewSeviceAccount returns a service account with the provided name and options
func NewServiceAccount(name string, opts ...ServiceAccountOpt) ServiceAccount {
	sa := ServiceAccount{
		ServiceAccount: corev1.ServiceAccount{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ServiceAccount",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		},
	}

	for _, v := range opts {
		v(&sa)
	}

	return sa
}

// ServiceAccountNamespace sets the namespace for the service account
func ServiceAccountNamespace(n string) ServiceAccountOpt {
	return func(s *ServiceAccount) {
		s.Namespace = n
	}
}

// ServiceAccountImagePullSecret sets an image pull secret for the service account
func ServiceAccountImagePullSecret(s string) ServiceAccountOpt {
	return serviceAccountImagePullSecrets()
}

// ServiceAccountImagePullSecrets sets multiple image pull secrets for the service account
func ServiceAccountImagePullSecrets(s []string) ServiceAccountOpt {
	return serviceAccountImagePullSecrets(s...)
}

func serviceAccountImagePullSecrets(s ...string) ServiceAccountOpt {
	return func(sa *ServiceAccount) {
		for _, v := range s {
			sa.ImagePullSecrets = append(sa.ImagePullSecrets, corev1.LocalObjectReference{
				Name: v,
			},
			)
		}
	}
}

// ServiceAccountAutoMountToken sets the automounttoken for the service account
func ServiceAccountAutoMountToken(b bool) ServiceAccountOpt {
	autoMount := b
	return func(s *ServiceAccount) {
		s.AutomountServiceAccountToken = &autoMount
	}
}
