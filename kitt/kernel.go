package kitt

type Kernel struct {
	Modules []Module
}

func (k Kernel) Boot() {
	for _, m := range k.Modules {
		m.Boot()
	}
}

func (k Kernel) RegisterTemplates() {
	for _, m := range k.Modules {
		registerTemplates(m.Templates(), "app/")
	}
}

func (k Kernel) RegisterEvents() {
	for _, m := range k.Modules {
		m.Events()
	}
}

func (k Kernel) RegisterRoutes(r *Router) {
	for _, m := range k.Modules {
		m.Routes(r)
	}
}

func (k Kernel) RegisterServices(s *Services) {
	for _, m := range k.Modules {
		m.Services(s)
	}
}
