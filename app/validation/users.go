package validation

import (
	"regexp"

	"github.com/Jon1701/property-reviews/app/errors"
)

type User struct {
	ID *string `json:"id,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
}

type UserValidationResults struct {
	Username *errors.ErrorMessage `json:"username,omitempty"`
	Password *errors.ErrorMessage `json:"password,omitempty"`
	EmailAddress *errors.ErrorMessage `json:"emailAddress,omitempty"`
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Performs field validation for the Create User route.
func ValidateCreateUser(user User) (*UserValidationResults) {
	results := UserValidationResults{}
	passValidation := true

	if user.Username == nil {
		// No username.
		msg := errors.UsernameInvalidFieldLength
		results.Username = &msg
		passValidation = false
	} else {
		// Length check.
		if len(*user.Username) < errors.UsernameMinLength || len(*user.Username) > errors.UsernameMaxLength {
			msg := errors.UsernameInvalidFieldLength
			results.Username = &msg
			passValidation = false
		}
	}

	if user.Password == nil {
		// No password.
		msg := errors.PasswordInvalidFieldLength
		results.Password = &msg
		passValidation = false
	} else {
		// Length check.
		if len(*user.Password) < errors.PasswordMinLength || len(*user.Password) > errors.PasswordMaxLength {
			msg := errors.PasswordInvalidFieldLength
			results.Password = &msg
			passValidation = false
		}
	}

	if user.EmailAddress == nil {
		// No email address.
		msg := errors.EmailAddressRequired
		results.EmailAddress = &msg
		passValidation = false
	} else {
		// Length check.
		if len(*user.EmailAddress) <= errors.EmailAddressMinLength || len(*user.EmailAddress) > errors.EmailAddressMaxLength {
			msg := errors.EmailAddressInvalidFieldLength
			results.EmailAddress = &msg
			passValidation = false
		} else {
			// Valid email check.
			if !emailRegex.MatchString(*user.EmailAddress) {
				msg := errors.EmailAddressRequired
				results.EmailAddress = &msg
				passValidation = false
			}
		}
	}

	if passValidation {
		return nil
	}

	return &results
}