package kitt

import (
	"fmt"
	"kitt/kitt/htmx"
	"kitt/kitt/render"
	"kitt/kitt/router"
)

type Controller struct{}

func (Controller) Boot() {}

// Routing
func (Controller) GET(pattern string, handler router.RouteHandler) router.Route {
	router := K().Router()
	route := K().Route(pattern).GET(handler)
	router.To(route)
	return route
}

func (Controller) POST(pattern string, handler router.RouteHandler) router.Route {
	router := K().Router()
	route := K().Route(pattern).POST(handler)
	router.To(route)
	return route
}

func (Controller) DELETE(pattern string, handler router.RouteHandler) router.Route {
	router := K().Router()
	route := K().Route(pattern).DELETE(handler)
	router.To(route)
	return route
}

func (Controller) Response(sendable router.RouteResponseSendable) router.RouteResponse {
	response := router.NewRouteResponse()
	response.WithSendable(sendable)
	return response
}

func (Controller) ResponseString(body string) router.RouteResponse {
	response := router.NewRouteResponse()
	response.WithStringResponse(body)
	return response
}

// Ctx
func (Controller) Ctx() KittContext {
	return K().Ctx()
}

// Views
func (Controller) View(name string) render.View {
	view := K().View(name)
	return view
}

func (Controller) Stack(renderables ...render.Renderable) render.ViewStack {
	stack := render.NewViewStack()

	for _, r := range renderables {
		stack.WithRenderable(r)
	}

	return stack
}

func (Controller) Htmx(view render.Renderable) htmx.HTMXElement {
	return htmx.NewElement(view)
}

func (c Controller) ToastHtmx(t string, message string, args ...any) htmx.HTMXElement {

	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	toastCtx := c.Ctx()
	toastCtx.Set("type", t)
	toastCtx.Set("message", message)
	toast := c.View("admin.toast").WithCtx(toastCtx.Basic())
	toastHtmx := c.Htmx(toast).WithId("toast-container").WithSwap(htmx.HTMX_SWAP_AFTER_BEGIN)

	return toastHtmx
}
