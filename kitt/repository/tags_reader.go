package repository

import (
	"reflect"
	"strings"
)

type TagsMetadata struct {
	Attr  string
	Type  reflect.Type
	Index int
	Tags  []string
}

type TagsReader[T interface{}] interface {
	Read() []TagsMetadata
}

type reader[T interface{}] struct {
	tag string
}

func (r reader[T]) Read() []TagsMetadata {
	tags := []TagsMetadata{}

	var zero T
	t := reflect.TypeOf(zero)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		panic("repository: model must be a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get(r.tag)
		if tag == "" {
			continue
		}

		meta := TagsMetadata{
			Attr:  field.Name,
			Type:  field.Type,
			Index: i,
			Tags:  strings.Split(tag, ","),
		}
		tags = append(tags, meta)
	}

	return tags
}

func NewTagsReader[T interface{}](tag string) TagsReader[T] {
	return &reader[T]{
		tag: tag,
	}
}
