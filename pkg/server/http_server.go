package server

import (
	"Kode_test/config"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewHttpServer(httpCfg config.HTTP, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           httpCfg.Host + ":" + httpCfg.Port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
