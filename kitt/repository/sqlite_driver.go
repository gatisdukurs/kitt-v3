package repository

import (
	"context"
	"fmt"
	"reflect"
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
	conn SqlConnection
}

func (sql sqliteDriver[ID]) Insert(table string, values DriverValues) (ID, error) {
	var zero ID
	args := []interface{}{}
	keys := []string{}
	bindings := []string{}

	for k, v := range values {
		keys = append(keys, k)
		args = append(args, v)
		bindings = append(bindings, "?")
	}

	q := fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s)`, table, strings.Join(keys, ","), strings.Join(bindings, ","))

	res, err := sql.conn.Exec(context.Background(), q, args...)

	if err != nil {
		return zero, err
	}

	raw, err := res.LastInsertId()

	if err != nil {
		return zero, err
	}

	return ID(raw), nil
}

func (sql sqliteDriver[ID]) Update(table string, values DriverValues, id ID) error {
	return nil
}

func (sql sqliteDriver[ID]) Delete(table string, id ID) error {
	return nil
}

func (sql sqliteDriver[ID]) ByID(table string, id ID) (DriverValues, error) {
	values := make(DriverValues)
	return values, nil
}

func (sql sqliteDriver[ID]) DropCollection(table string) error {
	q := fmt.Sprintf(`DROP TABLE IF EXISTS %s`, table)
	_, err := sql.conn.Exec(context.Background(), q)
	return err
}

func (sql sqliteDriver[ID]) EnsureCollectionExists(modelMeta ModelMeta) error {
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
