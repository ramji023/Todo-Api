package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	var ctx context.Context = context.Background()

	var config *pgxpool.Config
	var err error

	config, err = pgxpool.ParseConfig(databaseURL)

	if err != nil {
		fmt.Println("Unable to parse database url")
	}

	var pool *pgxpool.Pool

	pool, err = pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		fmt.Println("Unable to create connection pool")
		return nil, err
	}

	err = pool.Ping(ctx)

	if err != nil {
		fmt.Println("Unable to ping database")
		pool.Close()
		return nil, err
	}

	fmt.Println("Successfully connected to PostgreSQL database")

	return pool, nil
}
