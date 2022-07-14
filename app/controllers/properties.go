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
