package errormessages

import "fmt"

const (
	PaginationLimitMinValue = 1
	PaginationLimitMaxValue = 20
)

var InvalidPaginationLimit ErrorMessage = ErrorMessage(fmt.Sprintf("pagination limit must be an integer between %d and %d", PaginationLimitMinValue, PaginationLimitMaxValue))
