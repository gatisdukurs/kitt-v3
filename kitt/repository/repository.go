package repository

import (
	"reflect"
)

type Repository[T interface{}, ID comparable] interface {
	Query() FindQuery
	Create(m T) (ID, error)
	ByID(id ID) (T, error)
	Find(query FindQuery) []T
	Update(m T) error
	Delete(id ID) error
}

type repo[T interface{}, ID comparable] struct {
	driver    Driver[ID]
	modelMeta ModelMeta
}

func (r repo[T, ID]) Query() FindQuery {
	return r.driver.Query()
}

func (r repo[T, ID]) Find(q FindQuery) []T {
	items := []T{}
	values, err := r.driver.Find(q)

	if err == nil {
		for _, v := range values {
			items = append(items, r.toModel(v))
		}
	}

	return items
}

func (r repo[T, ID]) Create(m T) (ID, error) {
	values := r.toDriverValues(m)

	idKey := r.modelMeta.PrimaryKey
	delete(values, idKey)

	return r.driver.Insert(values)
}

func (r repo[T, ID]) ByID(id ID) (T, error) {
	values, err := r.driver.ByID(id)
	var m T

	if err != nil {
		return m, err
	}

	m = r.toModel(values)

	return m, nil
}

func (r repo[T, ID]) Update(m T) error {
	values := r.toDriverValues(m)

	// get primary and unset it
	idKey := r.modelMeta.PrimaryKey
	id, ok := values[idKey].(ID)

	if !ok {
		panic("primary key type not matching ID")
	}

	delete(values, idKey)

	return r.driver.Update(values, id)
}

func (r repo[T, ID]) Delete(id ID) error {
	return r.driver.Delete(id)
}

func (r repo[T, ID]) toModel(values DriverValues) T {
	var m T

	v := reflect.ValueOf(&m).Elem()
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

	return m
}

func (r repo[T, ID]) toDriverValues(m T) DriverValues {
	values := make(DriverValues)
	v := reflect.ValueOf(&m).Elem()

	for _, fieldMeta := range r.modelMeta.Fields {
		values[fieldMeta.Key] = v.Field(fieldMeta.Index).Interface()
	}

	return values
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

func Repo[T any, ID int64](dbType string, dbPath string) Repository[T, ID] {
	if dbType == DRIVER_SQL {
		conn := NewSqliteConn(dbPath)
		driver := NewSqliteDriver[ID](conn)
		repo, err := NewRepo[T, ID](driver)

		if err != nil {
			panic(err)
		}

		return repo
	} else {
		panic("DRIVER NOT IMPLEMENTED")
	}
}
