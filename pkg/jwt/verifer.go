package jwt

import (
	"crypto/rsa"
	"encoding/base64"

	jwtgo "github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

const (
	// DefaultSigningMethodAlg ...
	DefaultSigningMethodAlg = "RS256"
)

// Verifier ...
type Verifier struct {
	publicKey     *rsa.PublicKey
	signingMethod jwtgo.SigningMethod
}

// NewVerifier ...
func NewVerifier(publicKeyBase64 string, signingMethodAlg string) (*Verifier, error) {
	decoded, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		zap.S().Warnf("decode publickey failed with error: %+v", err)
		return nil, err
	}

	publicKey, err := jwtgo.ParseRSAPublicKeyFromPEM(decoded)
	if err != nil {
		zap.S().Warnf("parse publickey failed with error: %+v", err)
		return nil, err
	}

	signingMethod := jwtgo.GetSigningMethod(signingMethodAlg)

	return &Verifier{
		publicKey:     publicKey,
		signingMethod: signingMethod,
	}, nil
}

func (v *Verifier) verifyKeyFunc() jwtgo.Keyfunc {
	fn := func(*jwtgo.Token) (interface{}, error) {
		return v.publicKey, nil
	}
	return fn
}

// Verify ...
func (v *Verifier) Verify(signed string, claims jwtgo.Claims) error {
	_, err := jwtgo.ParseWithClaims(signed, claims, v.verifyKeyFunc())
	if err != nil {
		zap.S().Warnf("failed to parse token with error = %+v", err)
		return err
	}
	return nil
}
