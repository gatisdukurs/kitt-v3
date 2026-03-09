package repository

import (
	"reflect"
	"testing"
)

func Test_Model_Reader(t *testing.T) {
	t.Run("it reads", func(t *testing.T) {
		mr := NewModelReader[TestUser]("db")

		assertEqual(t, mr.Read(), ModelMeta{
			Collection: "testuser",
			Fields: []ModelFieldMeta{
				{Attr: "ID", Key: "id", Type: reflect.TypeOf(int(0)), Index: 0, Flags: []string{"pk", "required"}},
				{Attr: "Name", Key: "name", Type: reflect.TypeOf(""), Index: 1, Flags: []string{}},
			},
		})
	})
}
