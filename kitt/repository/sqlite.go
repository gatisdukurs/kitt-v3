package repository

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

var connInstances map[string]*sqlite

type sqlite struct {
	db *sql.DB
}

func (d *sqlite) Exec(ctx context.Context, q string, args ...any) (sql.Result, error) {
	return d.db.ExecContext(ctx, q, args...)
}

func (d *sqlite) Query(ctx context.Context, q string, args ...any) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, q, args...)
}

func (d *sqlite) QueryRow(ctx context.Context, q string, args ...any) *sql.Row {
	return d.db.QueryRowContext(ctx, q, args...)
}

func (d *sqlite) Close() error {
	return d.db.Close()
}

func (d *sqlite) WithDB(path string) SqlConnection {
	conn, err := sql.Open("sqlite", "file:"+path+"?_pragma=foreign_keys(1)")

	if err != nil {
		panic(err)
	}

	conn.SetMaxOpenConns(1)
	conn.SetMaxIdleConns(1)
	conn.SetConnMaxLifetime(0)

	d.db = conn

	return d
}

func NewSqliteConn(path string) SqlConnection {
	if connInstances == nil {
		connInstances = make(map[string]*sqlite)
	}

	if conn, ok := connInstances[path]; ok {
		return conn
	} else {
		conn := &sqlite{}
		conn.WithDB(path)

		return conn
	}
}
