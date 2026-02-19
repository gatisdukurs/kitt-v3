package pages

import (
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
		formCtx.SetSuccess("Page created!")
	}

	ctx.Response().Send(
		kitt.HTMX("admin.pages.form", formCtx),
		http.StatusOK,
	)
}

func GetPages(ctx *kitt.RouteCtx) {
	ctx.Response().Send(
		kitt.HTMX("admin.pages.content", ctx).
			WithFallback("admin.pages.index").
			WithOOB("navigation", "admin.navigation", ctx),
		http.StatusOK,
	)
}
