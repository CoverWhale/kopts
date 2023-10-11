package main

import (
	"fmt"
	"log"

	"github.com/CoverWhale/kopts"
)

func printYaml(i interface{}) {
	data, err := kopts.MarshalYaml(i)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(data)
}

func main() {

	namespace := "testing"

	n := kopts.NewNamespace(namespace,
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
		kopts.ConfigMapNamespace(namespace),
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
		kopts.ContainerEnvFromSecret("test", "thing", "apiKey"),
		kopts.ContainerImagePullPolicy("Always"),
		kopts.ContainerArgs([]string{"server", "start"}),
		kopts.ContainerPort("http", 8080),
		kopts.ContainerPort("https", 443),
		kopts.ContainerLivenessProbeHTTP(hp),
		kopts.ContainerVolumeSource(conf.Name, "/tmp", conf.AsVolumeSource()),
	)

	// can also call the options later for conditionals
	if true {
		f := kopts.ContainerEnvVar("added", "later")
		f(&c)
	}

	p := kopts.NewPodSpec("test",
		kopts.PodLabel("app", "testing"),
		kopts.PodContainer(c),
		kopts.PodConfigmapAsVolume(conf.Name, conf),
	)

	d := kopts.NewDeployment("testing",
		kopts.DeploymentNamespace(namespace),
		kopts.DeploymentSelector("app", "testing"),
		kopts.DeploymentPodSpec(p),
		kopts.DeploymentReplicas(3),
	)

	x := kopts.DeploymentReplicas(1)
	x(d)

	printYaml(d)

	s := kopts.NewService("test",
		kopts.ServiceNamespace(namespace),
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
				Type:    "Prefix",
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
		kopts.SecretNamespace(namespace),
		kopts.SecretData("apiKey", []byte("thekey")),
	)

	printYaml(sec)

}
