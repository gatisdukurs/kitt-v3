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

	t.Run("it does params", func(t *testing.T) {
		ctx := NewRouteCtx()
		ctx.WithParams(map[string]string{
			"id": "12",
		})

		assertEqual(t, ctx.Param("id"), "12")

		params := ctx.Params()

		assertEqual(t, params["id"], "12")
	})

	t.Run("it does int params", func(t *testing.T) {
		ctx := NewRouteCtx()
		ctx.WithParams(map[string]string{
			"id": "12",
		})

		assertEqual(t, ctx.ParamInt("id"), 12)
	})

	t.Run("it does int64 params", func(t *testing.T) {
		ctx := NewRouteCtx()
		ctx.WithParams(map[string]string{
			"id": "12",
		})

		assertEqual(t, ctx.ParamInt64("id"), int64(12))
	})
}
