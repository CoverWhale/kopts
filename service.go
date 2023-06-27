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
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Service holds a kubernetes service
type Service struct {
	corev1.Service
}

type ServiceOpt func(*Service)

// NewService returns a service with the given name and options
func NewService(name string, opts ...ServiceOpt) Service {
	service := Service{
		Service: corev1.Service{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Service",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
			Spec: corev1.ServiceSpec{
				Selector: make(map[string]string),
			},
		},
	}

	for _, v := range opts {
		v(&service)
	}

	return service
}

// Set service namespace
func ServiceNamespace(n string) ServiceOpt {
	return func(s *Service) {
		s.ObjectMeta.Namespace = n
	}
}

// Add service port
func ServicePort(port, targetPort int) ServiceOpt {
	return func(s *Service) {
		s.Spec.Ports = append(s.Spec.Ports, corev1.ServicePort{
			Port:       int32(port),
			TargetPort: intstr.FromInt(targetPort),
		})
	}
}

// Set service selector
func ServiceSelector(key, value string) ServiceOpt {
	return func(s *Service) {
		s.Spec.Selector[key] = value
	}
}
