package service

import (
	"errors"
	"time"

	"github.com/jonathanlazaro1/stone-challenge/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/jonathanlazaro1/stone-challenge/config"
)

// JwtClaims are the claims that are expected to be present on a JWT token sent to the application
type jwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// jwtIssuer is the issuer of the JWT token
const jwtIssuer = "Invoice API"

// jwtExpInHours is the time, in hours, that the JWT token is valid for
const jwtExpInHours = 10

// GenerateJWT gerenates a new JWT, given an email and name
func GenerateJWT(email string, name string) (string, error) {
	cfg := config.GetConfig()
	now := time.Now().Local().Add(time.Hour * time.Duration(jwtExpInHours))

	claims := jwtClaims{
		name,
		jwt.StandardClaims{
			Subject:   email,
			Issuer:    jwtIssuer,
			ExpiresAt: now.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfg.AppAuthSecret))
	return tokenString, err
}

// DecodeJWT decodes a string encoded in JWT format
func DecodeJWT(tokenString string) (*domain.AuthInfo, error) {
	cfg := config.GetConfig()
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.AppAuthSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwtClaims)
	err = claims.Valid()
	if err != nil {
		return nil, err
	}

	if !claims.VerifyIssuer(jwtIssuer, true) {
		err = errors.New("Issuer is invalid")
		return nil, err
	}
	return &domain.AuthInfo{Name: claims.Name, Email: claims.Subject}, nil
}
