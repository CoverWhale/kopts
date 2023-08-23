package kopts

import (
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

type CronJob struct {
	batchv1.CronJob
}

type CronJobOpt func(*CronJob)

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

func CronJobNamespace(n string) CronJobOpt {
	return func(c *CronJob) {
		setNamespace(n, &c.ObjectMeta)
	}
}

func CronJobRestartPolicy(r corev1.RestartPolicy) CronJobOpt {
	return func(c *CronJob) {
		c.Spec.JobTemplate.Spec.Template.Spec.RestartPolicy = r
	}
}

func CronJobParallelism(i int) CronJobOpt {
	parallel := int32(i)
	return func(c *CronJob) {
		c.Spec.JobTemplate.Spec.Parallelism = &parallel
	}
}

func CronJobCompletions(i int) CronJobOpt {
	completions := int32(i)
	return func(c *CronJob) {
		c.Spec.JobTemplate.Spec.Completions = &completions
	}
}

func CronJobActiveDeadlineSeconds(i int) CronJobOpt {
	seconds := int64(i)
	return func(c *CronJob) {
		c.Spec.JobTemplate.Spec.ActiveDeadlineSeconds = &seconds
	}
}

func CronJobBackoffLimit(i int) CronJobOpt {
	limit := int32(i)
	return func(c *CronJob) {
		c.Spec.JobTemplate.Spec.BackoffLimit = &limit
	}
}

func CronJobPodSpec(p PodSpec) CronJobOpt {
	return func(c *CronJob) {
		c.Spec.JobTemplate.Spec.Template = p.Spec
	}
}

func CronJobSchedule(cs CronSchedule) CronJobOpt {
	return func(c *CronJob) {
		c.Spec.Schedule = fmt.Sprintf("%s", string(cs))
	}
}

func CronJobConcurrency(p batchv1.ConcurrencyPolicy) CronJobOpt {
	return func(c *CronJob) {
		c.Spec.ConcurrencyPolicy = p
	}
}
