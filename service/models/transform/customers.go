package transform

import (
	"go-clean-arch/service/models/dto"
	"go-clean-arch/service/models/entities"
)

//ToCustomerSignUpResponse ...
func ToCustomerSignUpResponse(customer *entities.Customers) dto.CustomerSignUpResponse {
	if customer == nil {
		return dto.CustomerSignUpResponse{}
	}
	return dto.CustomerSignUpResponse{
		CustomerInfo: dto.CustomerInfo{
			CustomerID: customer.CustomerId,
			Name:       customer.Username,
			Email:      customer.Email,
			Phone:      customer.Phone,
			Dob:        customer.Dob,
		},
	}
}

//ToCustomerLoginResponse ...
func ToCustomerLoginResponse(customer *entities.Customers, token string) dto.CustomerLoginResponse {
	if customer == nil {
		return dto.CustomerLoginResponse{}
	}
	return dto.CustomerLoginResponse{
		CustomerInfo: dto.CustomerInfo{
			CustomerID: customer.CustomerId,
			Name:       customer.Username,
			Email:      customer.Email,
			Phone:      customer.Phone,
			Dob:        customer.Dob,
		},
		AccessToken: token,
	}
}

//ToCustomerInfoResponse ...
func ToCustomerInfoResponse(customer *entities.Customers) dto.CustomerInfo {
	if customer == nil {
		return dto.CustomerInfo{}
	}
	return dto.CustomerInfo{
		CustomerID: customer.CustomerId,
		Name:       customer.Username,
		Email:      customer.Email,
		Phone:      customer.Phone,
		Dob:        customer.Dob,
		Username:   customer.Username,
	}
}
