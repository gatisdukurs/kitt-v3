package router

import (
	"net/http"
	"testing"
)

func Test_Route_Response(t *testing.T) {
	t.Run("it sets status and returns response", func(t *testing.T) {
		str := "Hello World!"
		sendable := newFakeRenderable(str)
		r := NewRouteResponse(sendable)
		r.WithStatus(http.StatusBadGateway)

		assertEqual(t, r.Status(), http.StatusBadGateway)
		assertEqual(t, r.Body(), str)
	})

	t.Run("it supports htmx", func(t *testing.T) {
		str := "Hello World!"
		sendable := newFakeRenderable(str)
		r := NewRouteResponse(sendable)
		r.WithStatus(http.StatusBadGateway)

		assertEqual(t, r.Status(), http.StatusBadGateway)
		assertEqual(t, r.HTMX(), str)
	})
}
