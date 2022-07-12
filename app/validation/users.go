package validation

import (
	"regexp"

	"github.com/Jon1701/property-reviews/app/errormessages"
	"github.com/Jon1701/property-reviews/app/serializers"
)

type User struct {
	Username     *errormessages.ErrorMessage `json:"username,omitempty"`
	Password     *errormessages.ErrorMessage `json:"password,omitempty"`
	EmailAddress *errormessages.ErrorMessage `json:"emailAddress,omitempty"`
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Performs field validation for the Create User route.
func ValidateCreateUser(user serializers.User) *User {
	results := User{}
	passValidation := true

	if user.Password == nil {
		// No password.
		msg := errormessages.PasswordInvalidFieldLength
		results.Password = &msg
		passValidation = false
	} else {
		// Length check.
		if len(*user.Password) < errormessages.PasswordMinLength || len(*user.Password) > errormessages.PasswordMaxLength {
			msg := errormessages.PasswordInvalidFieldLength
			results.Password = &msg
			passValidation = false
		}
	}

	if user.EmailAddress == nil {
		// No email address.
		msg := errormessages.EmailAddressRequired
		results.EmailAddress = &msg
		passValidation = false
	} else {
		// Length check.
		if len(*user.EmailAddress) <= errormessages.EmailAddressMinLength || len(*user.EmailAddress) > errormessages.EmailAddressMaxLength {
			msg := errormessages.EmailAddressInvalidFieldLength
			results.EmailAddress = &msg
			passValidation = false
		} else {
			// Valid email check.
			if !emailRegex.MatchString(*user.EmailAddress) {
				msg := errormessages.EmailAddressRequired
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

// Performs field validation for the User Login route.
func ValidateUserLogin(user serializers.User) *User {
	results := User{}
	passValidation := true

	// Check if Password is provided.
	if user.Password == nil || (user.Password != nil && len(*user.Password) == 0) {
		// No password.
		msg := errormessages.FieldValueRequired
		results.Password = &msg
		passValidation = false
	}

	// Check if Email Address is provided.
	if user.EmailAddress == nil || (user.EmailAddress != nil && len(*user.EmailAddress) == 0) {
		// No email address.
		msg := errormessages.FieldValueRequired
		results.EmailAddress = &msg
		passValidation = false
	}

	if passValidation {
		return nil
	}

	return &results
}
