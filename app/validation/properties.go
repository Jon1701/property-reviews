package validation

import (
	"github.com/Jon1701/property-reviews/app/errormessages"
	"github.com/Jon1701/property-reviews/app/serializers"
)

type Property struct {
	Address      *Address                    `json:"address,omitempty"`
	Neighborhood *errormessages.ErrorMessage `json:"neighborhood,omitempty"`
	PropertyType *errormessages.ErrorMessage `json:"propertyType,omitempty"`
	BuildingType *errormessages.ErrorMessage `json:"buildingType,omitempty"`
}

// Performs field validation for Property struct values under the
// Create Property route.
func ValidateCreateProperty(property serializers.Property) *Property {
	results := Property{}
	passValidation := true

	// Check Address validity.
	resultsAddress := ValidateAddress(property.Address, false)
	if resultsAddress != nil {
		results.Address = resultsAddress
		passValidation = false
	}

	// Check Neighborhood validity.
	if property.Neighborhood != nil {
		minLength := errormessages.PropertyNeighborhoodMinLength
		maxLength := errormessages.PropertyNeighborhoodMaxLength
		isValidLength := len(*property.Neighborhood) >= minLength && len(*property.Neighborhood) <= maxLength
		if !isValidLength {
			results.Neighborhood = &errormessages.PropertyNeighborhoodInvalidFieldLength
			passValidation = false
		}
	}

	// Check Property Type.
	if property.PropertyType == nil {
		results.PropertyType = &errormessages.PropertyInvalidPropertyType
		passValidation = false
	} else {
		if !isInArrayOfPropertyTypes(*property.PropertyType) {
			results.PropertyType = &errormessages.PropertyInvalidPropertyType
			passValidation = false
		}
	}

	// Check Building Type.
	if property.BuildingType == nil {
		results.BuildingType = &errormessages.PropertyInvalidBuildingType
		passValidation = false
	} else {
		if !isInArrayOfBuildingTypes(*property.BuildingType) {
			results.BuildingType = &errormessages.PropertyInvalidBuildingType
			passValidation = false
		}
	}

	if passValidation {
		return nil
	}

	return &results
}

// Performs field validation for Property struct values under the
// Update Property route.
func ValidateUpdateProperty(property serializers.Property) *Property {
	results := Property{}
	passValidation := true

	// Check Address validity.
	if property.Address != nil {
		resultsAddress := ValidateAddressIgnoreNil(property.Address)
		if resultsAddress != nil {
			results.Address = resultsAddress
			passValidation = false
		}
	}

	// Check Neighborhood validity.
	if property.Neighborhood != nil {
		minLength := errormessages.PropertyNeighborhoodMinLength
		maxLength := errormessages.PropertyNeighborhoodMaxLength
		isValidLength := len(*property.Neighborhood) >= minLength && len(*property.Neighborhood) <= maxLength
		if !isValidLength {
			results.Neighborhood = &errormessages.PropertyNeighborhoodInvalidFieldLength
			passValidation = false
		}
	}

	// Check Property Type.
	if property.PropertyType != nil {
		if !isInArrayOfPropertyTypes(*property.PropertyType) {
			results.PropertyType = &errormessages.PropertyInvalidPropertyType
			passValidation = false
		}
	}

	// Check Building Type.
	if property.BuildingType != nil {
		if !isInArrayOfBuildingTypes(*property.BuildingType) {
			results.BuildingType = &errormessages.PropertyInvalidBuildingType
			passValidation = false
		}
	}

	if passValidation {
		return nil
	}

	return &results
}

// Checks if a given Property Type is in the array of allowed Property Types.
func isInArrayOfPropertyTypes(v serializers.PropertyType) bool {
	isInArray := false

	propertyType := string(v)
	for _, v := range errormessages.ArrayPropertyTypes {
		if propertyType == v {
			isInArray = true
		}
	}

	return isInArray
}

// Checks if a given Building Type is in the array of allowed Building Types.
func isInArrayOfBuildingTypes(v serializers.BuildingType) bool {
	isInArray := false

	buildingType := string(v)
	for _, v := range errormessages.ArrayBuildingTypes {
		if buildingType == v {
			isInArray = true
		}
	}

	return isInArray
}
