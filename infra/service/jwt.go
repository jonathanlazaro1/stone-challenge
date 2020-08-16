package service

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jonathanlazaro1/stone-challenge/config"
)

// JwtClaims are the claims that are expected to be present on a JWT token sent to the application
type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// JwtIssuer is the issuer of the JWT token
const JwtIssuer = "Invoice API"

// JwtExpInHours is the time, in hours, that the JWT token is valid for
const JwtExpInHours = 10

// GenerateJWT gerenates a new JWT, given an email and name
func GenerateJWT(email string, name string) (string, error) {
	config := config.GetConfig()
	now := time.Now().Local().Add(time.Hour * time.Duration(JwtExpInHours))

	claims := JwtClaims{
		name,
		jwt.StandardClaims{
			Subject:   email,
			Issuer:    JwtIssuer,
			ExpiresAt: now.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.AuthSecret))
	return tokenString, err
}

// DecodeJWT decodes a string encoded in JWT format
func DecodeJWT(tokenString string) (*JwtClaims, error) {
	config := config.GetConfig()
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AuthSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*JwtClaims)
	err = claims.Valid()
	if err != nil {
		return nil, err
	}

	if !claims.VerifyIssuer(JwtIssuer, true) {
		err = errors.New("Issuer is invalid")
		return nil, err
	}
	return claims, nil
}
