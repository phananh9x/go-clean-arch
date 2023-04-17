package jwt

// nolint
//go:generate mockgen -source=interfaces.go -destination=interfaces_mock.go -package=jwt

import (
	jwtgo "github.com/dgrijalva/jwt-go"
)

// IVerifier ...
type IVerifier interface {
	Verify(string, jwtgo.Claims) error
}
type ISigner interface {
	Sign(jwtgo.Claims) (string, error)
}
