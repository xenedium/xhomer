package kubernetes

import (
	"context"

	v1Core "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetServices(clientset *kubernetes.Clientset) (*v1Core.ServiceList, error) {

	services, err := clientset.CoreV1().Services("").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return services, nil
}
