package database

import (
	"database/sql"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

// Migrator declare method to run migrations
type Migrator interface {
	Up(db *sql.DB, dir string) error
}

// Migration is an empty structure to implement IMigrator method
type Migration struct{}

// Up implements IMigrator method
func (m *Migration) Up(db *sql.DB, dir string) error {
	return goose.Up(db, dir)
}

func runMigrations(pool *pgxpool.Pool, migrator Migrator) error {
	db := stdlib.OpenDBFromPool(pool)
	return migrator.Up(db, "migrations")
}
