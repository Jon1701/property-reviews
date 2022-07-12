package errormessages

const (
	PasswordMinLength     int = 8
	PasswordMaxLength     int = 1000
	EmailAddressMinLength int = 5
	EmailAddressMaxLength int = 255
)

var EmailAddressAlreadyRegistered ErrorMessage = "email address is already registered"
var EmailAddressInvalidFieldLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(EmailAddressMinLength, EmailAddressMaxLength)
var EmailAddressOrPasswordIsIncorrect ErrorMessage = "email address or password is incorrect"
var EmailAddressRequired ErrorMessage = "a valid email address is required"
var PasswordInvalidFieldLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(PasswordMinLength, PasswordMaxLength)
