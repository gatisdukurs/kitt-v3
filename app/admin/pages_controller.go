package admin

import (
	"kitt/kitt/router"
)

type PagesController struct {
	AdminController
}

func (c PagesController) Boot() {
	c.GET("/admin/pages", c.GetPages)
}

func (c PagesController) GetPages(ctx router.RouteCtx) {
	// View
	view := c.Layout("admin.layout")
	content := c.Partial("admin.pages.content")
	navigation := c.Navigation(ctx)

	view.WithPartial("content", content)
	view.WithPartial("navigation", navigation)
	// Send
	ctx.Response().Send(view)
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

// func GetView(ctx *kitt.RouteCtx) {
// 	ctx.Response().
// 		Send(
// 			kitt.View("admin.pages.view").
// 				WithCtx(ctx).
// 				WithContent(kitt.HTML("admin.pages.content", ctx)),
// 			http.StatusOK,
// 		)
// }
