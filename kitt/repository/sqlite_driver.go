package repository

import (
	"context"
	"database/sql"
	"errors"
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
	conn      SqlConnection
	modelMeta ModelMeta
}

func (sql sqliteDriver[ID]) Insert(values DriverValues) (ID, error) {
	table := sql.modelMeta.Collection

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

func (sql sqliteDriver[ID]) Update(values DriverValues, id ID) error {
	set := []string{}
	args := []interface{}{}

	for k, v := range values {
		set = append(set, fmt.Sprintf(`%s=?`, k))
		args = append(args, v)
	}

	args = append(args, id)

	q := fmt.Sprintf(`UPDATE %s SET %s WHERE id = ?`, sql.modelMeta.Collection, strings.Join(set, ","))

	_, err := sql.conn.Exec(context.Background(), q, args...)

	return err
}

func (sql sqliteDriver[ID]) Delete(id ID) error {
	q := fmt.Sprintf(`DELETE FROM %s WHERE id = ?`, sql.modelMeta.Collection)

	_, err := sql.conn.Exec(context.Background(), q, id)

	if err != nil {
		return err
	}

	return nil
}

func (s sqliteDriver[ID]) ByID(id ID) (DriverValues, error) {
	table := s.modelMeta.Collection
	keys := []string{}

	for _, field := range s.modelMeta.Fields {
		keys = append(keys, field.Key)
	}

	q := fmt.Sprintf(`SELECT %s FROM  %s WHERE id = ?`, strings.Join(keys, ","), table)
	row := s.conn.QueryRow(context.Background(), q, id)

	scanValues := make([]any, len(keys))
	scanPtrs := make([]any, len(keys))

	for i := range scanValues {
		scanPtrs[i] = &scanValues[i]
	}

	err := row.Scan(scanPtrs...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	// map values
	values := make(DriverValues)
	for i, key := range keys {
		v := scanValues[i]

		// sqlite often returns []byte for TEXT
		if b, ok := v.([]byte); ok {
			values[key] = string(b)
		} else {
			values[key] = v
		}
	}

	if len(values) == 0 {
		return nil, nil
	}

	return values, nil
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
