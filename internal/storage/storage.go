package storage

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

// NewConnection to the give Postgres database.
func NewConnection(cfg Config) (*pg.DB, error) {
	opt, err := pg.ParseURL(cfg.URL())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	db := pg.Connect(opt)

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
