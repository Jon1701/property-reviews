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
		return
	}

	// If a Management Company ID was provided, check if it exists in the table.
	company := models.ManagementCompany{}
	if property.ManagementCompany != nil && property.ManagementCompany.ID != nil {
		result := appCtx.DB.Where("id_hash = ?", property.ManagementCompany.ID).First(&company)
		if result.Error != nil {
			validationResults := validation.Property{
				ManagementCompany: &validation.ManagementCompany{
					ID: &errormessages.ManagementCompanyIDNotFound,
				},
			}

			body, err := json.Marshal(validationResults)
			if err != nil {
				panic(fmt.Sprintf("Failed to marshal the field validation results struct into JSON: %+v\n", err))
			}

			c.Data(http.StatusBadRequest, "application/json", body)
			return
		}
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
	if property.ManagementCompany != nil && property.ManagementCompany.ID != nil {
		m.ManagementCompanyIDHash = property.ManagementCompany.ID
	}
	result := appCtx.DB.Create(&m)
	if result.Error != nil {
		panic(fmt.Sprintf("Failed to persist Property in database: %+v\n", result.Error))
	}

	c.Header("Location", fmt.Sprintf("/api/property/%s", idHash))
	c.JSON(http.StatusNoContent, nil)
}

// Gets all Properties.
func (appCtx *AppContext) GetProperties(c *gin.Context) {
	results, err := appCtx.DBGetPropertiesSerialized()
	if err != nil {
		panic(fmt.Sprintf("Failed to serialize Properties from database: %+v\n", err))
	}

	body, err := json.Marshal(&results)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal Properties model to JSON: %+v\n", err))
	}

	c.Data(http.StatusOK, "application/json", body)
}

// Gets a Property by ID.
func (appCtx *AppContext) GetPropertyByID(c *gin.Context) {
	propertyID := c.Param("propertyID")

	// Check if Property exists in the database.
	m := models.Property{}
	result := appCtx.DBCheckIfPropertyIDExists(propertyID, &m)
	if !result {
		c.JSON(http.StatusNotFound, &gin.H{
			"message": errormessages.PropertyNotFound,
		})
		return
	}

	// Get Property by ID.
	property, err := appCtx.DBGetPropertySerializedByID(propertyID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get Property from database: %+v\n", err))
	}

	// Convert Property Model to JSON.
	body, err := json.Marshal(property)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal Property model into JSON: %+v\n", err))
	}

	c.Data(http.StatusOK, "application/json", body)
}

// Updates a Property.
func (appCtx *AppContext) UpdateProperty(c *gin.Context) {
	propertyID := c.Param("propertyID")

	// Check if Property exists in the database.
	m := models.Property{}
	result := appCtx.DBCheckIfPropertyIDExists(propertyID, &m)
	if !result {
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

	// Check if the Management Company ID Hash provided.
	isManagementCompanyIDHashProvided := property.ManagementCompany != nil && property.ManagementCompany.ID != nil && len(*property.ManagementCompany.ID) > 0

	// If a Management Company ID was provided, check if it exists in the table.
	company := models.ManagementCompany{}
	if isManagementCompanyIDHashProvided {
		result := appCtx.DB.Where("id_hash = ?", property.ManagementCompany.ID).First(&company)
		if result.Error != nil {
			validationResults := validation.Property{
				ManagementCompany: &validation.ManagementCompany{
					ID: &errormessages.ManagementCompanyIDNotFound,
				},
			}

			body, err := json.Marshal(validationResults)
			if err != nil {
				panic(fmt.Sprintf("Failed to marshal the field validation results struct into JSON: %+v\n", err))
			}

			c.Data(http.StatusBadRequest, "application/json", body)
			return
		}
	}

	// Persist updated Property into database.
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
	if isManagementCompanyIDHashProvided {
		m.ManagementCompanyIDHash = property.ManagementCompany.ID
	}
	resultUpdate := appCtx.DB.Updates(&m)
	if resultUpdate.Error != nil {
		panic(fmt.Sprintf("Failed to update Property in database: %+v\n", resultUpdate.Error))
	}
	// If the Management Company ID is not provided, remove it from the Property
	// in the database.
	if !isManagementCompanyIDHashProvided {
		result := appCtx.DB.Model(&m).Select("management_company_id_hash").Updates(map[string]interface{}{"management_company_id_hash": nil})
		if result.Error != nil {
			panic(fmt.Sprintf("Failed to delete the Management Company ID in database: %+v\n", result.Error))
		}
	}

	// Get the Property row and serialize it for JSON.
	s, err := appCtx.DBGetPropertySerializedByID(propertyID)
	if err != nil {
		panic(fmt.Sprintf("Failed to get the Property from the database: %+v", err))
	}

	// Convert to JSON.
	body, err := json.Marshal(s)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal the database row struct into JSON: %+v\n", err))
	}

	c.Data(http.StatusOK, "application/json", body)
}

// Checks if a given Property ID exists.
func (appCtx *AppContext) DBCheckIfPropertyIDExists(propertyID string, m *models.Property) bool {
	result := appCtx.DB.Where("id_hash = ?", propertyID).First(&m)

	return result.Error == nil
}

