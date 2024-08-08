// Copyright 2024 Cover Whale Insurance Solutions Inc.
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

import corev1 "k8s.io/api/core/v1"

type TolerationOperator = corev1.TolerationOperator
type TaintEffect = corev1.TaintEffect

type Toleration struct {
	corev1.Toleration
	Key               string
	Value             string
	TolerationSeconds int
	Operator          corev1.TolerationOperator
	Effect            corev1.TaintEffect
}
