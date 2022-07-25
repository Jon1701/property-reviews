package models

import "time"

type User struct {
	ID int `gorm:"primaryKey"`

	// UUID for this User.
	IDHash *string `gorm:"column:id_hash"`

	// Email Address.
	EmailAddress *string `gorm:"column:email_address"`

	// Hashed password.
	Password *string `gorm:"column:password"`

	// User creation time.
	CreatedAt time.Time `gorm:"column:created_at"`

	// User update time.
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
