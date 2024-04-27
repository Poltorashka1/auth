package pg

//
// import (
//	"auth/internal/client/db"
//	"auth/internal/logger"
//	"context"
//	"fmt"
//	"github.com/jackc/pgx/v4/pgxpool"
//)
//
//type pgClient struct {
//	db  db.DB
//	log logger.Logger
//}
//
//func New(ctx context.Context, dsn string, log logger.Logger) (db.Client, error) { // db.Client
//	pool, err := pgxpool.Connect(ctx, dsn)
//	if err != nil {
//		return nil, fmt.Errorf("failed to connect to db: %w", err)
//	}
//
//	return &pgClient{db: NewDB(pool, log)}, nil
//}
//
//func (c *pgClient) DB() db.DB {
//	return c.db
//}
//
//func (c *pgClient) Close() error {
//	if c.db != nil {
//		c.db.Close()
//	}
//	return nil
//}
