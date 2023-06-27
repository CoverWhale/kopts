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

type ConfigMap struct {
	corev1.ConfigMap
}

type ConfigMapOpt func(*ConfigMap)

// Returns a configmap with the given name and options
func NewConfigMap(name string, opts ...ConfigMapOpt) ConfigMap {
	c := ConfigMap{
		ConfigMap: corev1.ConfigMap{
			TypeMeta: metav1.TypeMeta{
				Kind:       "ConfigMap",
				APIVersion: "v1",
			},
			ObjectMeta: newObjectMeta(name),
			Data:       make(map[string]string),
			BinaryData: make(map[string][]byte),
		},
	}

	for _, v := range opts {
		v(&c)
	}

	return c
}

// Set configmap namespace
func ConfigMapNamespace(n string) ConfigMapOpt {
	return func(c *ConfigMap) {
		setNamespace(n, &c.ObjectMeta)
	}
}

// Set if configmap is immutable
func ConfigMapImmutable(b bool) ConfigMapOpt {
	return func(c *ConfigMap) {
		c.ConfigMap.Immutable = &b
	}
}

// Set singular configmap kv pair
func ConfigMapData(key, value string) ConfigMapOpt {
	return func(c *ConfigMap) {
		c.ConfigMap.Data[key] = value
	}
}

// Set multiple configmap kv pairs
func ConfigMapDataMap(data map[string]string) ConfigMapOpt {
	return func(c *ConfigMap) {
		for k, v := range data {
			c.ConfigMap.Data[k] = v
		}
	}
}

// Set singular binary kv data
func ConfigMapBinaryData(key string, value []byte) ConfigMapOpt {
	return func(c *ConfigMap) {
		c.ConfigMap.BinaryData[key] = value
	}
}

// Set multiple binary kv data
func ConfigMapBinaryDataMap(data map[string][]byte) ConfigMapOpt {
	return func(c *ConfigMap) {
		for k, v := range data {
			c.ConfigMap.BinaryData[k] = v
		}
	}
}
