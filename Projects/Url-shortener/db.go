package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func initDB() (*pgxpool.Pool, error) {

	dbUrl := os.Getenv("DATABASE_URL")
	ctx := context.Background()

	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}

	_, err = pool.Exec(ctx, `CREATE TABLE IF NOT EXISTS urls (short_code TEXT PRIMARY KEY, original_url TEXT NOT NULL)`)

	if err != nil {
		return nil, err
	}

	return pool, nil
}
