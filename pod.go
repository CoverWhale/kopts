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

type PodOpt func(*PodSpec)

// Podspec holds information for a Kubernetes pod spec
type PodSpec struct {
	Name      string
	Namespace string
	Image     string
	Spec      corev1.PodTemplateSpec
}

// Returns a pod spec with the given name and options
func NewPodSpec(name string, opts ...PodOpt) PodSpec {
	pod := PodSpec{
		Spec: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{},
			},
		},
	}

	for _, v := range opts {
		v(&pod)
	}

	return pod
}

// Set signle pod label
func PodLabel(key, value string) PodOpt {
	return func(p *PodSpec) {
		p.Spec.ObjectMeta.Labels = map[string]string{
			key: value,
		}
	}
}

// Add a pod container
func PodContainer(c Container) PodOpt {
	return func(p *PodSpec) {
		p.Spec.Spec.Containers = append(p.Spec.Spec.Containers, c.Container)
	}
}

// Add a pod init container
func PodInitContainer(c Container) PodOpt {
	return func(p *PodSpec) {
		p.Spec.Spec.InitContainers = append(p.Spec.Spec.InitContainers, c.Container)
	}
}
