package admin

import (
	"kitt/app/admin/internal/dashboard"
	"kitt/app/admin/internal/pages"
	"kitt/app/admin/internal/pages/jobs"
	"kitt/kitt"
	"kitt/kitt/queue"
	"kitt/kitt/repository"
)

type Module struct {
	pagesRepo   repository.Repository[pages.Page, int64]
	pagesCtrl   *pages.Controller
	dashCtrl    *dashboard.Controller
	pagesWorker *pages.SyncWorker
	worker      *queue.QueueWorker
}

func (Module) Id() string {
	return "admin"
}

func (m *Module) Boot() {
	// Queue setup
	q := queue.NewInMemoryQueue(100)
	d := queue.NewJobDispatcher()
	// Register jobs
	d.Register(jobs.Ajob{}, &jobs.AjobHandler{})

	m.worker = queue.NewQueueWorker("admin.worker", q, d)
	// Pages
	m.pagesRepo = repository.Repo[pages.Page, int64](repository.DRIVER_SQL, "db.sqlite")
	m.pagesCtrl = &pages.Controller{
		Pages: m.pagesRepo,
		Queue: q,
	}
	m.pagesWorker = &pages.SyncWorker{
		Pages: m.pagesRepo,
	}
	// Dashboard
	m.dashCtrl = &dashboard.Controller{}

	// Boot
	m.pagesCtrl.Boot()
	m.dashCtrl.Boot()
}

func (m *Module) Runnables() []kitt.Runnable {
	return []kitt.Runnable{
		m.pagesWorker,
		m.worker,
	}
}
