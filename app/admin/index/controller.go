package index

import (
	"kitt/kitt"
	"net/http"
)

func GetIndex(ctx *kitt.RouteCtx) {
	ctx.Response().Send(
		kitt.HTMX("admin.content", ctx).
			WithFallback("admin.index").
			WithOOB("navigation", "admin.navigation", ctx),
		http.StatusOK,
	)
}
