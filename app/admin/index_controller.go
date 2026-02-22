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
	view := c.View("admin.index")
	view.WithCtx(c.CtxWithNavigation(rctx).Basic())
	// Send
	rctx.Response().Send(view)
}
