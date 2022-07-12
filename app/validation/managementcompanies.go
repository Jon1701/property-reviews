package validation

import (
	"github.com/Jon1701/property-reviews/app/errormessages"
	"github.com/Jon1701/property-reviews/app/serializers"
)

type Address struct {
	Line1      *errormessages.ErrorMessage `json:"line1,omitempty"`
	Line2      *errormessages.ErrorMessage `json:"line2,omitempty"`
	City       *errormessages.ErrorMessage `json:"city,omitempty"`
	State      *errormessages.ErrorMessage `json:"state,omitempty"`
	PostalCode *errormessages.ErrorMessage `json:"postalCode,omitempty"`
	Country    *errormessages.ErrorMessage `json:"country,omitempty"`
}

type ManagementCompany struct {
	Name    *errormessages.ErrorMessage `json:"name,omitempty"`
	Address *Address                    `json:"address,omitempty"`
	Website *errormessages.ErrorMessage `json:"website,omitempty"`
}

// Performs field validation for Management Company struct values under the
// Create Management Company route.
func ValidateCreateManagementCompany(company serializers.ManagementCompany) *ManagementCompany {
	results := ManagementCompany{}
	passValidation := true

	// Check Name length.
	if company.Name == nil {
		msg := errormessages.FieldValueRequired
		results.Name = &msg
		passValidation = false
	} else {
		isLengthValid := len(*company.Name) >= errormessages.ManagementCompanyNameMinLength || len(*company.Name) <= errormessages.ManagementCompanyNameMinLength
		if !isLengthValid {
			results.Name = &errormessages.ManagementCompanyNameInvalidFieldLength
			passValidation = false
		}
	}

	// Check Address validity.
	resultsAddress := ValidateAddress(company.Address, false)
	if resultsAddress != nil {
		results.Address = resultsAddress
		passValidation = false
	}

	// Check Website URL validity and length.
	if company.Website != nil {
		if !IsValidURL(*company.Website) {
			results.Website = &errormessages.InvalidURL
			passValidation = false
		}

		if !IsURLOfCorrectLength(*company.Website) {
			results.Website = &errormessages.WebsiteInvalidLength
			passValidation = false
		}
	}

	if passValidation {
		return nil
	}

	return &results
}

// Performs field validation for Management Company struct values under the
// Update Management Company route.
func ValidateUpdateManagementCompany(company serializers.ManagementCompany) *ManagementCompany {
	results := ManagementCompany{}
	passValidation := true

	// Check Company Name length.
	if company.Name != nil {
		isValidLength := len(*company.Name) >= errormessages.ManagementCompanyNameMinLength && len(*company.Name) <= errormessages.ManagementCompanyNameMaxLength
		if !isValidLength {
			results.Name = &errormessages.ManagementCompanyNameInvalidFieldLength
			passValidation = false
		}
	}

	// Check Address length.
	if company.Address != nil {
		resultsAddress := ValidateAddressIgnoreNil(company.Address)
		if resultsAddress != nil {
			results.Address = resultsAddress
			passValidation = false
		}
	}

	// Check Website URL validity and length.
	if company.Website != nil {
		if !IsValidURL(*company.Website) {
			results.Website = &errormessages.InvalidURL
			passValidation = false
		}

		if !IsURLOfCorrectLength(*company.Website) {
			results.Website = &errormessages.WebsiteInvalidLength
			passValidation = false
		}
	}

	if passValidation {
		return nil
	}

	return &results
}
