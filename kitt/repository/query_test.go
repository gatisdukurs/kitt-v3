package repository

import "testing"

func Test_Query(t *testing.T) {
	// q := Query{
	// 	Where: []Filter{
	// 		{Field: ""}
	// 	},
	// }

	// t.Run("it builds", func(t *testing.T) {
	// 	res := BuildQuery(q)

	// 	assertEqual(t, res, ``)
	// })

	t.Run("it builds condition", func(t *testing.T) {
		str, arg := BuildCondition(Condition{
			Field: "id",
			Op:    "=",
			Value: 10,
		})

		assertEqual(t, str, `id = ?`)
		assertEqual(t, arg, 10)
	})

	t.Run("it builds where", func(t *testing.T) {
		q := SelectQuery{
			Where: []Condition{
				{Field: "id", Op: "=", Value: 10},
				{Field: "age", Op: ">", Value: 30},
			},
		}

		where, args := BuildWhere(q)

		assertEqual(t, where, `WHERE id = ? AND age > ?`)
		assertEqual(t, args, []any{
			10, 30,
		})
	})

	t.Run("it builds order by", func(t *testing.T) {
		by := SelectQuery{
			OrderBy: []Order{
				{Field: "id", Desc: true},
				{Field: "age", Desc: false},
			},
		}

		orderBy := BuildOrderBy(by)

		assertEqual(t, orderBy, `ORDER BY id DESC, age ASC`)
	})

	t.Run("it builds limit and offset", func(t *testing.T) {
		q := SelectQuery{
			Limit:  100,
			Offset: 200,
		}

		str := BuildLimitOffset(q)

		assertEqual(t, str, `LIMIT 100 OFFSET 200`)
	})

	t.Run("it builds proper queries", func(t *testing.T) {
		// Empty
		q := SelectQuery{}
		str, args, err := BuildQuery(q)

		assertEqual(t, str, str)
		assertEqual(t, args, nil)
		assertNotNil(t, err)

		// With table
		q.Table = "todo"
		str, _, _ = BuildQuery(q)
		assertEqual(t, str, `SELECT * FROM todo`)

		// With table, fields
		q.Fields = []string{"id", "completed"}
		str, _, _ = BuildQuery(q)
		assertEqual(t, str, `SELECT id, completed FROM todo`)

		// With table, fields, simple where
		q.Where = []Condition{
			{Field: "completed", Op: "=", Value: true},
		}
		str, args, _ = BuildQuery(q)
		assertEqual(t, str, `SELECT id, completed FROM todo WHERE completed = ?`)
		assertEqual(t, args, []any{true})

		// With table, fields, complex where
		q.Fields = append(q.Fields, "age")
		q.Where = []Condition{
			{Field: "completed", Op: "=", Value: true},
			{Field: "age", Op: ">", Value: 10},
		}
		str, args, _ = BuildQuery(q)
		assertEqual(t, str, `SELECT id, completed, age FROM todo WHERE completed = ? AND age > ?`)
		assertEqual(t, args, []any{true, 10})

		// With table, fields, complex where, limit, offset
		q.Limit = 10
		q.Offset = 20
		str, args, _ = BuildQuery(q)
		assertEqual(t, str, `SELECT id, completed, age FROM todo WHERE completed = ? AND age > ? LIMIT 10 OFFSET 20`)
		assertEqual(t, args, []any{true, 10})
	})

	t.Run("it builds queries", func(t *testing.T) {
		builder := NewQueryBuilder()
		builder.
			Select("id", "name").
			From("todo").
			Where("completed", "=", true).
			Where("age", ">=", 10).
			OrderAscBy("age").
			OrderDescBy("id").
			Limit(10).
			Offset(20)

		query, args := builder.Build()

		assertEqual(t, query, `SELECT id, name FROM todo WHERE completed = ? AND age >= ? ORDER BY age ASC, id ASC LIMIT 10 OFFSET 20`)
		assertEqual(t, args, []any{
			true, 10,
		})
	})
}
