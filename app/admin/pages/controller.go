package pages

import (
	"context"
	"kitt/kitt"
	"net/http"
)

func PostPages(ctx *kitt.RouteCtx) {
	formCtx := kitt.NewFormCtx().WithRequest(ctx.Request()).WithValidation(kitt.FormCtxValidators{
		"page.title": kitt.Validators{
			kitt.Required(),
			kitt.MinLength(3),
		},
		"page.content": kitt.Validators{
			kitt.Required(),
			kitt.MinLength(10),
		},
	})

	isValid := formCtx.Validate()

	if isValid {

		_, err := kitt.SQL().Exec(
			context.Background(),
			"INSERT INTO pages (title, content) VALUES (?,?)",
			formCtx.Value("page.title"),
			formCtx.Value("page.content"),
		)

		if err != nil {
			formCtx.SetError(err.Error())
		} else {
			formCtx.SetSuccess("Page created!")
		}
	}

	ctx.Response().Send(
		kitt.HTMX("admin.pages.form", formCtx),
		http.StatusOK,
	)
}

func GetPages(ctx *kitt.RouteCtx) {
	viewCtx := kitt.K().Ctx()
	viewCtx.Set("admin.navigation", ctx.Vars["admin.navigation"])

	view := kitt.
		K().
		Layout("admin.pages.index").
		WithCtx(viewCtx.Basic())

	ctx.Response().Send(view, http.StatusOK)

	// ctx.Response().Send(
	// 	kitt.HTMX("admin.pages.content", ctx).
	// 		WithFallback("admin.pages.index").
	// 		WithOOB("navigation", "admin.navigation", ctx),
	// 	http.StatusOK,
	// )
}

func GetView(ctx *kitt.RouteCtx) {
	ctx.Response().
		Send(
			kitt.View("admin.pages.view").
				WithCtx(ctx).
				WithContent(kitt.HTML("admin.pages.content", ctx)),
			http.StatusOK,
		)
}
