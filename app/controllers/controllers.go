package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppContext struct {
	DB *gorm.DB
}

func New(db *gorm.DB) AppContext {
	return AppContext{db}
}

// Serializes the request body into given struct.
func SerializeRequestBody(c *gin.Context, v interface{}) error {
	// Read request body.
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	// Parse request body.
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}
