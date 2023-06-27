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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// MarshalYaml returns the YAML for an API object
func MarshalYaml(i interface{}) (string, error) {
	o, err := yaml.Marshal(i)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("---\n%s\n", o), nil
}

func addAnnotation(key, value string, m *metav1.ObjectMeta) {
	if m.Annotations == nil {
		m.Annotations = make(map[string]string)
	}
	m.Annotations[key] = value
}

func setNamespace(n string, m *metav1.ObjectMeta) {
	m.Namespace = n
}

func newObjectMeta(name string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:        name,
		Labels:      make(map[string]string),
		Annotations: make(map[string]string),
	}
}

func addLabel(key, value string, m *metav1.ObjectMeta) {
	if m.Labels == nil {
		m.Labels = make(map[string]string)
	}

	m.Labels[key] = value
}
