package admin

import (
	"kitt/kitt"
	"kitt/kitt/router"
)

type AdminController struct {
	kitt.Controller
}

func (c AdminController) CtxWithNavigation(rctx router.RouteCtx) kitt.KittContext {
	ctx := c.Ctx()
	ctx.Set("admin.navigation", c.Navigation(rctx))
	return ctx
}

func (AdminController) Navigation(ctx router.RouteCtx) Navigation {
	nav := Navigation{
		Items: []NavigationItem{
			{Label: "Dashboard", Path: "/admin"},
			{Label: "Pages", Path: "/admin/pages"},
		},
	}
	return nav.WithActive(ctx.Request().Path())
}
