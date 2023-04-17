package dto

//CustomerInfo ...
type CustomerInfo struct {
	CustomerID string `json:"customer_id"`
	Username   string `json:"username,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Dob        string `json:"dob,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
}
