package usecase

import (
	"context"
	"go-clean-arch/service/models/dto"
	"go-clean-arch/service/models/transform"
)

func (c customerUseCase) GetCustomerInfo(context context.Context, customerId string) (dto.CustomerInfo, error) {
	result, err := c.repo.NewCustomersRepository().GetByID(context, customerId)
	if err != nil {
		return dto.CustomerInfo{}, err
	}
	if result == nil {
		return dto.CustomerInfo{}, nil
	}

	return transform.ToCustomerInfoResponse(result), nil
}
