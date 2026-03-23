package admin

import (
	"kitt/app/admin/internal/dashboard"
	"kitt/app/admin/internal/pages"
	"kitt/kitt/repository"
)

type Module struct {
	pagesRepo repository.Repository[pages.Page, int64]
	pagesCtrl *pages.Controller
	dashCtrl  *dashboard.Controller
}

func (Module) Id() string {
	return "admin"
}

func (m *Module) Boot() {
	// Pages
	m.pagesRepo = repository.Repo[pages.Page, int64](repository.DRIVER_SQL, "db.sqlite")
	m.pagesCtrl = &pages.Controller{
		Pages: m.pagesRepo,
	}
	m.pagesCtrl.Boot()

	// Dashboard
	m.dashCtrl = &dashboard.Controller{}
	m.dashCtrl.Boot()
}
