package dbclient

import (
	"context"
	"fmt"
	"time"
	"user-rest-api/internal/config"
	"user-rest-api/pkg/logger"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, cfg config.StorageConfig) *pgxpool.Pool {
	logger := logger.GetLogger()

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	logger.Infof("connecting to %s", dsn)
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		logger.Fatalf("unable to connect to %s: %v\n", dsn, err)
	}

	return pool
}
