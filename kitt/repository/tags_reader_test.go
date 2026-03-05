package repository

import "testing"

func Test_Tags_Reader(t *testing.T) {
	t.Run("it reads tags", func(t *testing.T) {
		r := NewTagsReader[struct {
			ID   string `db:"id,pk,required"`
			Name string `db:"name"`
		}]("db")

		want := []TagsMetadata{
			{Attr: "ID", Tags: []string{"id", "pk", "required"}},
			{Attr: "Name", Tags: []string{"name"}},
		}

		assertEqual(t, r.Read(), want)
	})
}
