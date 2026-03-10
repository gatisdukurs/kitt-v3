package repository

import (
	"fmt"
	"reflect"
	"slices"
)

type Repository[T interface{}, ID comparable] interface {
	Create(m *T) (ID, error)
	ByID(id ID) (T, error)
	Update(m *T) error
	Delete(id ID) error
}

type repo[T interface{}, ID comparable] struct {
	driver    Driver[ID]
	modelMeta ModelMeta
}

func (r repo[T, ID]) Create(m *T) (ID, error) {
	values := DriverValues{}
	v := reflect.ValueOf(m).Elem()

	for _, fieldMeta := range r.modelMeta.Fields {
		values[fieldMeta.Key] = v.Field(fieldMeta.Index).Interface()
	}

	return r.driver.Insert(values)
}

func (r repo[T, ID]) ByID(id ID) (T, error) {
	values, err := r.driver.ByID(id)
	var zero T

	if err != nil {
		return zero, err
	}

	v := reflect.ValueOf(&zero).Elem()
	for _, fieldMeta := range r.modelMeta.Fields {
		if _, ok := values[fieldMeta.Key]; !ok {
			continue
		}

		field := v.Field(fieldMeta.Index)
		if field.CanSet() {
			value := reflect.ValueOf(values[fieldMeta.Key])
			if value.Type().AssignableTo(fieldMeta.Type) {
				field.Set(value)
			} else if value.Type().ConvertibleTo(fieldMeta.Type) {
				field.Set(value.Convert(fieldMeta.Type))
			}
		}
	}

	return zero, nil
}

func (r repo[T, ID]) Update(m *T) error {
	values := DriverValues{}
	v := reflect.ValueOf(m).Elem()

	var id ID
	var zero ID

	for _, fieldMeta := range r.modelMeta.Fields {
		if slices.Contains(fieldMeta.Flags, "pk") {
			raw := v.Field(fieldMeta.Index).Interface()
			if converted, ok := raw.(ID); ok {
				id = converted
			}
			continue
		}
		values[fieldMeta.Key] = v.Field(fieldMeta.Index).Interface()
	}

	if id == zero {
		return fmt.Errorf("primary key not found")
	}

	return r.driver.Update(values, id)
}

func (r repo[T, ID]) Delete(id ID) error {
	return r.driver.Delete(id)
}

func NewRepo[T interface{}, ID comparable](driver Driver[ID]) (Repository[T, ID], error) {
	// Read the model
	reader := NewModelReader[T]("db")
	modelMeta := reader.Read()
	driver.WithModelMeta(modelMeta)

	// Make sure collection exists
	err := driver.CreateCollection()

	if err != nil {
		return nil, err
	}

	// Init
	repo := &repo[T, ID]{
		modelMeta: modelMeta,
		driver:    driver,
	}

	return repo, nil
}
