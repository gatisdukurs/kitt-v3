package repository

import (
	"fmt"
	"os"
	"testing"
)

func Test_Repository(t *testing.T) {
	dbPath := "test.db"
	os.Remove(dbPath)
	t.Run("it creates", func(t *testing.T) {
		d := NewTestFakeDriver[int]()
		d.InsertID = 1
		r, err := NewRepo[TestUser, int](d)

		id, err := r.Create(TestUser{
			ID:   10,
			Name: "Gatis",
		})

		assertEqual(t, id, d.InsertID)
		assertEqual(t, err, nil)
		assertEqual(t, d.InsertCalled, true)
		assertEqual(t, d.InsertValues["id"], 10)
		assertEqual(t, d.InsertValues["name"], "Gatis")
	})

	t.Run("it does by id", func(t *testing.T) {
		d := NewTestFakeDriver[int]()
		d.InsertID = 1
		r, err := NewRepo[TestUser, int](d)
		user, err := r.ByID(d.InsertID)

		assertEqual(t, err, nil)
		assertEqual(t, d.ByIDCalled, true)
		assertEqual(t, user.ID, 10)
		assertEqual(t, user.Name, "Gatis")
	})

	t.Run("it deletes", func(t *testing.T) {
		d := NewTestFakeDriver[int]()
		r, err := NewRepo[TestUser, int](d)

		err = r.Delete(10)

		assertEqual(t, err, nil)
		assertEqual(t, d.DeleteCalled, true)
	})

	t.Run("it updates", func(t *testing.T) {
		d := NewTestFakeDriver[int]()
		r, err := NewRepo[TestUser, int](d)

		user := TestUser{
			ID:   22,
			Name: "Gatis Dukurs",
		}

		err = r.Update(user)

		values := DriverValues{}
		values["name"] = "Gatis Dukurs"

		assertEqual(t, err, nil)
		assertEqual(t, d.UpdateCalled, true)
		assertEqual(t, d.UpdateValues, values)
		assertEqual(t, d.UpdateID, 22)
	})

	t.Run("it ensures collection exists", func(t *testing.T) {
		d := NewTestFakeDriver[int]()
		NewRepo[TestUser, int](d)

		assertEqual(t, d.EnsureCollectonCalled, true)

		d.EnsureCollectionError = fmt.Errorf("Collection creation error.")

		repo, err := NewRepo[TestUser, int](d)

		assertEqual(t, err, d.EnsureCollectionError)
		assertEqual(t, repo, nil)
	})

	t.Run("repo factory works", func(t *testing.T) {
		repo := Repo[TestUser, int64](DRIVER_SQL, dbPath)

		if _, ok := repo.(Repository[TestUser, int64]); !ok {
			t.Fatalf("repo factory not working")
		}
	})

	t.Run("it returns all", func(t *testing.T) {

	})
}
