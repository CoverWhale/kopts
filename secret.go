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

// Secret holds a Kubernetes secret
type Secret struct {
	corev1.Secret
}

type SecretOpt func(*Secret)

// NewSecret returns a secret with the given name and options
func NewSecret(name string, opts ...SecretOpt) Secret {
	s := Secret{
		Secret: corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		},
	}

	for _, v := range opts {
		v(&s)
	}

	return s
}

// Set secret namespace
func SecretNamespace(n string) SecretOpt {
	return func(s *Secret) {
		s.Namespace = n
	}
}

// Set secret data
func SecretData(key string, value []byte) SecretOpt {
	return func(s *Secret) {
		s.Data = map[string][]byte{
			key: value,
		}
	}
}
