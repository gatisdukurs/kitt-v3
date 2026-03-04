package shared

import (
	"kitt/kitt"
	"kitt/kitt/render"
	"kitt/kitt/router"
)

type Controller struct {
	kitt.Controller
}

func (c Controller) Navigation(rctx router.RouteCtx) render.View {
	nav := Navigation{
		Items: []NavigationItem{
			{Label: "Dashboard", Path: "/admin"},
			{Label: "Pages", Path: "/admin/pages"},
		},
	}

	ctx := c.Ctx()
	ctx.Set("admin.navigation", nav.WithActive(rctx.Request().Path()))
	return c.View("admin.navigation").WithCtx(ctx.Basic())
}
