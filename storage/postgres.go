package storage

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Creates a new database connection.
func NewConnection() (*gorm.DB, error) {
	connString := os.Getenv("POSTGRES_CONNSTRING")

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}
