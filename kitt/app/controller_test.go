package app

import (
	"kitt/kitt"
	"kitt/kitt/kernel"
	"kitt/kitt/render"
	"kitt/kitt/router"
	"kitt/kitt/services"
	"testing"
)

func Test_Contoller(t *testing.T) {
	rtr := router.NewRouter()
	rndr := render.NewEngine()
	container := services.NewContainer()
	container.Set(rtr)
	container.Set(rndr)

	k := kernel.NewKernel()
	k.WithServices(container)
	app := kernel.NewFakeApp()
	k.WithApp(app)
	k.Boot()

	t.Run("it provides GET route shortcut", func(t *testing.T) {
		c := &Controller{}
		app.BootController(c)
		route := c.GET("/home", func(ctx router.RouteCtx) router.RouteResponse {
			return nil
		})
		assertEqual(t, route.Pattern(), "GET /home")
	})

	t.Run("it provides POST route shortcut", func(t *testing.T) {
		c := &Controller{}
		app.BootController(c)
		route := c.POST("/home", func(ctx router.RouteCtx) router.RouteResponse {
			return nil
		})
		assertEqual(t, route.Pattern(), "POST /home")
	})

	t.Run("it provides DELETE route shortcut", func(t *testing.T) {
		c := &Controller{}
		app.BootController(c)
		route := c.DELETE("/home", func(ctx router.RouteCtx) router.RouteResponse {
			return nil
		})
		assertEqual(t, route.Pattern(), "DELETE /home")
	})

	t.Run("it provides view", func(t *testing.T) {
		c := &Controller{}
		app.BootController(c)
		view := c.View("none")

		if _, ok := view.(kitt.View); !ok {
			t.Fatalf("not providing layout")
		}
	})

	t.Run("it provides partial", func(t *testing.T) {
		c := &Controller{}
		app.BootController(c)
		view := c.View("none")

		if _, ok := view.(render.View); !ok {
			t.Fatalf("not providing partial")
		}
	})

	t.Run("it provides route response", func(t *testing.T) {
		c := &Controller{}
		app.BootController(c)
		response := c.Response(newFakeRenderable("Hello World!"))

		if _, ok := response.(router.RouteResponse); !ok {
			t.Fatalf("not providing route response")
		}
	})

	t.Run("it provides route response with string", func(t *testing.T) {
		c := &Controller{}
		app.BootController(c)
		str := "Hello World!"
		response := c.ResponseString(str)

		if _, ok := response.(router.RouteResponse); !ok {
			t.Fatalf("not providing route response")
		}

		assertEqual(t, response.Body(), str)
	})

	t.Run("it has stack shortcut", func(t *testing.T) {
		c := &Controller{}
		app.BootController(c)
		stack := c.Stack()

		assertEqual(t, stack.Render(), "")
	})
}
