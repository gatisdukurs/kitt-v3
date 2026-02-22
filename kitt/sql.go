package kitt

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

var dbInstance SQLDB

type sqlDB struct {
	db *sql.DB
}

func InitSQL() SQLDB {
	if dbInstance == nil {
		dbInstance = &sqlDB{}
	}
	return dbInstance
}

func SQL() SQLDB {
	return dbInstance
}

type SQLDB interface {
	Exec(ctx context.Context, q string, args ...any) (sql.Result, error)
	Query(ctx context.Context, q string, args ...any) (*sql.Rows, error)
	QueryRow(ctx context.Context, q string, args ...any) *sql.Row
	Close() error
	WithSQLite(path string) SQLDB
}

func (d *sqlDB) Exec(ctx context.Context, q string, args ...any) (sql.Result, error) {
	return d.db.ExecContext(ctx, q, args...)
}

func (d *sqlDB) Query(ctx context.Context, q string, args ...any) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, q, args...)
}

func (d *sqlDB) QueryRow(ctx context.Context, q string, args ...any) *sql.Row {
	return d.db.QueryRowContext(ctx, q, args...)
}

func (d *sqlDB) Close() error {
	return d.db.Close()
}

func (d *sqlDB) WithSQLite(path string) SQLDB {
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
