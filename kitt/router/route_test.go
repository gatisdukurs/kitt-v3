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
		route.GET(func(ctx RouteCtx) RouteResponse {
			ctx.Response().Send(str)
			return nil
		})

		route.Execute(ctx)

		assertEqual(t, w.Sent(), str)
	})

	t.Run("it returns correct pattern", func(t *testing.T) {
		r := NewRoute("/home")
		r.GET(func(ctx RouteCtx) RouteResponse {
			// handle
			return nil
		})
		assertEqual(t, r.Pattern(), "GET /home")
	})

	t.Run("it supports POST", func(t *testing.T) {
		r := NewRoute("/home")
		r.POST(func(ctx RouteCtx) RouteResponse {
			// handle
			return nil
		})
		assertEqual(t, r.Pattern(), "POST /home")
	})

	t.Run("it supports DELETE", func(t *testing.T) {
		r := NewRoute("/home")
		r.DELETE(func(ctx RouteCtx) RouteResponse {
			// handle
			return nil
		})
		assertEqual(t, r.Pattern(), "DELETE /home")
	})

	t.Run("it matches *", func(t *testing.T) {
		r := NewRoute("/assets/*")
		r.GET(func(ctx RouteCtx) RouteResponse {
			//
			return nil
		})

		assertEqual(t, r.Match("GET", "/assets/style.css"), true)
		assertEqual(t, r.Match("GET", "/asset"), false)
	})

	t.Run("it matches exactly", func(t *testing.T) {
		r := NewRoute("/home")
		r.GET(func(ctx RouteCtx) RouteResponse {
			//
			return nil
		})

		assertEqual(t, r.Match("GET", "/home"), true)
		assertEqual(t, r.Match("GET", "/home/"), true)
		assertEqual(t, r.Match("GET", "/hom"), false)
		assertEqual(t, r.Match("GET", "/homee"), false)
	})

	t.Run("it sends route response", func(t *testing.T) {
		str := "Hello World!"
		r := NewRoute("/home")
		r.GET(func(ctx RouteCtx) RouteResponse {
			sendable := newFakeRenderable(str)
			response := NewRouteResponse(sendable)

			return response
		})

		ctx := NewRouteCtx()
		response := r.Execute(ctx)

		assertEqual(t, response.Body(), str)
	})

	t.Run("it also works if no response is sent from handler", func(t *testing.T) {
		r := NewRoute("/home")
		r.GET(func(ctx RouteCtx) RouteResponse {
			return nil
		})

		ctx := NewRouteCtx()
		response := r.Execute(ctx)

		assertEqual(t, response, nil)
	})
}
