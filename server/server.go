package server

import (
	"flag"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type Server struct {
	Address   string
	client    *gin.Engine
	clientSet *kubernetes.Clientset
}

func NewServer(Address string) *Server {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	//config, err := rest.InClusterConfig()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	server := &Server{
		Address:   Address,
		client:    gin.Default(),
		clientSet: clientset,
	}

	// CORS
	server.client.Use(func() gin.HandlerFunc {
		return func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}

			c.Next()
		}
	}())
	v1 := server.client.Group("/api/v1")
	{
		v1.GET("/services", getServices(server.clientSet))
	}
	server.client.Static("/assets", "./dist/assets")
	server.client.StaticFile("/", "./dist/index.html")

	return server
}

func (s *Server) Run() {
	s.client.Run(s.Address)
}
