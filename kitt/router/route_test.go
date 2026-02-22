package router

import (
	"testing"
)

func Test_Route(t *testing.T) {
	t.Run("it handles", func(t *testing.T) {
		str := "Hello World!"
		w := newFakeResponseWriter()
		r := NewResponse()
		r.WithHttpResponse(w)

		ctx := NewRouteCtx()
		ctx.WithResponse(r)

		route := NewRoute("/home")
		route.GET(func(ctx RouteCtx) {
			ctx.Response().Send(str)
		})

		route.Execute(ctx)

		assertEqual(t, w.Sent(), str)
	})

	t.Run("it returns correct pattern", func(t *testing.T) {
		r := NewRoute("/home")
		r.GET(func(ctx RouteCtx) {
			// handle
		})
		assertEqual(t, r.Pattern(), "GET /home")
	})
}
