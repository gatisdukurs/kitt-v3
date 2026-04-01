package kernel

import (
	"context"
	"fmt"
	"kitt/kitt/router"
	"kitt/kitt/services"

	"golang.org/x/sync/errgroup"
)

type Kernel interface {
	WithServices(container services.Services) Kernel
	WithApp(module App) Kernel
	WithRunnable(runnable Runnable) Kernel
	Services() services.Services
	Boot()
	Run(ctx context.Context) error
}

type k struct {
	booted     bool
	httpServer router.HttpServer
	apps       map[string]App
	runnables  []Runnable
	container  services.Services
}

func (k *k) WithServices(container services.Services) Kernel {
	k.container = container
	return k
}

func (k *k) WithApp(app App) Kernel {
	k.apps[app.Id()] = app
	return k
}

func (k *k) WithRunnable(runnable Runnable) Kernel {
	k.runnables = append(k.runnables, runnable)
	return k
}

func (k *k) Run(ctx context.Context) error {
	if !k.booted {
		panic("NOT BOOTED")
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return k.runRunnables(ctx)
	})

	return g.Wait()
}

func (k *k) Services() services.Services {
	return k.container
}

func (k *k) Boot() {
	if k.booted {
		return
	}

	for _, app := range k.apps {
		app.Boot(k)
		k.runnables = append(k.runnables, app.Runnables()...)
	}

	k.booted = true
}

func (k *k) runRunnables(ctx context.Context) error {
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

func NewKernel() Kernel {
	return &k{
		apps: make(map[string]App),
	}
}
