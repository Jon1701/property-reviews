package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Jon1701/property-reviews/app/errormessages"
	"github.com/Jon1701/property-reviews/app/models"
	"github.com/Jon1701/property-reviews/app/serializers"
	"github.com/Jon1701/property-reviews/app/validation"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Generates a hyphenless UUID with "property_" prepended.
func generatePropertyIDHash() string {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")

	return fmt.Sprintf("property_%s", id)
}

// Creates a Property.
func (appCtx *AppContext) CreateProperty(c *gin.Context) {
	property := serializers.Property{}

	// Serialize the Request Body.
	err := SerializeRequestBody(c, &property)
	if err != nil {
		msg := errormessages.FailedToParseRequestBody
		c.JSON(http.StatusBadRequest, gin.H{
			"message": &msg,
		})
		return
	}

	// Field validation.
	results := validation.ValidateCreateProperty(property)
	if results != nil {
		body, err := json.Marshal(results)
		if err != nil {
			panic(fmt.Sprintf("Failed to marshal the field validation results struct into JSON: %+v\n", err))
		}

		c.Data(http.StatusBadRequest, "application/json", body)
	}

	// Persist into database.
	idHash := generatePropertyIDHash()
	m := models.Property{
		IDHash:            &idHash,
		AddressLine1:      property.Address.Line1,
		AddressCity:       property.Address.City,
		AddressState:      property.Address.State,
		AddressPostalCode: property.Address.PostalCode,
		AddressCountry:    property.Address.Country,
		PropertyType:      (*string)(property.PropertyType),
		BuildingType:      (*string)(property.BuildingType),
	}
	if property.Address.Line2 != nil {
		m.AddressLine2 = property.Address.Line2
	}
	if property.Neighborhood != nil {
		m.Neighborhood = property.Neighborhood
	}
	result := appCtx.DB.Create(&m)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to persist Property in database: %+v\n", result.Error))
	}

	c.Header("Location", fmt.Sprintf("/api/property/%s", idHash))
	c.JSON(http.StatusNoContent, nil)
}

// Updates a Property.
func (appCtx *AppContext) UpdateProperty(c *gin.Context) {
	propertyID := c.Param("propertyID")

	// Check if Property exists in the database.
	m := models.Property{}
	result := appCtx.DB.Where("id_hash = ?", propertyID).First(&m)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, &gin.H{
			"message": &errormessages.PropertyNotFound,
		})
		return
	}

	// Serialize response body.
	property := serializers.Property{}
	err := SerializeRequestBody(c, &property)
	if err != nil {
		msg := errormessages.FailedToParseRequestBody
		c.JSON(http.StatusBadRequest, gin.H{
			"message": &msg,
		})
		return
	}

	// Validate fields.
	results := validation.ValidateUpdateProperty(property)
	if results != nil {
		body, err := json.Marshal(results)
		if err != nil {
			panic(fmt.Sprintf("Failed to marshal the field validation results struct into JSON: %+v\n", err))
		}
		c.Data(http.StatusBadRequest, "application/json", body)
		return
	}

	// Persist into database.
	m.PropertyType = (*string)(property.PropertyType)
	m.BuildingType = (*string)(property.BuildingType)
	m.Neighborhood = property.Neighborhood
	if property.Address != nil {
		m.AddressLine1 = property.Address.Line1
		m.AddressLine2 = property.Address.Line2
		m.AddressCity = property.Address.City
		m.AddressState = property.Address.State
		m.AddressPostalCode = property.Address.PostalCode
		m.AddressCountry = property.Address.Country
	}
	result = appCtx.DB.Updates(&m)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to update Property in database: %+v\n", result.Error))
	}

	// Get updated Property.
	m = models.Property{}
	result = appCtx.DB.Where("id_hash = ?", propertyID).First(&m)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to get Property from the database: %+v\n", result.Error))
	}

	// Serialize row into JSON.
	property = serializers.Property{
		ID:           &propertyID,
		PropertyType: (*serializers.PropertyType)(m.PropertyType),
		BuildingType: (*serializers.BuildingType)(m.BuildingType),
		Neighborhood: m.Neighborhood,
		Address: &serializers.Address{
			Line1:      m.AddressLine1,
			Line2:      m.AddressLine2,
			City:       m.AddressCity,
			State:      m.AddressState,
			PostalCode: m.AddressPostalCode,
			Country:    m.AddressCountry,
		},
	}
	body, err := json.Marshal(property)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal the database row struct into JSON: %+v\n", err))
	}

	c.Data(http.StatusOK, "application/json", body)
}
