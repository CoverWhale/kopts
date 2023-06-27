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

// Role is a Kubernetes role
type Role struct {
	rbacv1.Role
}

type RoleOpt func(*Role)

// NewRole returns a role with the given name and options
func NewRole(name string, opts ...RoleOpt) Role {
	r := Role{
		Role: rbacv1.Role{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Role",
				APIVersion: "rbac.authorization.k8s.io/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		},
	}

	for _, v := range opts {
		v(&r)
	}

	return r
}

// Set role namespace
func RoleNamespace(n string) RoleOpt {
	return func(r *Role) {
		r.Namespace = n
	}
}

// Set role policy rule
func RolePolicyRule(pr PolicyRule) RoleOpt {
	return rolePolicyRules(pr)
}

// Set multiple role policy rules
func RolePolicyRules(pr []PolicyRule) RoleOpt {
	return rolePolicyRules(pr...)
}

func rolePolicyRules(pr ...PolicyRule) RoleOpt {
	return func(r *Role) {
		for _, v := range pr {
			r.Rules = append(r.Rules, v.PolicyRule)
		}
	}
}
