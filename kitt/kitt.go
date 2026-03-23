package kitt

import (
	"context"
	"fmt"
	"kitt/kitt/render"
	"kitt/kitt/router"

	"golang.org/x/sync/errgroup"
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
	WithRunnable(runnable Runnable) Kitt
	Boot()
	RunRunnables(ctx context.Context) error
	Run(ctx context.Context, host string) error
	ServeHttp(ctx context.Context, host string) error
	InTesting()
}

var kittInstance Kitt

func ensureInitialized() {
	if kittInstance == nil {
		kittInstance = newKitt()
	}
}

type k struct {
	booted     bool
	renderer   render.Engine
	router     router.Router
	httpServer router.HttpServer
	modules    map[string]Module
	runnables  []Runnable
}

func (k *k) View(name string) render.View {
	l := render.NewView(name, k.renderer)
	return l
}

func (k *k) Router() router.Router {
	return k.router
}

func (k *k) WithModule(module Module) Kitt {
	k.modules[module.Id()] = module
	return k
}

func (k *k) WithTemplates(patterns TemplatePatterns) Kitt {
	for _, pattern := range patterns {
		k.renderer = k.renderer.WithTemplates(pattern)
	}
	return k
}

func (k *k) WithTemplate(name string, str string) Kitt {
	k.renderer = k.renderer.WithTemplate(name, str)
	return k
}

func (k *k) WithTemplateFuncs(funcs render.Funcs) Kitt {
	k.renderer = k.renderer.WithFuncs(funcs)
	return k
}

func (k *k) WithHttpServer(server router.HttpServer) Kitt {
	k.httpServer = server
	return k
}

func (k *k) WithRunnable(runnable Runnable) Kitt {
	k.runnables = append(k.runnables, runnable)
	return k
}

func (k *k) Boot() {
	if k.booted {
		return
	}

	for _, m := range k.modules {
		m.Boot()
		k.runnables = append(k.runnables, m.Runnables()...)
	}

	k.booted = true
}

func (k *k) Run(ctx context.Context, host string) error {
	k.Boot()

	g, ctx := errgroup.WithContext(ctx)

	if host != "" {
		g.Go(func() error {
			return k.ServeHttp(ctx, host)
		})
	}

	g.Go(func() error {
		return k.RunRunnables(ctx)
	})

	return g.Wait()
}

func (k *k) ServeHttp(ctx context.Context, host string) error {
	if err := k.httpServer.ListenAndServe(ctx, host, k.router); err != nil {
		return fmt.Errorf("http server failed: %w", err)
	}
	return nil
}

func (k *k) RunRunnables(ctx context.Context) error {
	k.Boot()

	g, ctx := errgroup.WithContext(ctx)

	for _, runnable := range k.runnables {
		r := runnable

		g.Go(func() error {
			if err := r.Run(ctx); err != nil {
				return fmt.Errorf("runnable %q failed: %w", r.Id(), err)
			}
			return nil
		})
	}

	return g.Wait()
}

func (k *k) InTesting() {
	k.init()
}

func (k *k) init() {
	k.renderer = render.NewEngine()
	k.router = router.NewRouter()
	k.httpServer = router.NewHttpServer()
}

func K() Kitt {
	ensureInitialized()
	return kittInstance
}

func newKitt() Kitt {
	k := &k{
		modules: make(map[string]Module),
	}
	k.init()
	return k
}
