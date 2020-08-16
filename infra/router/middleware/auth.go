package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

// ContextAuthClaims hold claims that are passed to the Context, when an request is authenticated
type ContextAuthClaims struct {
	Name  string
	Email string
}

// ReqAuthInfoName is the special type for the context Auth info name, derived from string
type ReqAuthInfoName string

// RequestAuthInfo is the name of the auth info on Context
const RequestAuthInfo ReqAuthInfoName = "auth"

// AddJwtAuthentication returns an auth handler that searches for and validates a JWT on every request
func AddJwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed or not present token"))
		} else {

			token, _ := service.DecodeJWT(authHeader[1])
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && claims.VerifyIssuer(service.JwtIssuer, true) {
				ctxClaims := ContextAuthClaims{}
				for k, v := range claims {
					if k == "sub" {
						ctxClaims.Email = v.(string)
					}
					if k == "name" {
						ctxClaims.Name = v.(string)
					}
				}
				ctx := context.WithValue(r.Context(), RequestAuthInfo, ctxClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}
