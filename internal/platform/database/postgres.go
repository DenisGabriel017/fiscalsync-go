package database

import (
	"context"
	"fmt"

	"github.com/DenisGabriel017/fiscalsync-go/internal/platform/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	pool, err := pgxpool.New(ctx, connectionString)

	if err != nil {
		return nil, fmt.Errorf("criar pool do PostgreSQL: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("verificar conexão com PostgreSQL: %w", err)
	}
	return pool, nil
}
