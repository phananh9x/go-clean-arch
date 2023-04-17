package usecase

import (
	"context"
	"github.com/google/uuid"
	"go-clean-arch/service/models/dto"
	"go-clean-arch/service/models/entities"
	"go-clean-arch/service/models/transform"
	"time"
)

func (c customerUseCase) SignUp(ctx context.Context, customer dto.CustomerSignUpRequest) (dto.CustomerSignUpResponse, error) {
	customerBuilder := entities.NewCustomersBuilder().
		WithCustomerID(uuid.New().String()).
		WithUsername(customer.Username).
		WithPassword(customer.Password, uuid.New().String()).
		WithName(customer.Name).
		WithEmail(customer.Email).
		WithPhone(customer.Phone).
		WithDOB(customer.Dob).
		WithCreatedAt(time.Now().Unix()).
		WithUpdatedAt(time.Now().Unix()).
		Build()
	err := c.repo.NewCustomersRepository().Create(ctx, customerBuilder)
	if err != nil {
		return dto.CustomerSignUpResponse{}, err
	}
	return transform.ToCustomerSignUpResponse(customerBuilder), nil
}
