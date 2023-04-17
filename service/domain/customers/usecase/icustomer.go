package usecase

import (
	"context"
	"go-clean-arch/config"
	"go-clean-arch/pkg/jwt"
	"go-clean-arch/service/models/dto"
	"go-clean-arch/service/repository"
)

type ICustomerUseCase interface {
	GetCustomerInfo(context context.Context, customerId string) (dto.CustomerInfo, error)
	Login(context context.Context, request dto.CustomerLoginRequest) (dto.CustomerLoginResponse, error)
	SignUp(context context.Context, customer dto.CustomerSignUpRequest) (dto.CustomerSignUpResponse, error)
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
