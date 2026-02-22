package kitt

import (
	"context"
	"kitt/kitt/render"
	"kitt/kitt/router"
)

type TemplatePattern = string
type TemplatePatterns = []TemplatePattern

type Kitt interface {
	Layout(name string) render.Layout
	Partial(name string) render.Partial
	Router() router.Router
	Route(pattern string) router.Route
	Ctx() KittContext
	WithTemplate(name string, str string) Kitt
	WithTemplates(patterns TemplatePatterns) Kitt
	WithTemplateFuncs(funcs render.Funcs) Kitt
	WithHttpServer(handler router.HttpServer) Kitt
	ServeHttp(ctx context.Context, host string) error
	InTesting()
}

var kittInstance Kitt

func ensureInitialized() {
	if kittInstance == nil {
		kittInstance = newKitt()
	}
}

type kitt struct {
	renderer    render.Engine
	router      router.Router
	httpHandler router.HttpHandler
	httpServer  router.HttpServer
}

func (k kitt) Layout(name string) render.Layout {
	l := render.NewLayout(name, k.renderer)
	return l
}

func (k kitt) Partial(name string) render.Partial {
	p := render.NewPartial(name, k.renderer)
	return p
}

func (k kitt) Router() router.Router {
	return k.router
}

func (k kitt) Route(pattern string) router.Route {
	return router.NewRoute(pattern)
}

func (k kitt) Ctx() KittContext {
	return NewKittCtx()
}

func (k *kitt) WithTemplates(patterns TemplatePatterns) Kitt {
	for _, pattern := range patterns {
		k.renderer = k.renderer.WithTemplates(pattern)
	}
	return k
}

func (k *kitt) WithTemplate(name string, str string) Kitt {
	k.renderer = k.renderer.WithTemplate(name, str)
	return k
}

func (k *kitt) WithTemplateFuncs(funcs render.Funcs) Kitt {
	k.renderer = k.renderer.WithFuncs(funcs)
	return k
}

func (k *kitt) WithHttpServer(server router.HttpServer) Kitt {
	k.httpServer = server
	return k
}

func (k kitt) ServeHttp(ctx context.Context, host string) error {
	return k.httpServer.ListenAndServe(ctx, host, k.router)
}

func (k *kitt) InTesting() {
	k.init()
}

func (k *kitt) init() {
	k.renderer = render.NewEngine()
	k.router = router.NewRouter()
	k.httpServer = router.NewHttpServer()
}

func K() Kitt {
	ensureInitialized()
	return kittInstance
}

func newKitt() Kitt {
	k := &kitt{}
	k.init()
	return k
}
