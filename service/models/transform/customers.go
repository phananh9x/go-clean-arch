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
