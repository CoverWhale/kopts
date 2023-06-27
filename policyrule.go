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
)

type Verb string

const (
	Create           Verb = "create"
	Delete           Verb = "delete"
	Deletecollection Verb = "deletecollection"
	Get              Verb = "get"
	List             Verb = "list"
	Patch            Verb = "patch"
	Update           Verb = "update"
	Watch            Verb = "watch"
)

// PolicyRule is a Kubernetes policy rule
type PolicyRule struct {
	rbacv1.PolicyRule
}

type PolicyRuleOpt func(*PolicyRule)

// NewPolicyRule returns a policy rule with the given name and options
func NewPolicyRule(name string, opts ...PolicyRuleOpt) PolicyRule {
	pr := PolicyRule{
		PolicyRule: rbacv1.PolicyRule{},
	}

	for _, v := range opts {
		v(&pr)
	}

	return pr
}

// Set policy rule verb
func PolicyRuleVerb(v Verb) PolicyRuleOpt {
	return policyRuleVerbs(v)
}

// Set multiple rule verbs
func PolicyRuleVerbs(verbs []Verb) PolicyRuleOpt {
	return policyRuleVerbs(verbs...)
}

func policyRuleVerbs(verbs ...Verb) PolicyRuleOpt {
	return func(pr *PolicyRule) {
		for _, v := range verbs {
			pr.Verbs = append(pr.Verbs, string(v))
		}
	}
}

// Set policy rule API group
func PolicyRuleAPIGroup(group string) PolicyRuleOpt {
	return policyRuleAPIGroups(group)
}

// Set multiple rule API groups
func PolicyRuleAPIGroups(groups []string) PolicyRuleOpt {
	return policyRuleAPIGroups(groups...)
}

func policyRuleAPIGroups(groups ...string) PolicyRuleOpt {
	return func(pr *PolicyRule) {
		pr.APIGroups = groups
	}
}

// Set policy rule resource
func PolicyRuleResource(resource string) PolicyRuleOpt {
	return policyRuleResources(resource)
}

// Set multiple policy rule resources
func PolicyRuleResources(resources []string) PolicyRuleOpt {
	return policyRuleResources(resources...)
}

func policyRuleResources(resources ...string) PolicyRuleOpt {
	return func(pr *PolicyRule) {
		pr.Resources = resources
	}
}

// Set policy rule resource name
func PolicyRuleResourceName(rn string) PolicyRuleOpt {
	return policyRuleResourceNames(rn)
}

// Set multiple policy rule resource names
func PolicyRuleResourceNames(rn []string) PolicyRuleOpt {
	return policyRuleResourceNames(rn...)
}

func policyRuleResourceNames(rn ...string) PolicyRuleOpt {
	return func(pr *PolicyRule) {
		pr.ResourceNames = rn
	}
}

// Set policy rule resource URL
func PolicyRuleNonResourceURL(nru string) PolicyRuleOpt {
	return policyRuleNonResourceURLs(nru)
}

// Set multiple policy rule resource URLs
func PolicyRuleNonResourceURLs(nru []string) PolicyRuleOpt {
	return policyRuleNonResourceURLs(nru...)
}

func policyRuleNonResourceURLs(nru ...string) PolicyRuleOpt {
	return func(pr *PolicyRule) {
		pr.NonResourceURLs = nru
	}
}
