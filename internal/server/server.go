package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port int, handler http.Handler) error {
	addr := fmt.Sprintf(":%d", port)
	s.httpServer = &http.Server{
		Addr:           addr,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    14 * time.Second,
		WriteTimeout:   14 * time.Second,
	}
	log.Printf("Server run on http://localhost%s", addr)
	return s.httpServer.ListenAndServe()
}
