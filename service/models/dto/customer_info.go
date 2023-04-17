package dto

//CustomerInfo ...
type CustomerInfo struct {
	CustomerID string `json:"customer_id"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Dob        string `json:"dob"`
	CreatedAt  string `json:"created_at"`
}
