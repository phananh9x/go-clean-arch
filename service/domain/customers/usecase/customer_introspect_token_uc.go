package usecase

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"go-clean-arch/pkg/middleware"
	"time"
)

func (c customerUseCase) GetAccessToken(context context.Context, customerId string, password string) (string, error) {
	// TODO: implement your business logic here
	claims := middleware.AccessTokenClaims{
		UserID: customerId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(c.config.Authen.APIAuthenticator.JWTExpireTime)).Unix(),
		},
	}
	//Sign the token with our secret
	token, err := c.signer.Sign(claims)
	if err != nil {
		return "", err
	}
	return token, nil

}
