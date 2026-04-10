// Package db contains database connection logic.

package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB establishes a PostgreSQL connection using GORM.
func ConnectDB(dbURL string) (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}
	return db, nil
}
