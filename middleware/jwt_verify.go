package middleware

import (
	"blog/config"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

// IsAuthorized is a middleware function that checks if the request is authorized.
// It expects the JWT token to be included in the "Authorization" header as a Bearer token.
// The function validates the token and attaches the token claims to the request context.
// If the token is valid, the next handler is called; otherwise, an "Unauthorized" response is sent.
func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		if len(authHeader) != 2 {
			AuthorizationResponse("Malformed JWT token", w)
		} else {
			jwtToken := authHeader[1]
			token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(config.Cfg.JwtSecret), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims) // Attach token claims to the request context
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				AuthorizationResponse("Unauthorized", w)
			}
		}
	})
}
