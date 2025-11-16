package database

import (
	"context"
	"database/sql"
)

type DatabaseQuery interface {
	QueryContext(cx context.Context, query string, args ...any) (*sql.Rows, error)
	ExecContext(cx context.Context, query string, args ...any) (*sql.Result, error)
	QueryRowContext(cx context.Context, query string, args ...any) *sql.Row
}
