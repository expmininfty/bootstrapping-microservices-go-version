package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func New(dsn string) (*DB, error) {

	connConf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	connConf.MaxConnIdleTime = 10 * time.Minute
	connConf.MaxConnLifetime = 30 * time.Second
	connConf.MaxConns = 50
	connConf.MinConns = 1

	pool, err := pgxpool.NewWithConfig(context.Background(), connConf)

	if err != nil {
		return nil, err
	}

	return &DB{pool}, nil
}
