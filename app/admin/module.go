package admin

import (
	"context"
	"kitt/app/admin/index"
	"kitt/app/admin/pages"
	"kitt/kitt"
	"net/http"
)

type Module struct {
	nav Navigation
}

func (m *Module) Boot() {
	m.nav = Navigation{
		Items: []NavigationItem{
			{Label: "Dashboard", Path: "/admin"},
			{Label: "Pages", Path: "/admin/pages"},
		},
	}
}

func (m Module) Events() {
	kitt.Subscribe("router.onRequest", func(e kitt.Event) {
		ctx := kitt.GetEventContext[*kitt.RouteCtx](e)
		if ctx.Route().Module == "admin" {
			ctx.SetVar("admin.navigation", m.nav.WithActive(ctx.Request().Path()))
		}
	})
}

func (Module) Templates() kitt.TemplateSet {
	return kitt.TemplateSet{
		Pattern: "admin/*/templates/*.html",
	}
}

func (Module) Routes(r *kitt.Router) {
	r.To(kitt.Route{
		Module:  "admin",
		Method:  http.MethodGet,
		Pattern: "/admin",
		Handler: index.GetIndex,
	})
	r.To(kitt.Route{
		Module:  "admin",
		Method:  http.MethodGet,
		Pattern: "/admin/pages",
		Handler: pages.GetPages,
	})
	r.To(kitt.Route{
		Module:  "admin",
		Method:  http.MethodPost,
		Pattern: "/admin/pages",
		Handler: pages.PostPages,
	})
	// r.To(kitt.Route{
	// 	Module:  "admin",
	// 	Method:  http.MethodGet,
	// 	Pattern: "/admin/view",
	// 	Handler: pages.GetView,
	// })
}

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
