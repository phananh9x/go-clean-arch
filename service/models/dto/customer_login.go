package dto

//CustomerLoginRequest ...
type CustomerLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//CustomerLoginResponse ...
type CustomerLoginResponse struct {
	CustomerInfo CustomerInfo `json:"customer_info"`
	AccessToken  string       `json:"access_token"`
}
