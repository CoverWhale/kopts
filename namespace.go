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

// Namespace holds a Kubernetes namespace
type Namespace struct {
	corev1.Namespace
}

type NamespaceOpt func(*Namespace)

// NewNamespace returns a namespace with the given name and options
func NewNamespace(name string, opts ...NamespaceOpt) Namespace {
	ns := Namespace{
		corev1.Namespace{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Namespace",
				APIVersion: "v1",
			},
			ObjectMeta: newObjectMeta(name),
		},
	}

	for _, v := range opts {
		v(&ns)
	}

	return ns
}

// Set single annotation
func NamespaceAnnotation(key, value string) NamespaceOpt {
	return func(n *Namespace) {
		addAnnotation(key, value, &n.ObjectMeta)
	}
}

// Set multiple annotations
func NamespaceAnnotations(annotations map[string]string) NamespaceOpt {
	return func(n *Namespace) {
		for k, v := range annotations {
			addAnnotation(k, v, &n.ObjectMeta)
		}
	}
}
