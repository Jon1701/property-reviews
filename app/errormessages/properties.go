package errormessages

import (
	"fmt"
	"strings"

	"github.com/Jon1701/property-reviews/app/serializers"
)

const (
	PropertyNeighborhoodMinLength int = 3
	PropertyNeighborhoodMaxLength int = 255
)

// Array of Property/Building Types.
var ArrayPropertyTypes []string
var ArrayBuildingTypes []string

// String representations of the array of Property/Building Types.
var strArrayPropertyTypes string
var strArrayBuildingTypes string

// Error messages.
var PropertyNeighborhoodInvalidFieldLength ErrorMessage = FieldMustBeBetweenXandYCharactersInLength(PropertyNeighborhoodMinLength, PropertyNeighborhoodMaxLength)
var PropertyInvalidPropertyType ErrorMessage
var PropertyInvalidBuildingType ErrorMessage
var PropertyNotFound ErrorMessage = "property not found"
var FieldValueNotAllowedForCurrentBuildingType ErrorMessage = "field value not allowed for current building type"

func init() {
	// Build arrays of Building/Property Types.
	ArrayBuildingTypes = generateArrayBuildingTypes()
	ArrayPropertyTypes = generateArrayPropertyTypes()

	// Build string representations of the arrays of Building/Property Types.
	strArrayBuildingTypes = generateFormattedArrayString(serializers.TypesOfBuildings)
	strArrayPropertyTypes = generateFormattedArrayString(serializers.TypesOfProperties)

	// Set error messages.
	PropertyInvalidPropertyType = FieldMustBeOneOf(strArrayPropertyTypes)
	PropertyInvalidBuildingType = FieldMustBeOneOf(strArrayBuildingTypes)
}

// Generates an array of Property Types.
func generateArrayPropertyTypes() []string {
	var array []string

	for _, val := range serializers.TypesOfProperties {
		array = append(array, string(val))
	}

	return array
}

// Generates an array of Building Types.
func generateArrayBuildingTypes() []string {
	var array []string

	for _, val := range serializers.TypesOfBuildings {
		array = append(array, string(val))
	}

	return array
}

// Generates a string representation of an array of strings.
func generateFormattedArrayString(v interface{}) string {
	var sb strings.Builder

	sb.WriteString("[")

	switch v := v.(type) {
	case map[string]serializers.PropertyType:
		for _, val := range v {
			token := fmt.Sprintf("'%s', ", string(val))
			sb.WriteString(token)
		}

	case map[string]serializers.BuildingType:
		for _, val := range v {
			token := fmt.Sprintf("'%s', ", string(val))
			sb.WriteString(token)
		}

	default:
		panic("Unknown type")
	}

	// Get current string.
	currentStr := sb.String()

	// Remove last comma and space.
	currentStr = currentStr[0 : len(currentStr)-2]

	// Add closing array bracket, return.
	return fmt.Sprintf("%s]", currentStr)
}
