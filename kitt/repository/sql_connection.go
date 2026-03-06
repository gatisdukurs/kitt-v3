package repository

import (
	"context"
	"database/sql"
)

type SqlConnection interface {
	Exec(ctx context.Context, q string, args ...any) (sql.Result, error)
	Query(ctx context.Context, q string, args ...any) (*sql.Rows, error)
	QueryRow(ctx context.Context, q string, args ...any) *sql.Row
	Close() error
	WithDB(path string) SqlConnection
}
