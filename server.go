package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (serv *Server) Run(port string, handler http.Handler) error {
	serv.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return serv.httpServer.ListenAndServe()
}

func (serv *Server) Shutdown(ctx context.Context) error {
	return serv.httpServer.Shutdown(ctx)
}
