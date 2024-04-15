package db

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Client - connection to db
// type Client interface {
//	DB() DB
//	Close() error
//}

// DB - functions for working with db
type DB interface {
	SQLExecer
	Pinger
	Transaction
	Close() error
}

// Query - query to db with name and raw query
type Query struct {
	Name     string
	QueryRaw string
	Args     []interface{}
}

func NewQuery(name string, queryRaw string, args []interface{}) Query {
	return Query{
		Name:     name,
		QueryRaw: queryRaw,
		Args:     args,
	}
}

// SQLExecer - functions for working with db
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer - functions for working with db
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, q Query) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query) pgx.Row
}

type Pinger interface {
	Ping(ctx context.Context) error
}

// type Transaction interface {
//	StartTransaction() (pgx.Tx, error)
//}

type Transaction interface {
	StartTransaction(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}
