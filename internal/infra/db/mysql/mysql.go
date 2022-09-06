package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/hizzuu/beatic-backend/conf"
	"github.com/jmoiron/sqlx"
)

const txCtxKey = "tx"

type db struct {
	conn *sqlx.DB
}

type DB interface {
	DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

func New() (*db, error) {
	mysqlConf := &mysql.Config{
		User:                 conf.C.DB.User,
		Passwd:               conf.C.DB.Pass,
		Net:                  conf.C.DB.Net,
		Addr:                 conf.C.DB.Host + ":" + conf.C.DB.Port,
		DBName:               conf.C.DB.Name,
		ParseTime:            conf.C.DB.Parsetime,
		Loc:                  time.Local,
		AllowNativePasswords: conf.C.DB.AllowNativePasswords,
	}

	conn, err := sqlx.Open(conf.C.DB.Dbms, mysqlConf.FormatDSN())
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
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
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
