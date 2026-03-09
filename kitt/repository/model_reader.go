package repository

import (
	"reflect"
	"strings"
)

type ModelFieldMeta struct {
	Attr  string
	Index int
	Type  reflect.Type
	Key   string
	Flags []string
}

type ModelMeta struct {
	Collection string
	Fields     []ModelFieldMeta
}

type ModelReader[T interface{}] interface {
	Read() ModelMeta
}

type modelReader[T interface{}] struct {
	tag string
}

func (mr modelReader[T]) Read() ModelMeta {
	tagReader := NewTagsReader[T](mr.tag)
	tagsMeta := tagReader.Read()

	var zero T
	t := reflect.TypeOf(zero)
	table := strings.ToLower(t.Name())

	fields := []ModelFieldMeta{}

	for _, tag := range tagsMeta {
		key := tag.Tags[0]

		fields = append(fields, ModelFieldMeta{
			Attr:  tag.Attr,
			Key:   key,
			Type:  tag.Type,
			Index: tag.Index,
			Flags: tag.Tags[1:],
		})
	}

	return ModelMeta{
		Collection: table,
		Fields:     fields,
	}
}

func NewModelReader[T interface{}](tag string) ModelReader[T] {
	return &modelReader[T]{
		tag: tag,
	}
}
