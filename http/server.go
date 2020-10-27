package http

import (
	"log"
	"net/http"
)

type Server struct {
	Addr    string
	Handler http.Handler
}

func (s *Server) Start() {
	server := &http.Server{
		Addr:    s.Addr,
		Handler: s.Handler,
	}
	log.Printf("listening on %v", s.Addr)
	server.ListenAndServe()
}
