package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(url string, maxConns int) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = int32(maxConns)

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
