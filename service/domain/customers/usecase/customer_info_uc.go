package usecase

import (
	"context"
	"go-clean-arch/service/models/dto"
)

func (c customerUseCase) GetCustomerInfo(context context.Context, customerId string) (dto.CustomerInfo, error) {
	//Todo: implement your business logic here. and remove mock data below
	return dto.CustomerInfo{
		CustomerID: customerId,
		Name:       "John Doe",
		Email:      "john@gmail.com",
		Phone:      "0123456789",
		Dob:        "1990-01-01",
		CreatedAt:  "2020-01-01",
	}, nil
}
