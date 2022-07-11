package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Creates a new database connection.
func NewConnection(connString string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return db, err
	}

	return db, nil
}
