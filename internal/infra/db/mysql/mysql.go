package mysql

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hizzuu/beatic-backend/conf"
	"github.com/jmoiron/sqlx"
)

type db struct {
	conn *sqlx.DB
}

type DB interface {
	DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type key string

const (
	txCtxKey key = "tx"
)

func New() (*db, error) {
	conn, err := sqlx.Open("mysql", conf.C.DB.DSN)
	if err != nil {
		return nil, err
	}

	return &db{conn: conn}, nil
}

func (db *db) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := db.conn.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	v, err := f(context.WithValue(ctx, txCtxKey, tx))
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}

		return nil, err
	}

	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}

		return nil, err
	}

	return v, nil
}

func (db *db) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if tx, ok := ctx.Value(txCtxKey).(*sql.Tx); ok {
		return tx.PrepareContext(ctx, query)
	}
	return db.conn.PrepareContext(ctx, query)
}
