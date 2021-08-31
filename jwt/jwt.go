// https://pkg.go.dev/github.com/golang-jwt/jwt
package main

import (
	"crypto/rsa"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	// rsa
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey

	// hmac
	hmacSampleSecret []byte
)

type CustomerInfo struct {
	Name string
	Kind string
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	CustomerInfo
}

func CreateToken(user string) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))

	t.Claims = &CustomClaimsExample{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
		},
		CustomerInfo{user, "human"},
	}

	return t.SignedString(hmacSampleSecret)
}

func ParseToken(tokenString string) (*CustomClaimsExample, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return hmacSampleSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaimsExample); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
