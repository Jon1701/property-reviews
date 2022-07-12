package validation

import (
	"net/url"

	"github.com/Jon1701/property-reviews/app/errormessages"
	"github.com/Jon1701/property-reviews/app/serializers"
)

// Checks if a given URL is valid.
func IsValidURL(s string) bool {
	_, err := url.ParseRequestURI(s)

	return err == nil
}

// Check if a given URL has the correct length.
func IsURLOfCorrectLength(url string) bool {
	return len(url) >= errormessages.WebsiteMinLength && len(url) <= errormessages.WebsiteMaxLength
}

// Validates Address struct values.
func ValidateAddress(address *serializers.Address, line2Required bool) *Address {
	results := Address{}
	passValidation := true

	// Check if provided address struct is nil.
	if address == nil {
		msg := errormessages.FieldValueRequired

		results.Line1 = &msg
		results.City = &msg
		results.State = &msg
		results.PostalCode = &msg
		results.Country = &msg

		if line2Required {
			results.Line2 = &msg
		}

		return &results
	}

	if address.Line1 == nil {
		// Check if Line 1 is provided.
		msg := errormessages.FieldValueRequired
		results.Line1 = &msg
		passValidation = false
	} else {
		// Check Line 1 length.
		isValidLength := len(*address.Line1) >= errormessages.AddressLine1MinLength && len(*address.Line1) <= errormessages.AddressLine1MaxLength
		if !isValidLength {
			results.Line1 = &errormessages.AddressLine1InvalidLength
			passValidation = false
		}
	}

	// Check if Line 2 is provided if it is required.
	if line2Required {
		msg := errormessages.FieldValueRequired
		results.Line2 = &msg
		passValidation = false
	}

	// If Line 2 is provided, check it's length.
	isLine2ValidLength := address.Line2 != nil && len(*address.Line1) >= errormessages.AddressLine1MinLength && len(*address.Line1) <= errormessages.AddressLine1MaxLength
	if address.Line2 != nil && !isLine2ValidLength {
		results.Line2 = &errormessages.AddressLine2InvalidLength
		passValidation = false
	}

	if address.City == nil {
		// Check if City is provided.
		msg := errormessages.FieldValueRequired
		results.City = &msg
		passValidation = false
	} else {
		// Check City length.
		isValidLength := len(*address.City) >= errormessages.AddressCityMinLength && len(*address.City) <= errormessages.AddressCityMaxLength
		if !isValidLength {
			results.City = &errormessages.AddressCityInvalidLength
			passValidation = false
		}
	}

	if address.State == nil {
		// Check if State is provided.
		msg := errormessages.FieldValueRequired
		results.State = &msg
		passValidation = false
	} else {
		// Check State length.
		isValidLength := len(*address.State) >= errormessages.AddressStateMinLength && len(*address.State) <= errormessages.AddressStateMaxLength
		if !isValidLength {
			results.State = &errormessages.AddressStateInvalidLength
			passValidation = false
		}
	}

	if address.PostalCode == nil {
		// Check if Postal Code is provided.
		msg := errormessages.FieldValueRequired
		results.PostalCode = &msg
		passValidation = false
	} else {
		// Check Postal Code length.
		isValidLength := len(*address.PostalCode) >= errormessages.AddressPostalCodeMinLength && len(*address.PostalCode) <= errormessages.AddressPostalCodeMaxLength
		if !isValidLength {
			results.PostalCode = &errormessages.AddressPostalCodeInvalidLength
			passValidation = false
		}
	}

	if address.Country == nil {
		// Check if Country is provided.
		msg := errormessages.FieldValueRequired
		results.Country = &msg
		passValidation = false
	} else {
		// Check Country length.
		isValidLength := len(*address.Country) >= errormessages.AddressCountryMinLength && len(*address.Country) <= errormessages.AddressCountryMaxLength
		if !isValidLength {
			results.Country = &errormessages.AddressCountryInvalidLength
			passValidation = false
		}
	}

	if passValidation {
		return nil
	}

	return &results
}

// Validates Address struct values, but ignore nil values.
func ValidateAddressIgnoreNil(address *serializers.Address) *Address {
	results := Address{}
	passValidation := true

	// Check Line 1 length.
	if address.Line1 != nil {
		isValidLength := len(*address.Line1) >= errormessages.AddressLine1MinLength && len(*address.Line1) <= errormessages.AddressLine1MaxLength
		if !isValidLength {
			results.Line1 = &errormessages.AddressLine1InvalidLength
			passValidation = false
		}
	}

	// Check Line 2 length.
	if address.Line2 != nil {
		isValidLength := len(*address.Line2) >= errormessages.AddressLine2MinLength && len(*address.Line2) <= errormessages.AddressLine2MaxLength
		if !isValidLength {
			results.Line2 = &errormessages.AddressLine2InvalidLength
		}
	}

	// Check City length.
	if address.City != nil {
		isValidLength := len(*address.City) >= errormessages.AddressCityMinLength && len(*address.City) <= errormessages.AddressCityMaxLength
		if !isValidLength {
			results.City = &errormessages.AddressCityInvalidLength
			passValidation = false
		}
	}

	// Check State length.
	if address.State != nil {
		isValidLength := len(*address.State) >= errormessages.AddressStateMinLength && len(*address.State) <= errormessages.AddressStateMaxLength
		if !isValidLength {
			results.State = &errormessages.AddressStateInvalidLength
			passValidation = false
		}
	}

	// Check Postal Code length.
	if address.PostalCode != nil {
		isValidLength := len(*address.PostalCode) >= errormessages.AddressPostalCodeMinLength && len(*address.PostalCode) <= errormessages.AddressPostalCodeMaxLength
		if !isValidLength {
			results.PostalCode = &errormessages.AddressPostalCodeInvalidLength
			passValidation = false
		}
	}

	// Check Country length.
	if address.Country != nil {
		isValidLength := len(*address.Country) >= errormessages.AddressCountryMinLength && len(*address.Country) <= errormessages.AddressCountryMaxLength
		if !isValidLength {
			results.Country = &errormessages.AddressCountryInvalidLength
			passValidation = false
		}
	}

	if passValidation {
		return nil
	}

	return &results
}
