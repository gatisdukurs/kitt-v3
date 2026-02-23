package kitt

import (
	"kitt/kitt/render"
	"kitt/kitt/router"
	"testing"
)

func Test_Contoller(t *testing.T) {
	t.Run("it provides GET route shortcut", func(t *testing.T) {
		c := &Controller{}
		route := c.GET("/home", func(ctx router.RouteCtx) {})

		if _, ok := route.(router.Route); !ok {
			t.Fatalf("not providing route")
		}

		assertEqual(t, route.Pattern(), "GET /home")
	})

	t.Run("it provides POST route shortcut", func(t *testing.T) {
		c := &Controller{}
		route := c.POST("/home", func(ctx router.RouteCtx) {})

		if _, ok := route.(router.Route); !ok {
			t.Fatalf("not providing route")
		}
		assertEqual(t, route.Pattern(), "POST /home")
	})

	t.Run("it provides DELETE route shortcut", func(t *testing.T) {
		c := &Controller{}
		route := c.DELETE("/home", func(ctx router.RouteCtx) {})

		if _, ok := route.(router.Route); !ok {
			t.Fatalf("not providing route")
		}
		assertEqual(t, route.Pattern(), "DELETE /home")
	})

	t.Run("it provides context", func(t *testing.T) {
		c := &Controller{}
		ctx := c.Ctx()

		if _, ok := ctx.(KittContext); !ok {
			t.Fatalf("not providing ctx")
		}
	})

	t.Run("it provides layout", func(t *testing.T) {
		c := &Controller{}
		view := c.Layout("none")

		if _, ok := view.(render.Layout); !ok {
			t.Fatalf("not providing layout")
		}
	})

	t.Run("it provides partial", func(t *testing.T) {
		c := &Controller{}
		view := c.Partial("none")

		if _, ok := view.(render.Partial); !ok {
			t.Fatalf("not providing partial")
		}
	})
}
