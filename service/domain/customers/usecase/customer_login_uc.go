package usecase

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go-clean-arch/pkg/middleware"
	"go-clean-arch/service/models/dto"
	"go-clean-arch/service/models/transform"
	"time"
)

var (
	ErrCustomerNotFound = errors.New("customer not found")
	ErrInvalidPassword  = errors.New("invalid password")
)

func (c customerUseCase) Login(ctx context.Context, req dto.CustomerLoginRequest) (dto.CustomerLoginResponse, error) {
	// find customer by username
	result, err := c.repo.NewCustomersRepository().GetByUserName(ctx, req.Username)
	if err != nil {
		return dto.CustomerLoginResponse{}, ErrCustomerNotFound
	}
	if result == nil {
		return dto.CustomerLoginResponse{}, ErrCustomerNotFound
	}
	// compare password
	if !result.ComparePassword(req.Password) {
		return dto.CustomerLoginResponse{}, ErrInvalidPassword
	}
	// generate token
	claims := middleware.AccessTokenClaims{
		UserID: result.CustomerId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(c.config.Authen.APIAuthenticator.JWTExpireTime)).Unix(),
		},
	}
	//Sign the token with our secret
	token, err := c.signer.Sign(claims)
	if err != nil {
		return dto.CustomerLoginResponse{}, err
	}
	return transform.ToCustomerLoginResponse(result, token), nil

}
