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
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RoleBinding is a Kubernetes role binding
type RoleBinding struct {
	rbacv1.RoleBinding
}

type RoleBindingOpt func(*RoleBinding)

// NewRoleBinding returns a role binding with the given name and options
func NewRoleBinding(name string, opts ...RoleBindingOpt) RoleBinding {
	rb := RoleBinding{
		RoleBinding: rbacv1.RoleBinding{
			TypeMeta: metav1.TypeMeta{
				Kind:       "RoleBinding",
				APIVersion: "rbac.authorization.k8s.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		},
	}

	for _, v := range opts {
		v(&rb)
	}

	return rb
}

// Set role binding namespace
func RoleBindingNamespace(n string) RoleBindingOpt {
	return func(r *RoleBinding) {
		r.Namespace = n
	}
}

// Set role binding subject
func RoleBindingSubject(s rbacv1.Subject) RoleBindingOpt {
	return roleBindingSubject(s)
}

// Set multiple role binding subjects
func RoleBindingSubjects(s []rbacv1.Subject) RoleBindingOpt {
	return roleBindingSubject(s...)
}

func roleBindingSubject(s ...rbacv1.Subject) RoleBindingOpt {
	return func(r *RoleBinding) {
		r.Subjects = s
	}
}

// Set role binding roleref
func RoleBindingRoleRef(rf rbacv1.RoleRef) RoleBindingOpt {
	return func(r *RoleBinding) {
		r.RoleRef = rf
	}
}
