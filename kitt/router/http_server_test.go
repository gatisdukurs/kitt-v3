package router

import (
	"context"
	"testing"
	"time"
)

func Test_Http_Server(t *testing.T) {
	t.Run("it serves and shuts down", func(t *testing.T) {
		fakeHandler := newFakeHttpHandler()
		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*2)
		defer cancel()

		server := NewHttpServer()
		err := server.ListenAndServe(ctx, ":12345", fakeHandler)

		if err != nil {
			t.Fatal(err)
		}
	})

}
