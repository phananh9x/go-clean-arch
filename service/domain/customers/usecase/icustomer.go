package usecase

import (
	"context"
	"go-clean-arch/config"
	"go-clean-arch/pkg/jwt"
	dto2 "go-clean-arch/service/models/dto"
	"go-clean-arch/service/repository"
)

type ICustomerUseCase interface {
	GetCustomerInfo(context context.Context, customerId string) (dto2.CustomerInfo, error)
	GetAccessToken(context context.Context, username string, password string) (string, error)
	SignUp(context context.Context, customer dto2.CustomerSignUpRequest) (dto2.CustomerSignUpResponse, error)
}

type customerUseCase struct {
	signer jwt.ISigner
	config *config.AppConfig
	repo   repository.IRepo
}

func NewCustomerUseCase(config *config.AppConfig, signer jwt.ISigner, repo repository.IRepo) ICustomerUseCase {
	return &customerUseCase{
		config: config,
		signer: signer,
		repo:   repo,
	}
}
