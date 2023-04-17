package middleware

import (
	"errors"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-clean-arch/pkg/ginwrapper"
	"go-clean-arch/pkg/jwt"
	"strings"
)

var (
	ErrAuthorizationNotFound = errors.New("No Authorization header provided")
	ErrTokenNotFound         = errors.New("Could not find bearer token in Authorization header")
)

//AccessTokenClaims ...
type AccessTokenClaims struct {
	jwtgo.StandardClaims
	UserID string `json:"customer_id"`
}

// WebHookAuthentication return middleware for authentication of API
func WebHookAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// APIAuthentication return middleware for authentication of API
func APIAuthentication(verifier jwt.IVerifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &ginwrapper.Context{Context: c}
		tokenBearer := c.Request.Header.Get("Authorization")
		if tokenBearer == "" {
			ctx.BadRequest(ErrAuthorizationNotFound)
			return
		}
		token := strings.TrimPrefix(tokenBearer, "Bearer ")
		if token == tokenBearer {
			ctx.BadRequest(ErrTokenNotFound)
			return
		}

		claims := &AccessTokenClaims{}
		err := verifier.Verify(token, claims)

		if err != nil {
			if jwt.IsExpiredJWTError(err) {
				ctx.TokenExpired()
			} else {
				ctx.AccessDenied()
			}
			return
		}
		ctx.SetUserID(claims.UserID)
		ctx.SetToken(token)

		c.Next()
	}
}
