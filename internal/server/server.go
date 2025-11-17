package server

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port int, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + strconv.Itoa(port),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    14 * time.Second,
		WriteTimeout:   14 * time.Second,
	}
	log.Printf("Server run on http://localhost%s", port)
	return s.httpServer.ListenAndServe()
}
