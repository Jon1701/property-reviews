package controllers

import "gorm.io/gorm"

type AppContext struct {
	DB *gorm.DB
}

func New(db *gorm.DB) AppContext {
	return AppContext{db}
}
