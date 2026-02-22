package kitt

type Kernel struct {
	Modules []Module
}

func (k Kernel) Boot() {
	for _, m := range k.Modules {
		m.Boot()
	}
}

func (k Kernel) Migrate() {
	for _, m := range k.Modules {
		m.Migrate()
	}
}

func (k Kernel) RegisterEvents() {
	for _, m := range k.Modules {
		m.Events()
	}
}

func (k Kernel) RegisterServices(s *Services) {
	for _, m := range k.Modules {
		m.Services(s)
	}
}
