package pg

import (
	"auth/internal/client/db"
	"auth/internal/logger"
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type pg struct {
	log logger.Logger
	dbc *pgxpool.Pool
}

func NewDB(ctx context.Context, dsn string, log logger.Logger) (db.DB, error) {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	return &pg{
		dbc: pool,
		log: log,
	}, nil
}

func (p *pg) ScanOneContext(ctx context.Context, dest interface{}, q db.Query) error {
	// p.logQuery(q)
	rows, err := p.QueryContext(ctx, q)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, rows)
}

func (p *pg) ScanAllContext(ctx context.Context, dest interface{}, q db.Query) error {
	// p.logQuery(q)
	rows, err := p.QueryContext(ctx, q)
	if err != nil {
		return err
	}
	return pgxscan.ScanAll(dest, rows)
}

func (p *pg) QueryContext(ctx context.Context, q db.Query) (pgx.Rows, error) {
	p.logQuery(q)

	if tx, ok := ctx.Value("tx").(pgx.Tx); ok {
		return tx.Query(ctx, q.QueryRaw, q.Args...)
	}

	return p.dbc.Query(ctx, q.QueryRaw, q.Args...)
}

func (p *pg) QueryRowContext(ctx context.Context, q db.Query) pgx.Row {
	p.logQuery(q)

	if tx, ok := ctx.Value("tx").(pgx.Tx); ok {
		return tx.QueryRow(ctx, q.QueryRaw, q.Args...)
	}

	return p.dbc.QueryRow(ctx, q.QueryRaw, q.Args...)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.dbc.Ping(ctx)
}

func (p *pg) Close() error {
	p.dbc.Close()
	return nil
}

func (p *pg) ExecContext(ctx context.Context, q db.Query) (pgconn.CommandTag, error) {
	p.logQuery(q)

	if tx, ok := ctx.Value("tx").(pgx.Tx); ok {
		return tx.Exec(ctx, q.QueryRaw, q.Args...)
	}

	return p.dbc.Exec(ctx, q.QueryRaw, q.Args...)
}

func (p *pg) logQuery(q db.Query) {
	l := fmt.Sprintf("sql: %s, query: %s, args: %v", q.Name, q.QueryRaw, q.Args)

	p.log.Info(l)
}

func (p *pg) StartTransaction(ctx context.Context) (context.Context, error) {
	tx, err := p.dbc.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, "tx", tx), nil
}

func (p *pg) Rollback(ctx context.Context) error {
	if tx, ok := ctx.Value("tx").(pgx.Tx); ok {
		return tx.Rollback(ctx)
	}
	return nil
}

func (p *pg) Commit(ctx context.Context) error {
	if tx, ok := ctx.Value("tx").(pgx.Tx); ok {
		return tx.Commit(ctx)
	}
	return nil
}
