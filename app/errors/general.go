package errors

import "fmt"

type ErrorMessage string

const (
	FailedToParseRequestBody ErrorMessage = "failed to parse request body"
	FieldValueRequired ErrorMessage = "field value is required"
)

// Generates an error message indicating that a field must be between a certain
// length of characters.
func FieldMustBeBetweenXandYCharactersInLength(minLength int, maxLength int) ErrorMessage {
	msg := fmt.Sprintf("field must be between %d and %d characters in length", minLength, maxLength)
	return ErrorMessage(msg) 
}