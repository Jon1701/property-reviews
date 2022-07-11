package errors

const (
	UsernameMinLength     int = 1
	UsernameMaxLength     int = 50
	PasswordMinLength     int = 8
	PasswordMaxLength     int = 1000
	EmailAddressMinLength int = 5
	EmailAddressMaxLength int = 255
)

var UsernameInvalidFieldLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(UsernameMinLength, UsernameMaxLength)
var PasswordInvalidFieldLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(PasswordMinLength, PasswordMaxLength)
var EmailAddressRequired ErrorMessage = "a valid email address is required"
var EmailAddressInvalidFieldLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(EmailAddressMinLength, EmailAddressMaxLength)
