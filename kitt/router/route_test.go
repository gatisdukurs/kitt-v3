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

	t.Run("it supports POST", func(t *testing.T) {
		r := NewRoute("/home")
		r.POST(func(ctx RouteCtx) {
			// handle
		})
		assertEqual(t, r.Pattern(), "POST /home")
	})

	t.Run("it supports DELETE", func(t *testing.T) {
		r := NewRoute("/home")
		r.DELETE(func(ctx RouteCtx) {
			// handle
		})
		assertEqual(t, r.Pattern(), "DELETE /home")
	})

	t.Run("it matches *", func(t *testing.T) {
		r := NewRoute("/assets/*")
		r.GET(func(ctx RouteCtx) {
			//
		})

		assertEqual(t, r.Match("GET", "/assets/style.css"), true)
		assertEqual(t, r.Match("GET", "/asset"), false)
	})

	t.Run("it matches exactly", func(t *testing.T) {
		r := NewRoute("/home")
		r.GET(func(ctx RouteCtx) {
			//
		})

		assertEqual(t, r.Match("GET", "/home"), true)
		assertEqual(t, r.Match("GET", "/home/"), true)
		assertEqual(t, r.Match("GET", "/hom"), false)
		assertEqual(t, r.Match("GET", "/homee"), false)
	})
}
