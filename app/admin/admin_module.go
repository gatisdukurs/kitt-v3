package admin

import (
	"context"
	"kitt/kitt"
)

type AdminModule struct{}

func (AdminModule) Boot() {
	IndexController{}.Boot()
	PagesController{}.Boot()
}

func (m AdminModule) Events() {
	// kitt.Subscribe("router.onRequest", func(e kitt.Event) {
	// 	ctx := kitt.GetEventContext[*kitt.RouteCtx](e)
	// 	if ctx.Route().Module == "admin" {
	// 		ctx.SetVar("admin.navigation", m.nav.WithActive(ctx.Request().Path()))
	// 	}
	// })
}

func (AdminModule) Templates() {}

func (AdminModule) Services(s *kitt.Services) {
	kitt.Log("admin: Services")
}

func (AdminModule) Migrate() {
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
