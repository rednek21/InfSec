package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUserTable, downUserTable)
}

func upUserTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
    		username VARCHAR(64),
    		email VARCHAR(64),
    		role VARCHAR(16),
    		password VARCHAR(255)
		);
	`)
	return err
}

func downUserTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `DROP TABLE users`)
	return err
}
