package index

import (
	"kitt/kitt"
	"net/http"
)

func GetIndex(ctx *kitt.RouteCtx) {
	bctx := kitt.K().Ctx()
	bctx.Set("admin.navigation", ctx.Vars["admin.navigation"])

	ctx.Response().Send(
		kitt.K().Layout("admin.index").WithCtx(bctx.Basic()),
		http.StatusOK,
	)

	// ctx.Response().Send(
	// 	kitt.HTMX("admin.content", ctx).
	// 		WithFallback("admin.index").
	// 		WithOOB("navigation", "admin.navigation", ctx),
	// 	http.StatusOK,
	// )
}
