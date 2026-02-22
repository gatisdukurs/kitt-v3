package router

import (
	"net/http"
	"testing"
)

func Test_Response(t *testing.T) {
	t.Run("it writes response with string", func(t *testing.T) {
		w := newFakeResponseWriter()
		r := NewResponse()
		r.WithHttpResponse(w)

		str := "Hello World!"

		r.Send(str)

		assertEqual(t, w.Sent(), str)
	})

	t.Run("it writes response with renderable", func(t *testing.T) {
		str := "Hello World!"
		renderable := newFakeRenderable(str)
		w := newFakeResponseWriter()
		r := NewResponse()
		r.WithHttpResponse(w)

		r.Send(renderable)

		assertEqual(t, w.Sent(), str)
	})

	t.Run("it sets and sends status", func(t *testing.T) {
		str := "Hello World!"
		rw := newFakeResponseWriter()
		r := NewResponse()
		r.WithHttpResponse(rw)
		r.WithStatus(http.StatusAccepted)
		r.Send(str)

		assertEqual(t, rw.Status, http.StatusAccepted)
		assertEqual(t, rw.Sent(), str)
	})
}
