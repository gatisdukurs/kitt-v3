package pages

import (
	"kitt/app/admin/internal/shared"
	"kitt/kitt/router"
)

type Controller struct {
	shared.Controller
}

func (c Controller) Boot() {
	c.GET("/admin/pages", c.GetList)
	c.GET("/admin/pages/create", c.GetCreate)
	c.POST("/admin/pages", c.PostPage)
}

func (c Controller) GetList(ctx router.RouteCtx) router.RouteResponse {
	// View
	view := c.Layout("admin.layout")
	content := c.Partial("admin.pages.list")
	navigation := c.Navigation(ctx)

	view.WithPartial("content", content)
	view.WithPartial("navigation", navigation)
	view.WithHTMX("content", "navigation")
	// Send
	return c.Response(view)
}

func (c Controller) GetCreate(ctx router.RouteCtx) router.RouteResponse {
	// View
	view := c.Layout("admin.layout")
	content := c.Partial("admin.pages.create")
	navigation := c.Navigation(ctx)

	view.WithPartial("content", content)
	view.WithPartial("navigation", navigation)
	view.WithHTMX("content", "navigation")
	// Send
	return c.Response(view)
}

func (c Controller) PostPage(ctx router.RouteCtx) router.RouteResponse {
	view := c.Layout("admin.pages.form")
	return c.Response(view)
}

// func PostPages(ctx *kitt.RouteCtx) {
// 	formCtx := kitt.NewFormCtx().WithRequest(ctx.Request()).WithValidation(kitt.FormCtxValidators{
// 		"page.title": kitt.Validators{
// 			kitt.Required(),
// 			kitt.MinLength(3),
// 		},
// 		"page.content": kitt.Validators{
// 			kitt.Required(),
// 			kitt.MinLength(10),
// 		},
// 	})

// 	isValid := formCtx.Validate()

// 	if isValid {

// 		_, err := kitt.SQL().Exec(
// 			context.Background(),
// 			"INSERT INTO pages (title, content) VALUES (?,?)",
// 			formCtx.Value("page.title"),
// 			formCtx.Value("page.content"),
// 		)

// 		if err != nil {
// 			formCtx.SetError(err.Error())
// 		} else {
// 			formCtx.SetSuccess("Page created!")
// 		}
// 	}

// 	ctx.Response().Send(
// 		kitt.HTMX("admin.pages.form", formCtx),
// 		http.StatusOK,
// 	)
// }
