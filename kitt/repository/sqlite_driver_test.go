package repository

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_Sqlite_Driver(t *testing.T) {
	dbPath := "testdata/test.sqlite"
	os.Remove(dbPath)

	/*
	   Should support

	   ID        int64     `db:"pk,auto"`
	   Title     string    `db:"notnull,default:'Untitled'"`
	   Done      bool      `db:"notnull,default:0"`
	   CreatedAt time.Time `db:"default:CURRENT_TIMESTAMP"`
	   UserID    int64     `db:"references:users(id),ondelete:CASCADE,onupdate:CASCADE"`
	   Age       int64     `db:"check:age >= 0"`

	*/

	t.Run("it ensures table and drops it", func(t *testing.T) {
		conn := NewSqliteConn(dbPath)
		driver := NewSqliteDriver(conn)
		modelMeta := ModelMeta{
			Collection: "todo",
			Fields: []ModelFieldMeta{
				{Key: "id", Type: reflect.TypeOf(int64(0)), Flags: []string{"pk", "auto"}},
				{Key: "completed", Type: reflect.TypeOf(true), Flags: []string{"notnull", "default:false"}},
				{Key: "desc", Type: reflect.TypeOf(""), Flags: []string{}},
				{Key: "created_at", Type: reflect.TypeOf(time.Now()), Flags: []string{}},
				{Key: "slug", Type: reflect.TypeOf(""), Flags: []string{"unique"}},
				{Key: "test_default", Type: reflect.TypeOf(""), Flags: []string{"default:'Title'"}},
				{Key: "test_default_timestamp", Type: reflect.TypeOf(time.Now()), Flags: []string{"default:CURRENT_TIMESTAMP"}},
				{Key: "test_references", Type: reflect.TypeOf(time.Now()), Flags: []string{"references:users(id)", "ondelete:CASCADE", "onupdate:CASCADE"}},
				{Key: "age", Type: reflect.TypeOf(int64(0)), Flags: []string{"check:age >= 0"}},
			},
		}

		driver.WithModelMeta(modelMeta)
		err := driver.CreateCollection()

		assertEqual(t, err, nil)

		q := `SELECT sql FROM sqlite_master WHERE type='table' AND name='todo'`
		row := conn.QueryRow(t.Context(), q)

		assertNotNil(t, row)

		sql := ""
		row.Scan(&sql)

		assertEqual(t, err, nil)
		assertEqual(t, sql, `CREATE TABLE todo (id INTEGER PRIMARY KEY AUTOINCREMENT,completed BOOLEAN NOT NULL DEFAULT false,desc TEXT,created_at DATE,slug TEXT UNIQUE,test_default TEXT DEFAULT 'Title',test_default_timestamp DATE DEFAULT CURRENT_TIMESTAMP,test_references DATE REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,age INTEGER CHECK(age >= 0))`)

		driver.DropCollection()

		q = `SELECT sql FROM sqlite_master WHERE type='table' AND name='todo'`
		row = conn.QueryRow(t.Context(), q)

		sql = ""
		row.Scan(&sql)

		assertEqual(t, sql, "")
	})

	t.Run("it inserts and gets by ID and deletes", func(t *testing.T) {
		// Setup
		conn := NewSqliteConn(dbPath)
		driver := NewSqliteDriver(conn)

		modelMeta := ModelMeta{
			Collection: "users",
			Fields: []ModelFieldMeta{
				{Key: "id", Type: reflect.TypeOf(int64(0)), Flags: []string{"pk", "auto"}},
				{Key: "username", Type: reflect.TypeOf(""), Flags: []string{"notnull", "unique"}},
				{Key: "password", Type: reflect.TypeOf(""), Flags: []string{"notnull"}},
			},
		}

		driver.WithModelMeta(modelMeta)
		driver.CreateCollection()
		defer driver.DropCollection()

		// Insert
		iValues := make(DriverValues)
		iValues["username"] = "dumdum"
		iValues["password"] = "secret"

		id, err := driver.Insert(iValues)

		assertEqual(t, id, int64(1))
		assertEqual(t, err, nil)

		// Get by ID
		qValues, err := driver.ByID(id)

		assertEqual(t, err, nil)
		assertEqual(t, qValues["username"], "dumdum")
		assertEqual(t, qValues["password"], "secret")

		// Delete
		err = driver.Delete(id)
		assertEqual(t, err, nil)

		// Check on it
		qValues, err = driver.ByID(id)

		assertEqual(t, err, nil)
		assertEqual(t, qValues, nil)
	})

	t.Run("it updates", func(t *testing.T) {
		conn := NewSqliteConn(dbPath)
		driver := NewSqliteDriver(conn)

		modelMeta := ModelMeta{
			Collection: "users",
			Fields: []ModelFieldMeta{
				{Key: "id", Type: reflect.TypeOf(int64(0)), Flags: []string{"pk", "auto"}},
				{Key: "username", Type: reflect.TypeOf(""), Flags: []string{"notnull", "unique"}},
			},
		}

		driver.WithModelMeta(modelMeta)
		driver.CreateCollection()
		defer driver.DropCollection()

		// Insert
		iValues := make(DriverValues)
		iValues["username"] = "dumdum"

		id, err := driver.Insert(iValues)

		assertEqual(t, err, nil)

		// Update
		uValues := make(DriverValues)
		uValues["username"] = "Gatis"

		err = driver.Update(uValues, id)

		assertEqual(t, err, nil)

		// Query to see if updated
		qValues, err := driver.ByID(id)

		assertEqual(t, qValues["username"], "Gatis")
	})

	t.Run("it supports Find and First", func(t *testing.T) {
		conn := NewSqliteConn(dbPath)
		driver := NewSqliteDriver(conn)

		modelMeta := ModelMeta{
			Collection: "users",
			Fields: []ModelFieldMeta{
				{Key: "id", Type: reflect.TypeOf(int64(0)), Flags: []string{"pk", "auto"}},
				{Key: "username", Type: reflect.TypeOf(""), Flags: []string{"notnull", "unique"}},
			},
		}

		driver.WithModelMeta(modelMeta)
		driver.CreateCollection()
		defer driver.DropCollection()

		// Insert
		iValues := make(DriverValues)
		iValues["username"] = "dumdum"

		driver.Insert(iValues)
		iValues["username"] = "Gatis"
		driver.Insert(iValues)

		// Now query
		query := NewQueryBuilder()
		query.Select("id", "username")
		query.From("users")
		values, err := driver.Find(query)

		assertEqual(t, err, nil)
		assertEqual(t, len(values), 2)
		assertEqual(t, values[0]["username"], "dumdum")
		assertEqual(t, values[1]["username"], "Gatis")

		vs, err := driver.First(query)

		assertEqual(t, err, nil)
		assertEqual(t, vs["username"], "dumdum")
	})
}
