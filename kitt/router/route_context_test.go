package router

import "testing"

func Test_Route_Context(t *testing.T) {
	t.Run("it sets response", func(t *testing.T) {
		response := NewResponse()
		ctx := NewRouteCtx()

		ctx.WithResponse(response)

		assertEqual(t, response, ctx.Response())
	})

	t.Run("it sets request", func(t *testing.T) {
		request := NewRequest()
		ctx := NewRouteCtx()
		ctx.WithRequest(request)

		assertEqual(t, request, ctx.Request())
	})
}
