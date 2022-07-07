package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Jon1701/property-reviews/app/errors"
	"github.com/Jon1701/property-reviews/app/validation"
	"github.com/gin-gonic/gin"
)

// Creates a User.
func CreateUser(c *gin.Context) {
	user := validation.User{}

	data, err := ioutil.ReadAll(c.Request.Body);
	if err != nil {
		msg := errors.FailedToParseRequestBody;
		c.JSON(400, gin.H{
			"message": &msg,
		})
		return
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		msg := errors.FailedToParseRequestBody;
		c.JSON(400, gin.H{
			"message": &msg,
		})
		return
	}

	results := validation.ValidateCreateUser(user)
	if results != nil {
		body, err := json.Marshal(results)
		if err != nil {
			panic(error.Error(err))
		}

		c.Data(400, "application/json", []byte(body))
		return;
	}
	
	c.JSON(200, gin.H{
		"message": "Create User controller says hello world",
	})
}	