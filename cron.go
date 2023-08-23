// Copyright 2023 Cover Whale Insurance Solutions Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kopts

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CronSchedule is a cronjob schedule string
type CronSchedule string

const (
	Daily          CronSchedule = "0 0 * * *"
	Hourly         CronSchedule = "0 * * * *"
	Minute         CronSchedule = "* * * * *"
	Weekly         CronSchedule = "0 0 1 * *"
	Monthly        CronSchedule = "0 0 1 1/1 *"
	Yearly         CronSchedule = "0 0 1 1 *"
	Every15Minutes CronSchedule = "*/15 * * * *"
	Every5Minutes  CronSchedule = "*/5 * * * *"
	Every10Minutes CronSchedule = "*/10 * * *"
	EveryHalfHour  CronSchedule = "*/30 * * *"
)

// CronJob is a Kubernetes cron job
type CronJob struct {
	batchv1.CronJob
}

type CronJobOpt func(*CronJob)

// NewCronJob returns a cron job with the given name and options
func NewCronJob(name string, opts ...CronJobOpt) *CronJob {
	c := &CronJob{
		CronJob: batchv1.CronJob{
			TypeMeta: metav1.TypeMeta{
				Kind:       "CronJob",
				APIVersion: "batch/v1",
			},
			ObjectMeta: newObjectMeta(name),
			Spec: batchv1.CronJobSpec{
				JobTemplate: batchv1.JobTemplateSpec{
					Spec: batchv1.JobSpec{},
				},
			},
		},
	}

	for _, v := range opts {
		v(c)
	}

	return c
}

// CronJobNamespace sets the namespace for the cronjob
func CronJobNamespace(n string) CronJobOpt {
	return func(c *CronJob) {
		setNamespace(n, &c.ObjectMeta)
	}
}

// CronJobRestartPolicy sets the restart policy for the cronjob
func CronJobRestartPolicy(r corev1.RestartPolicy) CronJobOpt {
	return func(c *CronJob) {
		c.Spec.JobTemplate.Spec.Template.Spec.RestartPolicy = r
	}
}

// CronJobParallelism sets the parallelism for the cronjob
func CronJobParallelism(i int) CronJobOpt {
	parallel := int32(i)
	return func(c *CronJob) {
		c.Spec.JobTemplate.Spec.Parallelism = &parallel
	}
}

// CronJobCompletions sets the completions for the cronjob
func CronJobCompletions(i int) CronJobOpt {
	completions := int32(i)
	return func(c *CronJob) {
		c.Spec.JobTemplate.Spec.Completions = &completions
	}
}

// CronJobActiveDeadlineSeconds sets the active deadline seconds for the cronjob
func CronJobActiveDeadlineSeconds(i int) CronJobOpt {
	seconds := int64(i)
	return func(c *CronJob) {
		c.Spec.JobTemplate.Spec.ActiveDeadlineSeconds = &seconds
	}
}

// CronJobBackoffLimit sets the backoff limit for the cronjob
func CronJobBackoffLimit(i int) CronJobOpt {
	limit := int32(i)
	return func(c *CronJob) {
		c.Spec.JobTemplate.Spec.BackoffLimit = &limit
	}
}

// CronJobPodSpec sets the pod spec for the cronjob
func CronJobPodSpec(p PodSpec) CronJobOpt {
	return func(c *CronJob) {
		c.Spec.JobTemplate.Spec.Template = p.Spec
	}
}

// CronJobSchedule sets the schedule for the cronjob
func CronJobSchedule(cs CronSchedule) CronJobOpt {
	return func(c *CronJob) {
		c.Spec.Schedule = fmt.Sprintf("%s", string(cs))
	}
}

// CronJobConcurrency sets the concurrency for the cronjob
func CronJobConcurrency(p batchv1.ConcurrencyPolicy) CronJobOpt {
	return func(c *CronJob) {
		c.Spec.ConcurrencyPolicy = p
	}
}
