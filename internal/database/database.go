package database

import (
	"InfSec/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Database struct {
	Pool *pgxpool.Pool
}

var (
	dbInstance *Database
	dbOnce     sync.Once
	dbErr      error
)

func New(ctx context.Context, cfg *config.Config) (*Database, error) {
	dbOnce.Do(func() {
		dbInstance, dbErr = initializeDatabase(ctx, cfg)
		if dbErr != nil {
			return
		}

		if err := runMigrations(dbInstance.Pool, &Migration{}); err != nil {
			dbErr = fmt.Errorf("failed to run migrations: %w", err)
			return
		}
	})

	if dbErr != nil {
		return nil, dbErr
	}

	return dbInstance, nil
}

func initializeDatabase(ctx context.Context, cfg *config.Config) (*Database, error) {
	dbUrl := makeDBUrl(cfg)

	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Database{Pool: pool}, nil
}

func makeDBUrl(cfg *config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database)
}

// Close gracefully closes the database connection pool.
func (d *Database) Close() {
	if d.Pool != nil {
		d.Pool.Close()
	}
}
