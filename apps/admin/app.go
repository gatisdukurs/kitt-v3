package admin

import (
	"kitt/apps/admin/internal/dashboard"
	"kitt/apps/admin/internal/pages"
	"kitt/apps/admin/internal/pages/jobs"
	"kitt/kitt"
	"kitt/kitt/config"
	"kitt/kitt/queue"
	"kitt/kitt/repository"
)

type App struct {
	pagesRepo   repository.Repository[pages.Page, int64]
	pagesCtrl   *pages.Controller
	dashCtrl    *dashboard.Controller
	pagesWorker *pages.SyncWorker
	worker      *queue.QueueWorker
	kernel      kitt.Kernel
}

func (App) Id() string {
	return "admin"
}

func (a *App) Kernel() kitt.Kernel {
	return a.kernel
}

func (a *App) Boot(kernel kitt.Kernel) {
	// Kernel setup
	a.kernel = kernel

	conf := kitt.Service[kitt.Config](a.kernel.Services())

	// Setup templates
	rndr := kitt.Service[kitt.Renderer](a.kernel.Services())
	rndr.WithTemplates("./apps/admin/internal/*/*.html")

	// Queue setup
	q := queue.NewInMemoryQueue(100)
	d := queue.NewJobDispatcher()
	// Register jobs
	d.Register(jobs.Ajob{}, &jobs.AjobHandler{})

	a.worker = queue.NewQueueWorker("admin.worker", q, d)
	// Pages
	a.pagesRepo = repository.Repo[pages.Page, int64](repository.DRIVER_SQL, conf.Get(config.CONF_DB_SQLITE))
	a.pagesCtrl = &pages.Controller{
		Pages: a.pagesRepo,
		Queue: q,
	}
	a.pagesWorker = &pages.SyncWorker{
		Pages: a.pagesRepo,
	}
	// Dashboard
	a.dashCtrl = &dashboard.Controller{}

	// Boot
	a.BootController(a.pagesCtrl)
	a.BootController(a.dashCtrl)
}

func (a *App) BootController(c kitt.AppController) {
	c.Boot(a)
}

func (a *App) Runnables() []kitt.Runnable {
	return []kitt.Runnable{
		a.pagesWorker,
		a.worker,
	}
}
