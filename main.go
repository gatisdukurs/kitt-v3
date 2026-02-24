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
		"app/admin/templates/*/**.html",
		"app/admin/templates/**.html",
	})
	kitt.K().Router().With404(func(ctx router.RouteCtx) router.RouteResponse {
		ctx.Response().Send("Custom 404 here")
		return nil
	})
	kitt.K().Router().To(router.NewStaticRoute("/css", "./public/css"))

	kitt.InitSQL().WithSQLite("db.sqlite")
	defer kitt.SQL().Close()

	s := &kitt.Services{}
	k := &kitt.Kernel{
		Modules: []kitt.Module{
			&admin.AdminModule{},
		},
	}

	k.Boot()
	k.Migrate()
	k.RegisterEvents()
	k.RegisterServices(s)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	err := kitt.K().ServeHttp(ctx, ":3000")

	if err != nil {
		fmt.Println(err)
	}
}
