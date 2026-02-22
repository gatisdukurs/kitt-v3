package router

import (
	"context"
	"net/http"
	"time"
)

type HttpServer interface {
	ListenAndServe(ctx context.Context, addr string, handler HttpHandler) error
	Shutdown() error
}

type httpServer struct {
	server *http.Server
}

func (hs *httpServer) ListenAndServe(ctx context.Context, addr string, handler HttpHandler) error {
	hs.server = &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	errCh := make(chan error, 1)

	go func() {
		errCh <- hs.server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		return hs.Shutdown()
	case err := <-errCh:
		return err
	}
}

func (hs httpServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return hs.server.Shutdown(ctx)
}

func NewHttpServer() HttpServer {
	return &httpServer{}
}
