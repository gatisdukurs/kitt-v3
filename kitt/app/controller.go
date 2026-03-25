package app

import (
	"fmt"
	"kitt/kitt"
	"kitt/kitt/htmx"
	"kitt/kitt/render"
	"kitt/kitt/router"
)

type Controller struct {
	app kitt.App
}

func (c *Controller) Boot(app kitt.App) {
	c.app = app
}

func (c *Controller) services() kitt.Services {
	if c.app == nil {
		panic("NO APP")
	}
	return c.app.Kernel().Services()
}

func (c *Controller) router() kitt.Router {
	return kitt.Service[kitt.Router](c.services())
}

func (c *Controller) renderer() kitt.Renderer {
	return kitt.Service[kitt.Renderer](c.services())
}

// Routing
func (c *Controller) GET(pattern string, handler kitt.RouteHandler) kitt.Route {
	route := router.NewRoute(pattern).GET(handler)
	c.router().To(route)
	return route
}

func (c *Controller) POST(pattern string, handler kitt.RouteHandler) kitt.Route {
	route := router.NewRoute(pattern).POST(handler)
	c.router().To(route)
	return route
}

func (c *Controller) DELETE(pattern string, handler kitt.RouteHandler) kitt.Route {
	route := router.NewRoute(pattern).DELETE(handler)
	c.router().To(route)
	return route
}

func (c *Controller) Response(sendable kitt.RouteResponseSendable) kitt.RouteResponse {
	response := router.NewRouteResponse()
	response.WithSendable(sendable)
	return response
}

func (c *Controller) ResponseString(body string) kitt.RouteResponse {
	response := router.NewRouteResponse()
	response.WithStringResponse(body)
	return response
}

// Ctx
func (c *Controller) Ctx() kitt.Context {
	return make(kitt.Context)
}

// Views
func (c *Controller) View(name string) kitt.View {
	return render.NewView(name, c.renderer())
}

func (c *Controller) Stack(renderables ...kitt.Renderable) kitt.ViewStack {
	stack := render.NewViewStack()

	for _, r := range renderables {
		stack.WithRenderable(r)
	}

	return stack
}

func (c *Controller) Htmx(view kitt.Renderable) htmx.HTMXElement {
	return htmx.NewElement(view)
}

func (c *Controller) ToastHtmx(t string, message string, args ...any) htmx.HTMXElement {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	toastCtx := c.Ctx()
	toastCtx["type"] = t
	toastCtx["message"] = message
	toast := c.View("admin.toast").WithCtx(toastCtx)
	toastHtmx := c.Htmx(toast).WithId("toast-container").WithSwap(htmx.HTMX_SWAP_AFTER_BEGIN)

	return toastHtmx
}
