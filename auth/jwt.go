package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Payload struct {
	// Email Address.
	EmailAddress string

	// User ID.
	UserID string
}

// Generate a JWT using a given Payload of Claims.
func GenerateJWT(p Payload) (*string, error) {
	mapClaims := jwt.MapClaims{
		"emailAddress": p.EmailAddress,
		"userID":       p.UserID,
		"exp":          time.Now().UTC().Add(time.Hour * 2).Unix(),
		"iss":          "<Not Configured>",
		"iat":          time.Now().UTC().Unix(),
	}

	jwtSigningKey := []byte(os.Getenv("JWT_SIGNING_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	tokenString, err := token.SignedString(jwtSigningKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
