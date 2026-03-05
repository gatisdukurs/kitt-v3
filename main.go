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
	kitt.K().WithTemplateFuncs(render.Funcs{
		"asset": func(path string) string {
			return fmt.Sprintf("%s?v=%d", path, time.Now().Unix())
		},
	})
	kitt.K().WithTemplates(kitt.TemplatePatterns{
		"app/admin/internal/*/*.html",
	})
	kitt.K().Router().With404(func(ctx router.RouteCtx) router.RouteResponse {
		ctx.Response().Send("Custom 404 here")
		return nil
	})
	kitt.K().Router().To(router.NewStaticRoute("/css", "./public/css"))

	// ----- OLD START

	kitt.InitSQL().WithSQLite("db.sqlite")
	defer kitt.SQL().Close()

	admin.Module{}.Boot()

	// ---- OLD END

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	err := kitt.K().ServeHttp(ctx, ":3000")

	if err != nil {
		fmt.Println(err)
	}
}
