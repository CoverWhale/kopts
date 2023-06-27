package main

import (
	"fmt"
	"log"

	"github.com/CoverWhale/kopts"
	corev1 "k8s.io/api/core/v1"
)

func printYaml(i interface{}) {
	data, err := kopts.MarshalYaml(i)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(data)
}

func main() {

	pv := kopts.NewPersistentVolume("myvolume",
		kopts.PersistentVolumeHostPath("/", corev1.HostPathDirectory),
	)

	printYaml(pv)

	n := kopts.NewNamespace("test",
		kopts.NamespaceAnnotation("test", "test2"),
		kopts.NamespaceAnnotations(map[string]string{
			"hey": "there",
			"yo":  "what's up",
		}),
	)

	printYaml(n)

	multiLine := `this is
a multiline
config
`

	conf := kopts.NewConfigMap("myconfigmap",
		kopts.ConfigMapNamespace("testing"),
		kopts.ConfigMapData("multiline", multiLine),
		kopts.ConfigMapDataMap(map[string]string{
			"testing": "123",
			"hey":     "this is a test",
		}),
		kopts.ConfigMapBinaryData("test", []byte("gimme some bytes")),
	)

	printYaml(conf)

	hp := kopts.HTTPProbe{
		Path:          "/healthz",
		Port:          8080,
		PeriodSeconds: 10,
		InitialDelay:  10,
	}

	c := kopts.NewContainer("test",
		kopts.ContainerImage("myrepo/myimage:latest"),
		kopts.ContainerEnvVar("hey", "there"),
		kopts.ContainerEnvFromSecret("testsecret", "thing", "apiKey"),
		kopts.ContainerImagePullPolicy("Always"),
		kopts.ContainerArgs([]string{"server", "start"}),
		kopts.ContainerPort("http", 8080),
		kopts.ContainerPort("https", 443),
		kopts.ContainerLivenessProbeHTTP(hp),
	)

	// can also call the options later for conditionals
	if true {
		f := kopts.ContainerEnvVar("added", "later")
		f(&c)
	}

	p := kopts.NewPodSpec("test",
		kopts.PodLabel("testing", "again"),
		kopts.PodContainer(c),
	)

	d := kopts.NewDeployment("testing",
		kopts.DeploymentNamespace("testing"),
		kopts.DeploymentSelector("app", "testing"),
		kopts.DeploymentPodSpec(p),
		kopts.DeploymentReplicas(3),
	)

	x := kopts.DeploymentReplicas(1)
	x(d)

	printYaml(d)

	s := kopts.NewService("test",
		kopts.ServiceNamespace("testing"),
		kopts.ServicePort(80, 8080),
		kopts.ServiceSelector("app", "mytest"),
	)
	printYaml(s)

	r := kopts.Rule{
		Host: "test.test.com",
		TLS:  true,
		Paths: []kopts.Path{
			{
				Name:    "/test",
				Service: "test",
				Port:    8080,
				Type:    "PathPrefix",
			},
		},
	}
	i := kopts.NewIngress("test",
		kopts.IngressClass("nginx"),
		kopts.IngressNamespace("testing"),
		kopts.IngressRule(r),
	)

	printYaml(i)

	sec := kopts.NewSecret("test",
		kopts.SecretNamespace("testing"),
		kopts.SecretData("apiKey", []byte("thekey")),
	)

	printYaml(sec)

}
