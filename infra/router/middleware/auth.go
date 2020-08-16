package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/jonathanlazaro1/stone-challenge/infra/service"
)

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
			claims, err := service.DecodeJWT(authHeader[1])
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			} else if claims != nil {
				ctx := context.WithValue(r.Context(), RequestAuthInfo, struct {
					Name  string
					Email string
				}{Name: claims.Name, Email: claims.Subject})
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}
