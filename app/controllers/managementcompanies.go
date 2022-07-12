package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Jon1701/property-reviews/app/errormessages"
	"github.com/Jon1701/property-reviews/app/models"
	"github.com/Jon1701/property-reviews/app/serializers"
	"github.com/Jon1701/property-reviews/app/validation"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Generates a hyphenless UUID with "management_" prepended.
func generateManagementIDHash() string {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")

	return fmt.Sprintf("management_%s", id)
}

// Creates a Management Company.
func (appCtx *AppContext) CreateManagementCompany(c *gin.Context) {
	company := serializers.ManagementCompany{}

	// Read request body.
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		msg := errormessages.FailedToParseRequestBody
		c.JSON(http.StatusBadRequest, gin.H{
			"message": &msg,
		})
		return
	}

	// Parse request body.
	err = json.Unmarshal(data, &company)
	if err != nil {
		msg := errormessages.FailedToParseRequestBody
		c.JSON(http.StatusBadRequest, gin.H{
			"message": &msg,
		})
		return
	}

	// Field validation.
	results := validation.ValidateCreateManagementCompany(company)
	if results != nil {
		body, err := json.Marshal(results)
		if err != nil {
			panic(fmt.Sprintf("Failed to marshal the field validation results struct into JSON: %+v\n", err))
		}

		c.Data(http.StatusBadRequest, "application/json", []byte(body))
		return
	}

	// Persist into database.
	idHash := generateManagementIDHash()
	m := models.ManagementCompany{
		IDHash:            &idHash,
		Name:              company.Name,
		AddressLine1:      company.Address.Line1,
		AddressCity:       company.Address.City,
		AddressState:      company.Address.State,
		AddressPostalCode: company.Address.PostalCode,
		AddressCountry:    company.Address.Country,
	}
	if company.Address.Line2 != nil {
		m.AddressLine2 = company.Address.Line2
	}
	if company.Website != nil {
		m.Website = company.Website
	}
	result := appCtx.DB.Create(&m)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to persist Management Company in database: %+v\n", result.Error))
	}

	c.Header("Location", fmt.Sprintf("/api/management/%s", *m.IDHash))
	c.JSON(http.StatusNoContent, nil)
}

// Updates a Management Company.
func (appCtx *AppContext) UpdateManagementCompany(c *gin.Context) {
	company := serializers.ManagementCompany{}

	// Check if Management Company exists.
	managementID := c.Param("managementID")
	m := models.ManagementCompany{IDHash: &managementID}
	result := appCtx.DB.Where("id_hash = ?", managementID).First(&m)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, &gin.H{"message": "Management Company not found"})
		return
	}

	// Read request body.
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		msg := errormessages.FailedToParseRequestBody
		c.JSON(http.StatusBadRequest, gin.H{
			"message": &msg,
		})
		return
	}

	// Parse request body.
	err = json.Unmarshal(data, &company)
	if err != nil {
		msg := errormessages.FailedToParseRequestBody
		c.JSON(http.StatusBadRequest, gin.H{
			"message": &msg,
		})
		return
	}

	// Field validation.
	results := validation.ValidateUpdateManagementCompany(company)
	if results != nil {
		body, err := json.Marshal(results)
		if err != nil {
			panic(fmt.Sprintf("Failed to marshal the field validation results struct into JSON: %+v\n", err))
		}

		c.Data(http.StatusBadRequest, "application/json", []byte(body))
		return
	}

	// Persist into database.
	m.Name = company.Name
	if company.Address != nil {
		m.AddressLine1 = company.Address.Line1
		m.AddressLine2 = company.Address.Line2
		m.AddressCity = company.Address.City
		m.AddressState = company.Address.State
		m.AddressPostalCode = company.Address.PostalCode
		m.AddressCountry = company.Address.Country
	}
	m.Website = company.Website
	result = appCtx.DB.Updates(&m)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to persist Management Company in database: %+v\n", err))
	}

	// Get updated row.
	m = models.ManagementCompany{}
	result = appCtx.DB.Where("id_hash = ?", managementID).First(&m)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to get Management Company from database: %+v\n", err))
	}

	// Serialize row into JSON.
	company = serializers.ManagementCompany{
		ID:   m.IDHash,
		Name: m.Name,
		Address: &serializers.Address{
			Line1:      m.AddressLine1,
			Line2:      m.AddressLine2,
			City:       m.AddressCity,
			State:      m.AddressState,
			PostalCode: m.AddressPostalCode,
			Country:    m.AddressCountry,
		},
		Website: m.Website,
	}
	body, err := json.Marshal(company)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal the database row struct into JSON: %+v\n", err))
	}

	c.Data(http.StatusOK, "application/json", body)
}
