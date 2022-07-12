package errormessages

const (
	ManagementCompanyNameMinLength int = 3
	ManagementCompanyNameMaxLength int = 1000
)

var ManagementCompanyNameInvalidFieldLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(ManagementCompanyNameMinLength, ManagementCompanyNameMaxLength)
