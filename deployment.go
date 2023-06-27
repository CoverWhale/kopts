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

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ErrNameRequired = fmt.Errorf("name is required")

// Deployment holds a Kubernetes deployment
type Deployment struct {
	appsv1.Deployment
}

type DeploymentOpt func(*Deployment)

// NewDeployment returns a deployment with the given name and options
func NewDeployment(name string, depOpts ...DeploymentOpt) *Deployment {
	dep := &Deployment{
		appsv1.Deployment{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Deployment",
				APIVersion: "apps/v1",
			},
			ObjectMeta: newObjectMeta(name),
			Spec: appsv1.DeploymentSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: make(map[string]string),
				},
				Template: corev1.PodTemplateSpec{},
			},
		},
	}

	for _, v := range depOpts {
		v(dep)
	}

	return dep

}

// Set deployment namespace
func DeploymentNamespace(n string) DeploymentOpt {
	return func(d *Deployment) {
		setNamespace(n, &d.ObjectMeta)
	}
}

// Add deployment selector
func DeploymentSelector(key, value string) DeploymentOpt {
	return func(d *Deployment) {
		metav1.AddLabelToSelector(d.Spec.Selector, key, value)
	}
}

// Add single deployment label
func DeploymentLabel(key, value string) DeploymentOpt {
	return func(d *Deployment) {
		addLabel(key, value, &d.ObjectMeta)
	}
}

// Add multiple deployment labels
func DeploymentLabels(labels map[string]string) DeploymentOpt {
	return func(d *Deployment) {
		for k, v := range labels {
			addLabel(k, v, &d.ObjectMeta)
		}
	}
}

// Set deployment pod spec
func DeploymentPodSpec(p PodSpec) DeploymentOpt {
	return func(d *Deployment) {
		d.Spec.Template = p.Spec
	}
}

// Set deployment replicas
func DeploymentReplicas(r int) DeploymentOpt {
	replicas := int32(r)
	return func(d *Deployment) {
		d.Spec.Replicas = &replicas
	}
}
