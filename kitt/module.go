package kitt

type Module interface {
	Boot()
	Events()
	Templates() TemplateSet
	Routes(r *Router)
	Services(s *Services)
	Migrate()
}
