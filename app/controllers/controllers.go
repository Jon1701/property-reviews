package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Jon1701/property-reviews/app/errormessages"
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
// If serialization fails, send a response to the client with Status
// Bad Request.
func SerializeRequestBodyAndRespondIfErrored(c *gin.Context, v interface{}) {
	// Read request body.
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		msg := errormessages.FailedToParseRequestBody
		c.JSON(http.StatusBadRequest, gin.H{
			"message": &msg,
		})
		return
	}

	// Parse request body.
	err = json.Unmarshal(data, v)
	if err != nil {
		msg := errormessages.FailedToParseRequestBody
		c.JSON(http.StatusBadRequest, gin.H{
			"message": &msg,
		})
		return
	}
}
