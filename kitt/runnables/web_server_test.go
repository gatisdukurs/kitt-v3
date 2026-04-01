package runnables

import (
	"context"
	"io"
	"net"
	"net/http"
	"testing"
	"time"
)

func Test_Web_Server(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	// Reserve a free port first.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen on random port: %v", err)
	}
	addr := ln.Addr().String()
	_ = ln.Close()

	r := NewWebServer(addr, handler)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)

	go func() {
		errCh <- r.Run(ctx)
	}()

	// Wait until server is reachable.
	client := &http.Client{Timeout: 200 * time.Millisecond}
	url := "http://" + addr

	deadline := time.Now().Add(3 * time.Second)
	for {
		resp, err := client.Get(url)
		if err == nil {
			body, readErr := io.ReadAll(resp.Body)
			_ = resp.Body.Close()
			if readErr != nil {
				t.Fatalf("read response body: %v", readErr)
			}

			if resp.StatusCode != http.StatusOK {
				t.Fatalf("unexpected status: got %d want %d", resp.StatusCode, http.StatusOK)
			}

			if string(body) != "ok" {
				t.Fatalf("unexpected body: got %q want %q", string(body), "ok")
			}
			break
		}

		if time.Now().After(deadline) {
			t.Fatalf("server did not start in time: %v", err)
		}

		time.Sleep(50 * time.Millisecond)
	}

	// Stop server via context cancellation.
	cancel()

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("Run returned error: %v", err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("server did not shut down in time")
	}
}
