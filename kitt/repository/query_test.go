package repository

import "testing"

func Test_SELECT(t *testing.T) {
	t.Run("SELECT id, name FROM users", func(t *testing.T) {
		q := SELECT("users").Columns("id", "name")
		sql, _ := q.Build()
		assertEqual(t, sql, `SELECT id, name FROM users`)
	})

	t.Run("SELECT * FROM users", func(t *testing.T) {
		q := SELECT("users")
		sql, _ := q.Build()
		assertEqual(t, sql, `SELECT * FROM users`)
	})

	t.Run("it supports where", func(t *testing.T) {
		q := SELECT("users")
		q.Where(Gt("id", 1))

		sql, args := q.Build()
		assertEqual(t, sql, `SELECT * FROM users WHERE id > ?`)
		assertEqual(t, args, []any{1})
	})

	t.Run("It supports order by", func(t *testing.T) {
		q := SELECT("users")
		q.OrderBy("id", DESC)
		q.OrderBy("age", ASC)

		sql, args := q.Build()

		assertEqual(t, sql, `SELECT * FROM users ORDER BY id DESC, age ASC`)
		assertEqual(t, args, []any{})
	})

	t.Run("It supports limit", func(t *testing.T) {
		q := SELECT("users")
		q.Limit(10)

		sql, args := q.Build()

		assertEqual(t, sql, `SELECT * FROM users LIMIT 10`)
		assertEqual(t, args, []any{})
	})

	t.Run("It supports offset", func(t *testing.T) {
		q := SELECT("users")
		q.Offset(20)

		sql, args := q.Build()

		assertEqual(t, sql, `SELECT * FROM users OFFSET 20`)
		assertEqual(t, args, []any{})
	})
}

func Test_INSERT(t *testing.T) {
	t.Run("INSERT INTO users (name, age) VALUES (?, ?)", func(t *testing.T) {
		q := INSERT("users").Columns("name", "age")
		q.Row("Gatis", 18)

		sql, args := q.Build()
		assertEqual(t, sql, `INSERT INTO users (name, age) VALUES (?, ?)`)
		assertEqual(t, len(args), 2)
	})

	t.Run("It supports rows", func(t *testing.T) {
		q := INSERT("users").Columns("name", "age")
		q.Row("Gatis", 42)
		q.Row("Kristine", 18)

		sql, args := q.Build()
		assertEqual(t, sql, `INSERT INTO users (name, age) VALUES (?, ?), (?, ?)`)
		assertEqual(t, len(args), 4)
	})
}

func Test_UPDATE(t *testing.T) {
	t.Run("UPDATE users SET age = ? WHERE id = ?", func(t *testing.T) {
		q := UPDATE("users")
		q.Set("age", 18)
		q.Where(Eq("id", 1))

		sql, args := q.Build()

		assertEqual(t, sql, `UPDATE users SET age = ? WHERE id = ?`)
		assertEqual(t, args, []any{18, 1})
	})
}

func Test_DELETE(t *testing.T) {
	t.Run("DELETE FROM users WHERE id = ?", func(t *testing.T) {
		q := DELETE("users")
		q.Where(Eq("id", 1))

		sql, args := q.Build()
		assertEqual(t, sql, `DELETE FROM users WHERE id = ?`)
		assertEqual(t, len(args), 1)
	})
}

func Test_GroupClause(t *testing.T) {
	t.Run("AND", func(t *testing.T) {
		sql, args := And(
			Lt("id", 10),
			Gt("id", 1),
		).Build()

		assertEqual(t, sql, `(id < ? AND id > ?)`)
		assertEqual(t, args, []any{10, 1})
	})

	t.Run("OR", func(t *testing.T) {
		sql, args := Or(
			Lt("id", 10),
			Gt("id", 1),
		).Build()

		assertEqual(t, sql, `(id < ? OR id > ?)`)
		assertEqual(t, args, []any{10, 1})
	})

	t.Run("AND OR", func(t *testing.T) {
		sql, args := Or(
			Lt("id", 10),
			Gt("id", 1),
			And(Eq("age", 10), Ne("age", 20)),
		).Build()

		assertEqual(t, sql, `(id < ? OR id > ? OR (age = ? AND age != ?))`)
		assertEqual(t, args, []any{10, 1, 10, 20})
	})
}

func Test_Expressions(t *testing.T) {
	t.Run("=", func(t *testing.T) {
		sql, args := Eq("id", 1).Build()
		assertEqual(t, sql, "id = ?")
		assertEqual(t, args, []any{1})
	})

	t.Run("!=", func(t *testing.T) {
		sql, args := Ne("id", 1).Build()
		assertEqual(t, sql, "id != ?")
		assertEqual(t, args, []any{1})
	})

	t.Run(">", func(t *testing.T) {
		sql, args := Gt("id", 1).Build()
		assertEqual(t, sql, "id > ?")
		assertEqual(t, args, []any{1})
	})

	t.Run(">=", func(t *testing.T) {
		sql, args := Gte("id", 1).Build()
		assertEqual(t, sql, "id >= ?")
		assertEqual(t, args, []any{1})
	})

	t.Run("<", func(t *testing.T) {
		sql, args := Lt("id", 1).Build()
		assertEqual(t, sql, "id < ?")
		assertEqual(t, args, []any{1})
	})

	t.Run("<=", func(t *testing.T) {
		sql, args := Lte("id", 1).Build()
		assertEqual(t, sql, "id <= ?")
		assertEqual(t, args, []any{1})
	})

	t.Run("LIKE", func(t *testing.T) {
		sql, args := Like("id", 1).Build()
		assertEqual(t, sql, "id LIKE ?")
		assertEqual(t, args, []any{1})
	})

	t.Run("IS NULL", func(t *testing.T) {
		sql, args := IsNull("id").Build()
		assertEqual(t, sql, "id IS NULL")
		assertEqual(t, args, nil)
	})

	t.Run("IS NOT NULL", func(t *testing.T) {
		sql, args := IsNotNull("id").Build()
		assertEqual(t, sql, "id IS NOT NULL")
		assertEqual(t, args, nil)
	})
}
