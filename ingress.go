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
	"fmt"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Ingress holds a Kubernetes ingress
type Ingress struct {
	networkingv1.Ingress
}

type IngressOpt func(*Ingress)

// NewIngress returns an ingress with the given name and options
func NewIngress(name string, opts ...IngressOpt) Ingress {
	i := Ingress{
		Ingress: networkingv1.Ingress{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Ingress",
				APIVersion: "networking.k8s.io/v1",
			},
			ObjectMeta: newObjectMeta(name),
		},
	}

	for _, v := range opts {
		v(&i)
	}

	return i
}

// Set ingress namespace
func IngressNamespace(n string) IngressOpt {
	return func(i *Ingress) {
		i.Ingress.ObjectMeta.Namespace = n
	}
}

// Set ingress class
func IngressClass(c string) IngressOpt {
	return func(i *Ingress) {
		i.Ingress.Spec.IngressClassName = &c
	}
}

// Rule holds an ingress rule
type Rule struct {
	Host        string
	Paths       []Path
	LetsEncrypt bool
	TLS         bool
}

// Path holds an ingress path
type Path struct {
	Name    string
	Service string
	Port    int
	Type    networkingv1.PathType
}

// Append a rule to paths
func IngressRule(r Rule) IngressOpt {
	var paths []networkingv1.HTTPIngressPath
	for _, v := range r.Paths {
		paths = append(paths, networkingv1.HTTPIngressPath{
			Path:     v.Name,
			PathType: &v.Type,
			Backend: networkingv1.IngressBackend{
				Service: &networkingv1.IngressServiceBackend{
					Name: v.Service,
					Port: networkingv1.ServiceBackendPort{
						Number: int32(v.Port),
					},
				},
			},
		})
	}

	return func(i *Ingress) {
		if r.LetsEncrypt {
			addAnnotation("cert-manager.io/cluster-issuer", "letsencrypt-prod", &i.ObjectMeta)
			i.Spec.TLS = append(i.Spec.TLS, networkingv1.IngressTLS{
				Hosts:      []string{r.Host},
				SecretName: fmt.Sprintf("%s-tls", i.Ingress.ObjectMeta.Name),
			})
		}
		i.Ingress.Spec.Rules = append(i.Ingress.Spec.Rules, networkingv1.IngressRule{
			Host: r.Host,
			IngressRuleValue: networkingv1.IngressRuleValue{
				HTTP: &networkingv1.HTTPIngressRuleValue{
					Paths: paths,
				},
			},
		})
	}
}
