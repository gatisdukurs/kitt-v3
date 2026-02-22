package router

import (
	"io"
	"testing"
)

func Test_Route(t *testing.T) {
	t.Run("it handles", func(t *testing.T) {
		str := "Hello World!"
		buf := newBuf()
		ctx := makeRouteCtx(buf)

		r := NewRoute("/home")
		r.GET(func(ctx RouteCtx) {
			ctx.Response().Send(str)
		})

		r.Execute(ctx)

		assertEqual(t, getBufStr(buf), str)
	})

	t.Run("it returns correct pattern", func(t *testing.T) {
		r := NewRoute("/home")
		r.GET(func(ctx RouteCtx) {
			// handle
		})
		assertEqual(t, r.Pattern(), "GET /home")
	})
}

func makeRouteCtx(buf io.Writer) RouteCtx {
	response := NewResponse()
	response.WithWriter(buf)

	ctx := NewRouteCtx()
	ctx.WithResponse(response)

	return ctx
}
