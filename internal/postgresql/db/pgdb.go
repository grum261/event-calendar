package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type pgdb interface {
	Begin(context.Context) (pgx.Tx, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
}

type Queries struct {
	db pgdb
}

func NewQueries(db pgdb) *Queries {
	return &Queries{
		db: db,
	}
}

func (q *Queries) Begin(ctx context.Context) (pgx.Tx, error) {
	return q.db.Begin(ctx)
}

func (q *Queries) WithTx(tx pgx.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}
