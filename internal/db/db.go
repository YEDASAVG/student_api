package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), dbURL)

	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Unable to ping to database")
	}
	return pool, nil
}
