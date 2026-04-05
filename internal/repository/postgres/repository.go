package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/iskendernarynbaev-lab/exchange-rate-grpc/internal/model"
)

type Repository struct {
	db *sql.DB
}

func New(ctx context.Context, databaseURL string) (*Repository, error) {
	poolCfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse database URL: %w", err)
	}

	db, err := sql.Open("pgx", poolCfg.ConnString())
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	ctxPing, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctxPing); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &Repository{db: db}, nil
}

func (r *Repository) StoreRate(ctx context.Context, rate model.Rate) error {
	const query = `INSERT INTO rates (ask, bid, created_at) VALUES ($1, $2, $3)`
	if _, err := r.db.ExecContext(ctx, query, rate.Ask, rate.Bid, rate.CreatedAt.UTC()); err != nil {
		return fmt.Errorf("insert rate: %w", err)
	}
	return nil
}

func (r *Repository) Close() error {
	if r == nil || r.db == nil {
		return nil
	}
	return r.db.Close()
}
