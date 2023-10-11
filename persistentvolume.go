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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PersistentVolume holds a Kubernetes persistent volume
type PersistentVolume struct {
	corev1.PersistentVolume
}

type PersistentVolumeOpt func(*PersistentVolume)

// NewPersistentVolume returns a persistent volume with the given name and options
func NewPersistentVolume(name string, opts ...PersistentVolumeOpt) PersistentVolume {
	pv := PersistentVolume{
		PersistentVolume: corev1.PersistentVolume{
			TypeMeta: metav1.TypeMeta{
				Kind:       "PersistentVolume",
				APIVersion: "v1",
			},
			ObjectMeta: newObjectMeta(name),
			Spec:       corev1.PersistentVolumeSpec{},
		},
	}

	for _, v := range opts {
		v(&pv)
	}

	return pv
}

// Set volume capacity
func PersistentvolumeCapacity(capacity resource.Quantity) PersistentVolumeOpt {
	return func(pv *PersistentVolume) {
		if pv.Spec.Capacity == nil {
			pv.Spec.Capacity = corev1.ResourceList{
				"capacity": capacity,
			}
		}
	}
}

// Set the volume host path
func PersistentVolumeHostPath(path string, pathType corev1.HostPathType) PersistentVolumeOpt {
	return func(pv *PersistentVolume) {
		pv.Spec.HostPath = &corev1.HostPathVolumeSource{
			Path: path,
			Type: &pathType,
		}
	}
}

// Set the local volume path
func PersistentVolumeLocal(path string, fsType string) PersistentVolumeOpt {
	return func(pv *PersistentVolume) {
		pv.Spec.Local = &corev1.LocalVolumeSource{
			Path:   path,
			FSType: &fsType,
		}
	}
}
