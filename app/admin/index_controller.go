package admin

import (
	"kitt/kitt/router"
)

type IndexController struct {
	AdminController
}

func (c IndexController) Boot() {
	c.GET("/admin", c.GetIndex)
}

func (c IndexController) GetIndex(rctx router.RouteCtx) router.RouteResponse {
	// View
	view := c.Layout("admin.layout")
	content := c.Partial("admin.content")
	navigation := c.Navigation(rctx)

	view.WithPartial("content", content)
	view.WithPartial("navigation", navigation)
	view.WithHTMX("content", "navigation")
	// Send
	return c.Response(view)
}
