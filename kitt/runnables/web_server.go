package runnables

import (
	"context"
	"fmt"
	"kitt/kitt/kernel"
	"net/http"
)

type webServer struct {
	addr    string
	handler http.Handler
}

func (w webServer) Id() string {
	return "web.server.runable"
}

func (w webServer) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    w.addr,
		Handler: w.handler,
	}
	go func() {
		<-ctx.Done()
		_ = server.Shutdown(context.Background())
	}()

	fmt.Printf("Web Server running on: %s\n", w.addr)
	err := server.ListenAndServe()

	if err == http.ErrServerClosed {
		return nil
	}

	return err
}

func NewWebServer(addr string, handler http.Handler) kernel.Runnable {
	return &webServer{
		addr:    addr,
		handler: handler,
	}
}
