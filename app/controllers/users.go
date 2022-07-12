package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/Jon1701/property-reviews/app/errors"
	"github.com/Jon1701/property-reviews/app/models"
	"github.com/Jon1701/property-reviews/app/serializers"
	"github.com/Jon1701/property-reviews/app/validation"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Generates a hyphenless UUID with "user_" prepended.
func generateUserIDHash() string {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")

	return fmt.Sprintf("user_%s", id)
}

// Generates a hashed and salted password.
func hashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate hash from password - Error %v", err))
	}

	return string(hash)
}

// Compares a hashed password and plaintext password.
func compareHashAndPassword(hashedPassword string, plainPassword string) bool {
	byteHash := []byte(hashedPassword)
	bytePassword := []byte(plainPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)

	return err == nil
}

// Creates a User.
func (appCtx *AppContext) CreateUser(c *gin.Context) {
	user := serializers.User{}

	// Read request body.
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		msg := errors.FailedToParseRequestBody
		c.JSON(400, gin.H{
			"message": &msg,
		})
		return
	}

	// Parse request body.
	err = json.Unmarshal(data, &user)
	if err != nil {
		msg := errors.FailedToParseRequestBody
		c.JSON(400, gin.H{
			"message": &msg,
		})
		return
	}

	// Validation request body object.
	results := validation.ValidateCreateUser(user)
	if results != nil {
		body, err := json.Marshal(results)
		if err != nil {
			panic(error.Error(err))
		}

		c.Data(400, "application/json", []byte(body))
		return
	}

	// Check if Email Address is already registered.
	m := models.User{}
	appCtx.DB.Where("email_address = ?", user.EmailAddress).First(&m)
	if m.IDHash != nil {
		msg := errors.EmailAddressAlreadyRegistered

		v := validation.User{EmailAddress: &msg}
		body, err := json.Marshal(v)
		if err != nil {
			panic(error.Error(err))
		}

		c.Data(400, "application/json", []byte(body))
		return
	}

	// Hash and salt the password.
	pw := hashAndSalt(*user.Password)

	// Generate ID hash.
	idHash := generateUserIDHash()

	// Persist into database.
	m = models.User{
		IDHash:       &idHash,
		Username:     user.Username,
		EmailAddress: user.EmailAddress,
		Password:     &pw,
	}
	result := appCtx.DB.Create(&m)
	if result.Error != nil {
		log.Fatal("Failed to create User", result.Error)
		return
	}

	c.Header("Location", fmt.Sprintf("/api/users/%s", *m.IDHash))
	c.JSON(204, nil)
}
