package repository

import "testing"

func Test_Repository(t *testing.T) {
	t.Run("it works with driver and type", func(t *testing.T) {
		d := NewTestFakeDriver[int64]()
		r := NewRepo[TestUser, int64](d)

		r.Create(&TestUser{
			Name: "Gatis",
		})

		assertEqual(t, d.InsertCalled, false)
	})
}
