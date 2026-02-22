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

// Ctx
func (Controller) Ctx() KittContext {
	return K().Ctx()
}

// Views
func (Controller) View(name string) render.Layout {
	view := K().Layout(name)
	return view
}
