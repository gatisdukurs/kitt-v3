package kitt

type Module interface {
	Boot()
	Events()
	Templates()
	Services(s *Services)
	Migrate()
}
