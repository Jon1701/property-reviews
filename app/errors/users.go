package errors

const (
	UsernameMinLength     int = 1
	UsernameMaxLength     int = 50
	PasswordMinLength     int = 8
	PasswordMaxLength     int = 1000
	EmailAddressMinLength int = 5
	EmailAddressMaxLength int = 255
)

var EmailAddressAlreadyRegistered ErrorMessage = "email address is already registered"
var EmailAddressInvalidFieldLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(EmailAddressMinLength, EmailAddressMaxLength)
var EmailAddressRequired ErrorMessage = "a valid email address is required"
var PasswordInvalidFieldLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(PasswordMinLength, PasswordMaxLength)
var UsernameInvalidFieldLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(UsernameMinLength, UsernameMaxLength)
