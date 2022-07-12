package serializers

type User struct {
	ID           *string `json:"id,omitempty"`
	Password     *string `json:"password,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
}

type Address struct {
	Line1      *string `json:"line1,omitempty"`
	Line2      *string `json:"line2,omitempty"`
	City       *string `json:"city,omitempty"`
	State      *string `json:"state,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	Country    *string `json:"country,omitempty"`
}

type ManagementCompany struct {
	ID      *string  `json:"id,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
	Website *string  `json:"website,omitempty"`
}
