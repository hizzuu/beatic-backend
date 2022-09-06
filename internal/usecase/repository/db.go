package repository

import (
	"context"
	"database/sql"
)

type DB interface {
	DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}
