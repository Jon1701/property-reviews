package serializers

type User struct {
	ID           *string `json:"id,omitempty"`
	Password     *string `json:"password,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
}
