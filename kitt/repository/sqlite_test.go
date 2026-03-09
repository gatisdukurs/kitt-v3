package repository

import (
	"os"
	"testing"
	"time"
)

func Test_Sqlite(t *testing.T) {
	dbPath := "testdata/test.sqlite"
	os.Remove(dbPath)

	t.Run("it connects and closes", func(t *testing.T) {
		conn := NewSqliteConn(dbPath)
		defer conn.Close()
	})

	t.Run("it executes", func(t *testing.T) {
		conn := NewSqliteConn(dbPath)
		defer conn.Close()

		q := `CREATE TABLE users (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name TEXT
			)`

		_, err := conn.Exec(t.Context(), q)
		assertEqual(t, err, nil)

		// Clean up
		q = `DROP TABLE users`

		_, err = conn.Exec(t.Context(), q)
		assertEqual(t, err, nil)
	})

	t.Run("it inserts values and queries it", func(t *testing.T) {
		conn := NewSqliteConn(dbPath)
		defer conn.Close()

		q := `CREATE TABLE todo (
				id TEXT PRIMARY KEY,
				completed BOOLEAN,
				desc TEXT
			)`

		_, err := conn.Exec(t.Context(), q)
		assertEqual(t, err, nil)

		// try some inserts
		docId := "1234"
		q = `INSERT INTO todo (id, completed, desc) VALUES (?, ?, ?)`
		res, err := conn.Exec(t.Context(), q, docId, false, "Buy a Milk!")

		rowID, err := res.LastInsertId()
		assertEqual(t, rowID, int64(1))

		// Lets query it now
		q = `SELECT * FROM todo WHERE id = ?`
		row := conn.QueryRow(t.Context(), q, docId)

		assertNotNil(t, row)

		var id string
		var completed bool
		var desc string

		err = row.Scan(&id, &completed, &desc)
		assertEqual(t, err, nil)
		assertEqual(t, id, docId)
		assertEqual(t, completed, false)
		assertEqual(t, desc, "Buy a Milk!")

		conn.Exec(t.Context(), `DROP TABLE todo`)
	})

	t.Run("it queries multiple things", func(t *testing.T) {
		conn := NewSqliteConn(dbPath)
		defer conn.Close()

		q := `CREATE TABLE calendar (
				id INTEGER PRIMARY KEY AUTOINCREMENT, 
				date DATE
			)`

		_, err := conn.Exec(t.Context(), q)
		assertEqual(t, err, nil)

		// insert some
		q = `INSERT INTO calendar (date) VALUES (?)`
		res, err := conn.Exec(t.Context(), q, time.Now())

		lastId, err := res.LastInsertId()

		assertEqual(t, err, nil)
		assertEqual(t, lastId, int64(1))

		// insert again

		q = `INSERT INTO calendar (date) VALUES (?)`
		res, err = conn.Exec(t.Context(), q, time.Now().Add(time.Hour*10))
		lastId, err = res.LastInsertId()

		assertEqual(t, err, nil)
		assertEqual(t, lastId, int64(2))

		// now the query part

		q = `SELECT * FROM calendar`

		rows, err := conn.Query(t.Context(), q)
		assertEqual(t, err, nil)

		defer rows.Close()

		data := []struct {
			Id   int64
			Date time.Time
		}{}

		for rows.Next() {
			var id int64
			var date time.Time

			err := rows.Scan(&id, &date)

			assertEqual(t, err, nil)

			data = append(data, struct {
				Id   int64
				Date time.Time
			}{
				Id:   id,
				Date: date,
			})
		}

		assertEqual(t, len(data), 2)

		assertEqual(t, data[0].Id, int64(1))
		assertEqual(t, data[1].Id, int64(2))

		conn.Exec(t.Context(), `DROP TABLE calendar`)
	})
}
