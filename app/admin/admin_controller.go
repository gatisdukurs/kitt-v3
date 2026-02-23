package admin

import (
	"kitt/kitt"
	"kitt/kitt/render"
	"kitt/kitt/router"
)

type AdminController struct {
	kitt.Controller
}

func (c AdminController) Navigation(rctx router.RouteCtx) render.Partial {
	nav := Navigation{
		Items: []NavigationItem{
			{Label: "Dashboard", Path: "/admin"},
			{Label: "Pages", Path: "/admin/pages"},
		},
	}

	ctx := c.Ctx()
	ctx.Set("admin.navigation", nav.WithActive(rctx.Request().Path()))
	return c.Partial("admin.navigation").WithCtx(ctx.Basic())
}
