package server

import "github.com/gin-gonic/gin"

type Server struct {
	Address string
	client  *gin.Engine
}

func NewServer(Address string) *Server {
	server := &Server{
		Address: Address,
		client:  gin.Default(),
	}

	return server
}

func (s *Server) Run() {
	s.client.Run(s.Address)
}
