package dashboard

import (
	"kitt/app/admin/internal/shared"
	"kitt/kitt/router"
)

type Controller struct {
	shared.Controller
}

func (c Controller) Boot() {
	c.GET("/admin", c.GetDashboard)
}

func (c Controller) GetDashboard(rctx router.RouteCtx) router.RouteResponse {
	// View
	view := c.View("admin.layout")
	content := c.View("admin.dashboard")
	navigation := c.Navigation(rctx)

	view.WithPartial("content", content)
	view.WithPartial("navigation", navigation)
	view.WithHTMX("content", "navigation")
	// Send
	return c.Response(view)
}
