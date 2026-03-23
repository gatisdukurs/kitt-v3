package dashboard

import (
	"kitt/app/admin/internal/shared"
	"kitt/kitt/htmx"
	"kitt/kitt/render"
	"kitt/kitt/router"
)

type Controller struct {
	shared.Controller
}

func (c *Controller) Boot() {
	c.GET("/admin", c.GetDashboard)
}

func (c *Controller) GetDashboard(ctx router.RouteCtx) router.RouteResponse {
	// View
	content := c.View("admin.dashboard")
	navigation := c.Navigation(ctx)

	if ctx.Request().HTMX() {
		contentHtmx := htmx.NewElement(content).WithId("content")
		navigationHtmx := htmx.NewElement(navigation).
			WithId("navigation").
			WithSwap(htmx.HTMX_SWAP_DEFAULT)

		stack := render.NewViewStack()
		stack.WithRenderable(contentHtmx)
		stack.WithRenderable(navigationHtmx)
		return c.Response(stack)
	} else {
		view := c.View("admin.layout")
		view.WithPartial("content", content)
		view.WithPartial("navigation", navigation)
		return c.Response(view)
	}
}
