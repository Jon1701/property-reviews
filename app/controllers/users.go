package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Jon1701/property-reviews/app/errormessages"
	"github.com/Jon1701/property-reviews/app/models"
	"github.com/Jon1701/property-reviews/app/serializers"
	"github.com/Jon1701/property-reviews/app/validation"
	"github.com/Jon1701/property-reviews/auth"
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

	// Serialize the request body, respond with Bad Request if unable to.
	err := SerializeRequestBody(c, &user)
	if err != nil {
		msg := errormessages.FailedToParseRequestBody
		c.JSON(http.StatusBadRequest, gin.H{
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

		c.Data(http.StatusBadRequest, "application/json", []byte(body))
		return
	}

	// Check if Email Address is already registered.
	lowercaseEmail := strings.ToLower(*user.EmailAddress)
	m := models.User{}
	appCtx.DB.Where("email_address = ?", lowercaseEmail).First(&m)
	if m.IDHash != nil {
		msg := errormessages.EmailAddressAlreadyRegistered

		v := validation.User{EmailAddress: &msg}
		body, err := json.Marshal(v)
		if err != nil {
			panic(fmt.Sprintf("Failed to Marshal the User Validation struct: %v\n", err))
		}

		c.Data(http.StatusBadRequest, "application/json", []byte(body))
		return
	}

	// Hash and salt the password.
	pw := hashAndSalt(*user.Password)

	// Generate ID hash.
	idHash := generateUserIDHash()

	// Persist into database.
	m = models.User{
		IDHash:       &idHash,
		EmailAddress: &lowercaseEmail,
		Password:     &pw,
	}
	result := appCtx.DB.Create(&m)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to persist User in the database: %v\n", result.Error))
	}

	c.Header("Location", fmt.Sprintf("/api/users/%s", *m.IDHash))
	c.JSON(http.StatusNoContent, nil)
}

// Logs in a User.
func (appCtx *AppContext) UserLogin(c *gin.Context) {
	user := serializers.User{}

	// Serialize the request body, respond with Bad Request if unable to.
	err := SerializeRequestBody(c, &user)
	if err != nil {
		msg := errormessages.FailedToParseRequestBody
		c.JSON(http.StatusBadRequest, gin.H{
			"message": &msg,
		})
		return
	}

	// Validation request body object.
	results := validation.ValidateUserLogin(user)
	if results != nil {
		body, err := json.Marshal(results)
		if err != nil {
			panic(error.Error(err))
		}

		c.Data(http.StatusBadRequest, "application/json", []byte(body))
		return
	}

	// Check if Email Address is registered.
	lowercaseEmail := strings.ToLower(*user.EmailAddress)
	m := models.User{}
	result := appCtx.DB.Where("email_address = ?", lowercaseEmail).First(&m)
	if result.Error != nil {
		body, err := json.Marshal(validation.User{
			EmailAddress: &errormessages.EmailAddressOrPasswordIsIncorrect,
			Password:     &errormessages.EmailAddressOrPasswordIsIncorrect,
		})

		if err != nil {
			panic(fmt.Sprintf("Failed to Marshal User Login validation struct (Email not found): %+v\n", err))
		}

		c.Data(http.StatusBadRequest, "application/json", []byte(body))
		return
	}

	// Check if Password matches Hashed Password.
	isMatching := compareHashAndPassword(*m.Password, *user.Password)
	if !isMatching {
		body, err := json.Marshal(validation.User{
			EmailAddress: &errormessages.EmailAddressOrPasswordIsIncorrect,
			Password:     &errormessages.EmailAddressOrPasswordIsIncorrect,
		})

		if err != nil {
			panic(fmt.Sprintf("Failed to Marshal User Login validation struct (Password mismatch): %+v\n", err))
		}

		c.Data(http.StatusBadRequest, "application/json", []byte(body))
		return
	}

	// Generate JWT.
	lowercaseEmail = strings.ToLower(*m.EmailAddress)
	p := auth.Payload{EmailAddress: lowercaseEmail, UserID: *m.IDHash}
	token, err := auth.GenerateJWT(p)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate JWT: %+v\n", err))
	}

	c.Data(http.StatusOK, "application/json", []byte(*token))
}
