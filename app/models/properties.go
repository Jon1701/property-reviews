package models

import "time"

type Property struct {
	ID int `gorm:"primaryKey"`

	// UUID for this Property.
	IDHash *string `gorm:"column:id_hash"`

	// Address Line 1.
	AddressLine1 *string `gorm:"column:address_line1"`

	// Address Line 2.
	AddressLine2 *string `gorm:"column:address_line2"`

	// Address City.
	AddressCity *string `gorm:"column:address_city"`

	// Address State.
	AddressState *string `gorm:"column:address_state"`

	// Address Postal Code.
	AddressPostalCode *string `gorm:"column:address_postal_code"`

	// Address Country.
	AddressCountry *string `gorm:"column:address_country"`

	// Property Type.
	PropertyType *string `gorm:"column:property_type"`

	// Building Type.
	BuildingType *string `gorm:"column:building_type"`

	// Neighborhood.
	Neighborhood *string `gorm:"column:neighborhood"`

	// ID Hash of the Management Company.
	ManagementCompanyIDHash *string `gorm:"column:management_company_id_hash"`

	// Creation Time.
	CreatedAt time.Time `gorm:"column:created_at"`

	// Update Time.
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type PropertyWithManageCompany struct {
	// UUID for this Property.
	PropertyIDHash *string `gorm:"column:properties_id_hash"`

	// Address Line 1.
	PropertyAddressLine1 *string `gorm:"column:properties_address_line1"`

	// Address Line 2.
	PropertyAddressLine2 *string `gorm:"column:properties_address_line2"`

	// Address City.
	PropertyAddressCity *string `gorm:"column:properties_address_city"`

	// Address State.
	PropertyAddressState *string `gorm:"column:properties_address_state"`

	// Address Postal Code.
	PropertyAddressPostalCode *string `gorm:"column:properties_address_postal_code"`

	// Address Country.
	PropertyAddressCountry *string `gorm:"column:properties_address_country"`

	// Property Type.
	PropertyType *string `gorm:"column:properties_property_type"`

	// Building Type.
	PropertyBuildingType *string `gorm:"column:properties_building_type"`

	// Neighborhood.
	PropertyNeighborhood *string `gorm:"column:properties_neighborhood"`

	// UUID for this Management Company.
	ManagementCompanyIDHash *string `gorm:"column:management_companies_id_hash"`

	// Management Company Name.
	ManagementCompanyName *string `gorm:"column:management_companies_name"`

	// Management Company Address Line 1.
	ManagementCompanyAddressLine1 *string `gorm:"column:management_companies_address_line1"`

	// Management Company Address Line 2.
	ManagementCompanyAddressLine2 *string `gorm:"column:management_companies_address_line2"`

	// Management Company Address City.
	ManagementCompanyAddressCity *string `gorm:"column:management_companies_address_city"`

	// Management Company Address State.
	ManagementCompanyAddressState *string `gorm:"column:management_companies_address_state"`

	// Management Company Address Postal Code.
	ManagementCompanyAddressPostalCode *string `gorm:"column:management_companies_address_postal_code"`

	// Management Company Address Country.
	ManagementCompanyAddressCountry *string `gorm:"column:management_companies_address_country"`

	// Management Company Website.
	ManagementCompanyWebsite *string `gorm:"column:management_companies_website"`
}
