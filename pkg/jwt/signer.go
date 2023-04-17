package jwt

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	jwtgo "github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

type Signer struct {
	privateKey    *rsa.PrivateKey
	signingMethod jwtgo.SigningMethod
	hash          crypto.Hash
}

func NewSigner(privateKeyBase64 string, signingMethodAlg string, hash crypto.Hash) (*Signer, error) {
	decoded, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		zap.S().Warnf("decode private key failed with error: %+v", err)
		return nil, err
	}
	privateKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(decoded)
	if err != nil {
		return nil, err
	}
	signingMethod := jwtgo.GetSigningMethod(signingMethodAlg)
	return &Signer{
		privateKey:    privateKey,
		signingMethod: signingMethod,
		hash:          hash,
	}, nil
}

func (s *Signer) Sign(claims jwtgo.Claims) (string, error) {
	token := jwtgo.NewWithClaims(s.signingMethod, claims)
	return token.SignedString(s.privateKey)
}
