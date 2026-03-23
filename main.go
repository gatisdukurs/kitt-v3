package main

import (
	"context"
	"fmt"
	"kitt/app/admin"
	"kitt/kitt"
	"kitt/kitt/render"
	"kitt/kitt/router"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	k := kitt.K()

	// Some globals
	k.WithTemplateFuncs(render.Funcs{
		"asset": func(path string) string {
			return fmt.Sprintf("%s?v=%d", path, time.Now().Unix())
		},
	})
	k.WithTemplates(kitt.TemplatePatterns{
		"app/admin/internal/*/*.html",
	})
	k.Router().With404(func(ctx router.RouteCtx) router.RouteResponse {
		ctx.Response().Send("Custom 404 here")
		return nil
	})
	k.Router().To(router.NewStaticRoute("/css", "./public/css"))

	// Modules
	k.WithModule(&admin.Module{})

	// Boot
	k.Boot()

	// Start server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	err := k.ServeHttp(ctx, ":3000")

	if err != nil {
		fmt.Println(err)
	}
}
