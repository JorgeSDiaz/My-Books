package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Addr string
}

func New(port string) *Server {
	return &Server{
		Addr: port,
	}
}

func (s *Server) Start() error {
	routes()

	server := http.Server{
		Addr: s.Addr,
	}

	fmt.Println("Server running on port", s.Addr)

	return server.ListenAndServe()
}
