package kitt

import (
	"context"
	"kitt/kitt/render"
	"kitt/kitt/router"
)

type TemplatePattern = string
type TemplatePatterns = []TemplatePattern

type Kitt interface {
	View(name string) render.View
	Router() router.Router
	WithTemplate(name string, str string) Kitt
	WithTemplates(patterns TemplatePatterns) Kitt
	WithTemplateFuncs(funcs render.Funcs) Kitt
	WithModule(module Module) Kitt
	WithHttpServer(handler router.HttpServer) Kitt
	Boot()
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
	renderer   render.Engine
	router     router.Router
	httpServer router.HttpServer
	modules    map[string]Module
}

func (k kitt) View(name string) render.View {
	l := render.NewView(name, k.renderer)
	return l
}

func (k kitt) Router() router.Router {
	return k.router
}

func (k *kitt) WithModule(module Module) Kitt {
	k.modules[module.Id()] = module
	return k
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

func (k kitt) Boot() {
	for _, m := range k.modules {
		m.Boot()
	}
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
	k := &kitt{
		modules: make(map[string]Module),
	}
	k.init()
	return k
}
