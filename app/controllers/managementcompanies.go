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
	"gorm.io/gorm"
)

// Generates a hyphenless UUID with "management_" prepended.
func generateManagementIDHash() string {
	id := strings.ReplaceAll(uuid.New().String(), "-", "")

	return fmt.Sprintf("management_%s", id)
}

// Creates a Management Company.
func (appCtx *AppContext) CreateManagementCompany(c *gin.Context) {
	company := serializers.ManagementCompany{}

	// Serialize the request body, respond with Bad Request if unable to.
	err := SerializeRequestBody(c, &company)
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

// Gets a list of Management Companies.
func (appCtx *AppContext) GetManagementCompanies(c *gin.Context) {
	// Get pagination parameters.
	params := PaginationParameters{}
	SerializePaginationParameters(c, &params)

	// Get paginated Management Company results.
	results, err := appCtx.DBGetManagementCompanies(params.AfterID, params.BeforeID, params.Limit)
	if err != nil {
		panic(fmt.Sprintf("Failed to get Management Companies: %+v\n", err))
	}

	// Serialize results into JSON.
	body, err := json.Marshal(results)
	if err != nil {
		panic(fmt.Sprintf("Failed to serialize Management Company results into JSON: %+v\n", err))
	}

	c.Data(http.StatusOK, "application/json", body)
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

	// Serialize the request body, respond with Bad Request if unable to.
	err := SerializeRequestBody(c, &company)
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

		c.Data(http.StatusBadRequest, "application/json", body)
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
		panic(fmt.Sprintf("Failed to persist Management Company in database: %+v\n", result.Error))
	}

	// Get updated row.
	m = models.ManagementCompany{}
	result = appCtx.DB.Where("id_hash = ?", managementID).First(&m)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to get Management Company from database: %+v\n", result.Error))
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

// Gets a Management Company from the database.
func (appCtx *AppContext) DBGetManagementCompanyByID(managementID string) (*models.ManagementCompany, error) {
	m := models.ManagementCompany{}
	result := appCtx.DB.Raw(fmt.Sprintf(`
		SELECT *
		FROM management_companies
		WHERE id_hash = '%s';
	`, managementID)).Scan(&m)
	if result.Error != nil {
		return nil, result.Error
	}

	return &m, nil
}

// Gets an array of Management Companies.
func (appCtx *AppContext) DBGetManagementCompanies(afterID *string, beforeID *string, paramLimit *uint64) (*[]serializers.ManagementCompany, error) {
	var limit uint64 = uint64(20)
	if *paramLimit > 0 {
		limit = *paramLimit
	}

	// Get Management Companies from the database.
	var result *gorm.DB
	m := []models.ManagementCompany{}
	if afterID != nil && beforeID != nil { // Between.
		result = appCtx.DB.Raw(fmt.Sprintf(`
			SELECT * FROM	management_companies
			WHERE id BETWEEN
				(SELECT id	FROM management_companies WHERE id_hash = '%s')
					AND
				(SELECT id	FROM management_companies WHERE id_hash = '%s')
			LIMIT %d;
	`, *afterID, *beforeID, limit)).Scan(&m)
	} else if afterID != nil && beforeID == nil { // After.
		result = appCtx.DB.Raw(fmt.Sprintf(`
			SELECT * FROM	management_companies
			WHERE id > (SELECT id FROM management_companies WHERE id_hash = '%s')
			LIMIT %d;
		`, *afterID, limit)).Scan(&m)
	} else if afterID == nil && beforeID != nil { // Before.
		result = appCtx.DB.Raw(fmt.Sprintf(`
			SELECT * FROM	management_companies
			WHERE id < (SELECT id FROM management_companies WHERE id_hash = '%s')
			LIMIT %d;
		`, *beforeID, limit)).Scan(&m)
	} else {
		result = appCtx.DB.Raw(fmt.Sprintf(`
			SELECT * FROM management_companies
			LIMIT %d
		`, limit)).Scan(&m)
	}
	if result.Error != nil {
		return nil, result.Error
	}

	// Serialize results.
	companiesSerialized := []serializers.ManagementCompany{}
	for _, item := range m {
		property := appCtx.DBSerializeManagementCompanyModel(item)
		companiesSerialized = append(companiesSerialized, *property)
	}

	return &companiesSerialized, nil
}

// Updates a Management Company in the database.
func (appCtx *AppContext) DBUpdateManagementCompany(managementID string, s *serializers.ManagementCompany) error {
	m := models.ManagementCompany{
		Name:    s.Name,
		Website: s.Website,
	}
	if s.Address != nil {
		m.AddressLine1 = s.Address.Line1
		m.AddressLine2 = s.Address.Line2
		m.AddressCity = s.Address.City
		m.AddressState = s.Address.State
		m.AddressPostalCode = s.Address.PostalCode
		m.AddressCountry = s.Address.Country
	}
	result := appCtx.DB.Model(&models.ManagementCompany{}).Where("id_hash = ?", managementID).Updates(&m)
	return result.Error
}

// Serializes a Management Company Model into its Serializer.
func (appCtx *AppContext) DBSerializeManagementCompanyModel(m models.ManagementCompany) *serializers.ManagementCompany {
	return &serializers.ManagementCompany{
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
}
