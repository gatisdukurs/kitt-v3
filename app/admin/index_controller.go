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

func (c IndexController) GetIndex(rctx router.RouteCtx) {
	// View
	view := c.Layout("admin.layout")
	content := c.Partial("admin.content")
	navigation := c.Navigation(rctx)

	view.WithPartial("content", content)
	view.WithPartial("navigation", navigation)
	// Send
	rctx.Response().Send(view)
}
