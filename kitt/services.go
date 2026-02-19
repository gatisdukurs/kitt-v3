package kitt

type Service interface{}

type Services struct {
	Services map[string]Service
}

func (s *Services) Register(name string, service Service) {
	if s.Services == nil {
		s.Services = map[string]Service{}
	}

	s.Services[name] = service
}

func (s Services) Get(name string) Service {
	if service, exists := s.Services[name]; exists {
		return service
	}

	return nil
}
