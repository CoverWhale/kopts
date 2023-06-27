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
	"k8s.io/apimachinery/pkg/util/intstr"
)

// Container holds a Kubernetes container
type Container struct {
	corev1.Container
}

type ContainerOpt func(*Container)

// NewContainer returns a container with the provided namd and any options
func NewContainer(name string, opts ...ContainerOpt) Container {
	c := Container{
		corev1.Container{
			Name: name,
		},
	}

	for _, v := range opts {
		v(&c)
	}

	return c
}

// Set container image
func ContainerImage(image string) ContainerOpt {
	return func(c *Container) {
		c.Image = image
	}
}

// Add container environment variable
func ContainerEnvVar(key, value string) ContainerOpt {
	return func(c *Container) {
		c.Env = append(c.Env, corev1.EnvVar{
			Name:  key,
			Value: value,
		})
	}
}

// Add container environment variable from secret
func ContainerEnvFromSecret(secret, name, key string) ContainerOpt {
	return func(c *Container) {
		c.Env = append(c.Env, corev1.EnvVar{
			Name: name,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: secret,
					},
					Key: key,
				},
			},
		})
	}
}

// Add container environment variable from config map
func ContainerEnvFromConfigMap(configmap, name, key string) ContainerOpt {
	return func(c *Container) {
		c.Env = append(c.Env, corev1.EnvVar{
			Name: name,
			ValueFrom: &corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: configmap,
					},
					Key: key,
				},
			},
		})
	}
}

// Set container commands
func ContainerCommands(commands []string) ContainerOpt {
	return func(c *Container) {
		c.Command = commands
	}
}

// Set container pull policy
func ContainerImagePullPolicy(policy corev1.PullPolicy) ContainerOpt {
	return func(c *Container) {
		c.ImagePullPolicy = policy
	}
}

// Set Container args
func ContainerArgs(args []string) ContainerOpt {
	return func(c *Container) {
		c.Args = args
	}
}

// Add Container port
func ContainerPort(name string, port int) ContainerOpt {
	return func(c *Container) {
		c.Ports = append(c.Ports, corev1.ContainerPort{
			Name:          name,
			ContainerPort: int32(port),
		})
	}
}

// Add container Volume
func ContainerVolume(path string, pv PersistentVolume) ContainerOpt {
	return func(c *Container) {
		c.VolumeMounts = append(c.VolumeMounts, corev1.VolumeMount{
			MountPath: path,
			Name:      pv.ObjectMeta.Name,
		})
	}
}

// Liveness probe holds the information for a Kubernetes liveness probe
type HTTPProbe struct {
	Path          string
	Port          int
	InitialDelay  int
	PeriodSeconds int
}

// Add liveness probe
func ContainerLivenessProbeHTTP(h HTTPProbe) ContainerOpt {
	return func(c *Container) {
		c.LivenessProbe = &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				HTTPGet: &corev1.HTTPGetAction{
					Path: h.Path,
					Port: intstr.FromInt(h.Port),
				},
			},
			InitialDelaySeconds: int32(h.InitialDelay),
			PeriodSeconds:       int32(h.PeriodSeconds),
		}
	}
}
