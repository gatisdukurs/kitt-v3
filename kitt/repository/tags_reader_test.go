package repository

import (
	"reflect"
	"testing"
)

func Test_Tags_Reader(t *testing.T) {
	t.Run("it reads tags", func(t *testing.T) {
		r := NewTagsReader[TestUser]("db")

		want := []TagsMetadata{
			{Attr: "ID", Type: reflect.TypeOf(int(0)), Index: 0, Tags: []string{"id", "pk", "required"}},
			{Attr: "Name", Type: reflect.TypeOf(""), Index: 1, Tags: []string{"name"}},
		}

		assertEqual(t, r.Read(), want)
	})
}
