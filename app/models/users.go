package models

type User struct {
	ID int `gorm:"primaryKey"`

	// UUID for this User.
	IDHash *string `gorm:"column:id_hash"`

	// Username.
	Username *string `gorm:"column:username"`

	// Email Address.
	EmailAddress *string `gorm:"column:email_address"`

	// Hashed password.
	Password *string `gorm:"column:password"`
}
