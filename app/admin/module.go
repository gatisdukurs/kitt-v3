package admin

import (
	"context"
	"kitt/app/admin/internal/dashboard"
	"kitt/app/admin/internal/pages"
	"kitt/kitt"
)

type Module struct{}

func (Module) Boot() {
	dashboard.Controller{}.Boot()
	pages.Controller{}.Boot()
}

func (m Module) Events() {
	// kitt.Subscribe("router.onRequest", func(e kitt.Event) {
	// 	ctx := kitt.GetEventContext[*kitt.RouteCtx](e)
	// 	if ctx.Route().Module == "admin" {
	// 		ctx.SetVar("admin.navigation", m.nav.WithActive(ctx.Request().Path()))
	// 	}
	// })
}

func (Module) Templates() {}

func (Module) Services(s *kitt.Services) {
	kitt.Log("admin: Services")
}

func (Module) Migrate() {
	_, err := kitt.SQL().Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS pages (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL
		);
	`)

	kitt.Log("admin: Migrate")
	if err != nil {
		kitt.Log(err.Error())
	}
}
