package main

import (
	"fmt"
	"kitt/app/admin"
	"kitt/kitt"
	"text/template"
	"time"
)

func main() {
	kitt.K().WithTemplateFuncs(kitt.KittTemplateFuncs{
		"asset": func(path string) string {
			return fmt.Sprintf("%s?v=%d", path, time.Now().Unix())
		},
	})
	kitt.K().WithTemplates(kitt.KittTemplatePatterns{
		"app/admin/*/templates/*.html",
	})

	kitt.InitSQL().WithSQLite("db.sqlite")
	defer kitt.SQL().Close()
	kitt.RegisterTemplateFuncs(template.FuncMap{
		"asset": func(path string) string {
			return fmt.Sprintf("%s?v=%d", path, time.Now().Unix())
		},
	})

	r := &kitt.Router{}
	s := &kitt.Services{}
	k := &kitt.Kernel{
		Modules: []kitt.Module{
			&admin.Module{},
		},
	}

	k.Boot()
	k.Migrate()
	k.RegisterEvents()
	k.RegisterTemplates()
	k.RegisterRoutes(r)
	k.RegisterServices(s)
	kitt.Run(r)
}
