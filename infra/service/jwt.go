package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jonathanlazaro1/stone-challenge/config"
)

// JwtIssuer is the issuer of the JWT token
const JwtIssuer = "Invoice API"

// JwtExpInHours is the time, in hours, that the JWT token is valid for
const JwtExpInHours = 10

// GenerateJWT gerenates a new JWT, given an email and name
func GenerateJWT(email string, name string) (string, error) {
	config := config.GetConfig()
	now := time.Now().Local().Add(time.Hour * time.Duration(JwtExpInHours))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  JwtIssuer,
		"sub":  email,
		"exp":  now.Unix(),
		"name": name})

	tokenString, err := token.SignedString(config.AuthSecret)
	return tokenString, err
}

// DecodeJWT decodes a string encoded in JWT format
func DecodeJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return token, nil
	})
	return token, err
}
