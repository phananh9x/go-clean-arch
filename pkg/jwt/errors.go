package jwt

import (
	jwtgo "github.com/dgrijalva/jwt-go"
)

// IsExpiredJWTError checks if err is JWT ValidationErrorExpired
func IsExpiredJWTError(err error) bool {
	if err == nil {
		return false
	}

	switch v := err.(type) {
	case *jwtgo.ValidationError:
		return (v != nil) && (v.Errors&jwtgo.ValidationErrorExpired != 0)
	default:
		return false
	}
}
