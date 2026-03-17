package repository

import (
	"context"
	"fmt"
	"reflect"
	"slices"
	"strings"
	"time"
)

func tableType(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		return "INTEGER"

	case reflect.String:
		return "TEXT"

	case reflect.Bool:
		return "BOOLEAN"

	case reflect.Float32, reflect.Float64:
		return "REAL"

	case reflect.Struct:
		if t == reflect.TypeOf(time.Time{}) {
			return "DATE"
		}
	}

	return "TEXT"
}

func parseDefaultFlag(flag string) string {
	const prefix = "default:"

	if !strings.HasPrefix(flag, prefix) {
		return ""
	}

	val := strings.TrimPrefix(flag, prefix)
	if val == "" {
		return ""
	}

	return " DEFAULT " + val
}

type sqliteDriver[ID int64] struct {
	conn      SqlConnection
	modelMeta ModelMeta
}

func (d sqliteDriver[ID]) Query() FindQuery {
	return SELECT(d.modelMeta.Collection)
}

func (d sqliteDriver[ID]) Insert(values DriverValues) (ID, error) {
	table := d.modelMeta.Collection
	q := INSERT(table)

	columns := []string{}
	row := []any{}
	for k, v := range values {
		columns = append(columns, k)
		row = append(row, v)
	}

	q.Columns(columns...)
	q.Row(row...)

	sql, args := q.Build()

	res, err := d.conn.Exec(context.Background(), sql, args...)

	var zero ID
	if err != nil {
		return zero, err
	}

	raw, err := res.LastInsertId()

	if err != nil {
		return zero, err
	}

	return ID(raw), nil
}

func (d sqliteDriver[ID]) Update(values DriverValues, id ID) error {
	table := d.modelMeta.Collection
	q := UPDATE(table)
	for column, value := range values {
		q.Set(column, value)
	}

	q.Where(Eq("id", id))

	sql, args := q.Build()

	_, err := d.conn.Exec(context.Background(), sql, args...)

	return err
}

func (d sqliteDriver[ID]) Delete(id ID) error {
	table := d.modelMeta.Collection
	q := DELETE(table)
	q.Where(Eq("id", id))
	sql, args := q.Build()

	_, err := d.conn.Exec(context.Background(), sql, args...)

	if err != nil {
		return err
	}

	return nil
}

func (d sqliteDriver[ID]) ByID(id ID) (DriverValues, error) {
	table := d.modelMeta.Collection
	q := SELECT(table)
	keys := []string{}
	primaryKey := "id"

	for _, field := range d.modelMeta.Fields {
		keys = append(keys, field.Key)
		if slices.Contains(field.Flags, "pk") {
			primaryKey = field.Key
		}
	}

	q.Columns(keys...)
	q.Where(Eq(primaryKey, id))

	return d.First(q)
}

func (d sqliteDriver[ID]) Find(q FindQuery) ([]DriverValues, error) {
	keys := []string{}
	values := []DriverValues{}
	sql, args := q.Build()

	for _, field := range d.modelMeta.Fields {
		keys = append(keys, field.Key)
	}

	rows, err := d.conn.Query(context.Background(), sql, args...)

	if err != nil {
		return values, err
	}

	for rows.Next() {
		v := make(DriverValues)
		scanValues := make([]any, len(keys))
		scanPtrs := make([]any, len(keys))

		for i := range scanValues {
			scanPtrs[i] = &scanValues[i]
		}

		err := rows.Scan(scanPtrs...)
		if err != nil {
			return values, err
		}

		for i, key := range keys {
			value := scanValues[i]

			if b, ok := value.([]byte); ok {
				v[key] = string(b)
			} else {
				v[key] = value
			}
		}

		values = append(values, v)
	}

	return values, nil
}

func (d sqliteDriver[ID]) First(q FindQuery) (DriverValues, error) {
	q.Limit(1)
	values, err := d.Find(q)

	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, nil
	}

	return values[0], err
}

func (sql sqliteDriver[ID]) DropCollection() error {
	table := sql.modelMeta.Collection
	q := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, table)
	_, err := sql.conn.Exec(context.Background(), q)
	return err
}

func (sql *sqliteDriver[ID]) WithModelMeta(modelMeta ModelMeta) Driver[ID] {
	sql.modelMeta = modelMeta
	return sql
}

func (sql sqliteDriver[ID]) CreateCollection() error {
	modelMeta := sql.modelMeta
	ctx := context.Background()

	tableFields := []string{}

	for _, f := range modelMeta.Fields {
		format := `%s %s`

		for _, flag := range f.Flags {
			switch {
			case flag == "pk":
				format += ` PRIMARY KEY`
			case flag == "auto":
				format += ` AUTOINCREMENT`
			case flag == "unique":
				format += ` UNIQUE`
			case flag == "notnull":
				format += ` NOT NULL`
			case strings.HasPrefix(flag, "default:"):
				format += parseDefaultFlag(flag)
			case strings.HasPrefix(flag, "check:"):
				format += ` CHECK(` + strings.TrimPrefix(flag, "check:") + `)`
			case strings.HasPrefix(flag, "references:"):
				format += ` REFERENCES ` + strings.TrimPrefix(flag, "references:")
			case strings.HasPrefix(flag, "ondelete:"):
				format += ` ON DELETE ` + strings.TrimPrefix(flag, "ondelete:")
			case strings.HasPrefix(flag, "onupdate:"):
				format += ` ON UPDATE ` + strings.TrimPrefix(flag, "onupdate:")
			}
		}

		tField := fmt.Sprintf(format, f.Key, tableType(f.Type))
		tableFields = append(tableFields, tField)
	}

	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s)`,
		modelMeta.Collection,
		strings.Join(tableFields, ","),
	)

	_, err := sql.conn.Exec(ctx, q)

	return err
}

func NewSqliteDriver[ID int64](conn SqlConnection) Driver[ID] {
	return &sqliteDriver[ID]{
		conn: conn,
	}
}
