package services

import (
	"fmt"
	"reflect"
	"strings"
)

type Services interface {
	Set(service any) Services
	SetWithKey(key string, service any) Services
	Get(t string) any
}

type services struct {
	container map[string]any
}

func (s *services) Set(service any) Services {
	key := strings.ToLower(reflect.TypeOf(service).String())

	fmt.Println("SET SERVICE", key)

	s.container[key] = service

	return s
}

func (s *services) SetWithKey(key string, service any) Services {
	s.container[key] = service
	return s
}

func (s *services) Get(key string) any {
	service, ok := s.container[key]

	if !ok {
		var zero any
		return zero
	}

	return service
}

func NewContainer() Services {
	return &services{
		container: make(map[string]any),
	}
}

func GetService[T any](s Services) T {
	key := strings.ToLower(reflect.TypeOf((*T)(nil)).String())
	raw := s.Get(key)
	service, ok := raw.(T)
	if !ok {
		fmt.Println(s)
		panic(fmt.Sprintf("Can't get service: %s", key))
	}

	return service
}

func GetServiceWithKey[T any](key string, s Services) T {
	raw := s.Get(key)
	service, ok := raw.(T)
	if !ok {
		fmt.Println(s)
		panic(fmt.Sprintf("Can't get service: %s", key))
	}

	return service
}
