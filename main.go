package main

import (
	"context"
	"fmt"
	"kitt/apps/admin"
	"kitt/kitt/config"
	"kitt/kitt/kernel"
	"kitt/kitt/render"
	"kitt/kitt/router"
	"kitt/kitt/runnables"
	"kitt/kitt/services"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Config
	config.LoadEnv("./")
	conf := config.NewConfigFromEnv()

	// Render
	renderService := render.NewEngine()
	renderService.WithFuncs(render.Funcs{
		"asset": func(path string) string {
			return fmt.Sprintf("%s?v=%d", path, time.Now().Unix())
		},
	})

	// Router
	routerService := router.NewRouter()
	routerService.With404(func(ctx router.RouteCtx) router.RouteResponse {
		ctx.Response().Send("Custom 404 here")
		return nil
	})
	routerService.To(router.NewStaticRoute("/css", "./public/css"))

	// Services
	container := services.NewServices()
	container.Set(renderService)
	container.Set(routerService)
	container.Set(conf)

	// Kernel
	k := kernel.NewKernel()
	k.WithServices(container)
	k.WithRunnable(runnables.NewWebServer(":3000", routerService))
	k.WithApp(&admin.App{})
	k.Boot()

	// Start server
	err := k.Run(ctx)

	if err != nil {
		fmt.Println(err)
	}
}
