package http

import (
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func (s *Server) Routes() {
	router := &httprouter.Router{
		HandleOPTIONS: true,
	}

	router.HandlerFunc("GET", "/", s.handlePing())
	router.HandlerFunc("GET", "/health/ready", s.handleReady())
	router.HandlerFunc("GET", "/health/live", s.handleLive())
	router.HandlerFunc("GET", "/test", s.handleTest())

	middle := negroni.Classic()
	middle.UseHandler(router)
	s.Handler = middle
}
