package dto

//CustomerSignUpRequest ...
type CustomerSignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Dob      string `json:"dob"`
}

//CustomerSignUpResponse ...
type CustomerSignUpResponse struct {
	CustomerInfo
}
