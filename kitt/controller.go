package kitt

import (
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
	response := K().Response(sendable)
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
