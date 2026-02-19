package main

import (
	"fmt"
	"kitt/app/admin"
	"kitt/kitt"
	"text/template"
	"time"
)

func main() {
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
	k.RegisterEvents()
	k.RegisterTemplates()
	k.RegisterRoutes(r)
	k.RegisterServices(s)
	kitt.Run(r)
}
