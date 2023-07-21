package http

import (
	"context"

	"net"
	"net/http"
	"time"

	"github.com/mado/config"
	"github.com/mado/transport/http/handler"
)

type Server struct {
	handler *handler.Handler
	srv     *http.Server
	Notify  chan error
}

func NewServer(cfg *config.Config, handler *handler.Handler) *Server {
	srv := &http.Server{
		Addr:         net.JoinHostPort("", cfg.AppPort),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{srv: srv, handler: handler}
}

func (s *Server) StartServer() {
	s.srv.Handler = s.InitRoute()

	go func() {
		s.Notify <- s.srv.ListenAndServe()
		close(s.Notify)
	}()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return s.srv.Shutdown(ctx)
}