// Gets a Property and Serializes it.
func (appCtx *AppContext) DBGetPropertySerializedByID(propertyID string) (*serializers.Property, error) {
	// Get the updated Property from the database.
	m := &models.PropertyWithManageCompany{}

	result := appCtx.DB.Raw(fmt.Sprintf(`
		SELECT
			properties.id_hash							AS properties_id_hash,
			properties.address_line1 				AS properties_address_line1,
			properties.address_line2 				AS properties_address_line2,
			properties.address_city 				AS properties_address_city,
			properties.address_state 				AS properties_address_state,
			properties.address_postal_code 	AS properties_address_postal_code,
			properties.address_country 			AS properties_address_country,
			properties.property_type 				AS properties_property_type,
			properties.building_type 				AS properties_building_type,
			properties.neighborhood 				AS properties_neighborhood,
			management_companies.id_hash 		AS management_companies_id_hash,
			management_companies.name 			AS management_companies_name,
			management_companies.address_line1 				AS management_companies_address_line1,
			management_companies.address_line2 				AS management_companies_address_line2,
			management_companies.address_city 				AS management_companies_address_city,
			management_companies.address_state 				AS management_companies_address_state,
			management_companies.address_postal_code 	AS management_companies_address_postal_code,
			management_companies.address_country 			AS management_companies_address_country,
			management_companies.website 							AS management_companies_website	
		FROM
			properties
		LEFT JOIN
			management_companies
		ON
			properties.management_company_id_hash = management_companies.id_hash
		WHERE
			properties.id_hash = '%s';`, propertyID)).Find(&m)
	if result.Error != nil {
		return nil, result.Error
	}

	// Serialize the Property Model.
	property := appCtx.DBSerializePropertyModel(*m)

	return property, nil
}

// Gets all Properties and Serializes them.
func (appCtx *AppContext) DBGetPropertiesSerialized() (*[]serializers.Property, error) {
	propertiesModel := []models.PropertyWithManageCompany{}

	result := appCtx.DB.Raw(`
	SELECT
		properties.id_hash							AS properties_id_hash,
		properties.address_line1 				AS properties_address_line1,
		properties.address_line2 				AS properties_address_line2,
		properties.address_city 				AS properties_address_city,
		properties.address_state 				AS properties_address_state,
		properties.address_postal_code 	AS properties_address_postal_code,
		properties.address_country 			AS properties_address_country,
		properties.property_type 				AS properties_property_type,
		properties.building_type 				AS properties_building_type,
		properties.neighborhood 				AS properties_neighborhood,
		management_companies.id_hash 		AS management_companies_id_hash,
		management_companies.name 			AS management_companies_name,
		management_companies.address_line1 				AS management_companies_address_line1,
		management_companies.address_line2 				AS management_companies_address_line2,
		management_companies.address_city 				AS management_companies_address_city,
		management_companies.address_state 				AS management_companies_address_state,
		management_companies.address_postal_code 	AS management_companies_address_postal_code,
		management_companies.address_country 			AS management_companies_address_country,
		management_companies.website 							AS management_companies_website	
	FROM
		properties
	LEFT JOIN
		management_companies
	ON
		properties.management_company_id_hash = management_companies.id_hash
	`).Scan(&propertiesModel)
	if result.Error != nil {
		return nil, result.Error
	}

	// Serialize the model for JSON.
	propertiesSerialized := []serializers.Property{}
	for _, item := range propertiesModel {
		property := appCtx.DBSerializePropertyModel(item)
		propertiesSerialized = append(propertiesSerialized, *property)
	}

	return &propertiesSerialized, nil
}

// Serializes a Property Model for JSON.
func (appCtx *AppContext) DBSerializePropertyModel(m models.PropertyWithManageCompany) *serializers.Property {
	property := &serializers.Property{
		ID: m.PropertyIDHash,
		Address: &serializers.Address{
			Line1:      m.PropertyAddressLine1,
			Line2:      m.PropertyAddressLine2,
			City:       m.PropertyAddressCity,
			State:      m.PropertyAddressState,
			PostalCode: m.PropertyAddressPostalCode,
			Country:    m.PropertyAddressCountry,
		},
		BuildingType: (*serializers.BuildingType)(m.PropertyBuildingType),
		PropertyType: (*serializers.PropertyType)(m.PropertyType),
		Neighborhood: m.PropertyNeighborhood,
	}
	if m.ManagementCompanyIDHash != nil {
		property.ManagementCompany = &serializers.ManagementCompany{
			ID:   m.ManagementCompanyIDHash,
			Name: m.ManagementCompanyName,
			Address: &serializers.Address{
				Line1:      m.ManagementCompanyAddressLine1,
				Line2:      m.ManagementCompanyAddressLine2,
				City:       m.ManagementCompanyAddressCity,
				State:      m.ManagementCompanyAddressState,
				PostalCode: m.ManagementCompanyAddressPostalCode,
				Country:    m.ManagementCompanyAddressCountry,
			},
		}
	}

	return property
}
