package server

import (
	"github.com/gin-gonic/gin"
	xkube "github.com/xenedium/xhomer/kubernetes"
	"k8s.io/client-go/kubernetes"
)

type Response struct {
	Message  string    `json:"message"`
	Services []Service `json:"services"`
}

type Service struct {
	Name      string  `json:"name"`
	Ports     []Port  `json:"ports"`
	Namespace string  `json:"namespace"`
	Type      string  `json:"type"`
	Ip        string  `json:"ip"`
	Image     *string `json:"image"`
}

type Port struct {
	Name string `json:"name"`
	Port int32  `json:"port"`
}

func getServices(clientSet *kubernetes.Clientset) gin.HandlerFunc {
	return func(c *gin.Context) {
		services, err := xkube.GetServices(clientSet)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		serviceList := []Service{}

		for _, service := range services.Items {
			ports := []Port{}
			for _, port := range service.Spec.Ports {
				ports = append(ports, Port{
					Name: port.Name,
					Port: port.Port,
				})
			}

			serviceList = append(serviceList, Service{
				Name:      service.Name,
				Ports:     ports,
				Namespace: service.Namespace,
				Type:      string(service.Spec.Type),
				Ip:        service.Spec.ClusterIP,
				Image:     nil,
			})
		}

		c.JSON(200, Response{
			Message:  "OK",
			Services: serviceList,
		})
	}
}
