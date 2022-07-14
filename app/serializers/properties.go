package serializers

type PropertyType string
type BuildingType string

var TypesOfProperties = map[string]PropertyType{
	"CommercialProperty":  "commercial",
	"ResidentialProperty": "residential",
}

var TypesOfBuildings = map[string]BuildingType{
	"Apartment":     "apartment",
	"BasementSuite": "basement-suite",
	"Bungalow":      "bungalow",
	"Cabin":         "cabin",
	"Condominium":   "condominium",
	"CoOp":          "co-op",
	"Cottage":       "cottage",
	"DetachedHouse": "detached",
	"Townhome":      "townhome",
}

type Property struct {
	ID                *string            `json:"id,omitempty"`
	Address           *Address           `json:"address,omitempty"`
	Neighborhood      *string            `json:"neighborhood,omitempty"`
	PropertyType      *PropertyType      `json:"propertyType,omitempty"`
	BuildingType      *BuildingType      `json:"buildingType,omitempty"`
	ManagementCompany *ManagementCompany `json:"managementCompany,omitempty"`
}
