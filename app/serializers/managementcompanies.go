package serializers

type ManagementCompany struct {
	ID      *string  `json:"id,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
	Website *string  `json:"website,omitempty"`
}
