package kitt

import (
	"kitt/kitt/config"
	"kitt/kitt/kernel"
	"kitt/kitt/render"
	"kitt/kitt/router"
	"kitt/kitt/services"
)

// Types
type Kernel = kernel.Kernel
type App = kernel.App
type AppController = kernel.AppController
type Runnable = kernel.Runnable

type Services = services.Services
type Router = router.Router
type Route = router.Route
type RouteResponse = router.RouteResponse
type RouteResponseSendable = router.RouteResponseSendable
type RouteHandler = router.RouteHandler
type Renderer = render.Engine
type Renderable = render.Renderable
type View = render.View
type ViewStack = render.ViewStack
type Context = render.AnyCtx
type Config = config.Config

// Methods
func Service[T any](s services.Services) T {
	return services.GetService[T](s)
}

func ServiceWithKey[T any](key string, s services.Services) T {
	return services.GetServiceWithKey[T](key, s)
}
