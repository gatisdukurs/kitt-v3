package shared

import (
	"kitt/kitt"
	"kitt/kitt/app"
	"kitt/kitt/render"
	"kitt/kitt/router"
)

type Controller struct {
	app.Controller
}

func (c *Controller) Boot(app kitt.App) {
	c.Controller.Boot(app)
}

func (c Controller) Navigation(rctx router.RouteCtx) render.View {
	nav := Navigation{
		Items: []NavigationItem{
			{Label: "Dashboard", Path: "/admin"},
			{Label: "Pages", Path: "/admin/pages"},
		},
	}

	ctx := c.Ctx()
	ctx["admin.navigation"] = nav.WithActive(rctx.Request().Path())
	return c.View("admin.navigation").WithCtx(ctx)
}
