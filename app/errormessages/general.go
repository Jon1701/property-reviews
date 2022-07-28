package errormessages

import "fmt"

type ErrorMessage string

const (
	AddressLine1MinLength      int = 3
	AddressLine1MaxLength      int = 1000
	AddressLine2MinLength      int = 3
	AddressLine2MaxLength      int = 1000
	AddressCityMinLength       int = 1
	AddressCityMaxLength       int = 1000
	AddressStateMinLength      int = 1
	AddressStateMaxLength      int = 1000
	AddressPostalCodeMinLength int = 1
	AddressPostalCodeMaxLength int = 20
	AddressCountryMinLength    int = 2
	AddressCountryMaxLength    int = 100
	WebsiteMinLength           int = 3
	WebsiteMaxLength           int = 255
)

const (
	FailedToParseRequestBody ErrorMessage = "failed to parse request body"
	FieldValueRequired       ErrorMessage = "field value is required"
)

var AddressLine1InvalidLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(AddressLine1MinLength, AddressLine1MaxLength)
var AddressLine2InvalidLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(AddressLine2MinLength, AddressLine2MaxLength)
var AddressCityInvalidLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(AddressCityMinLength, AddressCityMaxLength)
var AddressStateInvalidLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(AddressStateMinLength, AddressStateMaxLength)
var AddressPostalCodeInvalidLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(AddressPostalCodeMinLength, AddressPostalCodeMaxLength)
var AddressCountryInvalidLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(AddressCountryMinLength, AddressCountryMaxLength)
var WebsiteInvalidLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(WebsiteMinLength, WebsiteMaxLength)
var InvalidURL ErrorMessage = "invalid URL"
var FieldValueNotAllowed ErrorMessage = "field value not allowed"

// Generates an error message indicating that a field must be between a certain
// length of characters.
func FieldMustBeBetweenXandYCharactersInLength(minLength int, maxLength int) ErrorMessage {
	msg := fmt.Sprintf("field must be between %d and %d characters in length", minLength, maxLength)
	return ErrorMessage(msg)
}

// Generates an error message indicating that a field must have one value in an
// array of values.
func FieldMustBeOneOf(values string) ErrorMessage {
	msg := fmt.Sprintf("field must be one of: %s", values)
	return ErrorMessage(msg)
}
