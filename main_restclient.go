package main

import (
	"context"
	"fmt"
	v12 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	// rest client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// get pod data
	pod := v1.Pod{}
	err1 := restClient.Get().Namespace("default").Resource("pod").Name(
		"nginx-deployment-648fc488dc-9wc27").Do(context.TODO()).Into(&pod)
	if err1 != nil {
		panic(err1)
	} else {
		fmt.Println(pod.Name)
	}

	// get deployment data
	deployment := v12.Deployment{}
	err2 := restClient.Get().Namespace("default").Resource("deployment").Name(
		"nginx-deployment").Do(context.TODO()).Into(&deployment)
	if err2 != nil {
		panic(err2)
	} else {
		fmt.Println(deployment.Name)
	}

}
