package models

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
}
