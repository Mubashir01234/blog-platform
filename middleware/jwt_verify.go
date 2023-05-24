package middleware

import (
	"blog/config"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt" // Import golang jwt library
)

// IsAuthorized is a middleware function that checks if the request is authorized.
// It expects the JWT token to be included in the "Authorization" header as a Bearer token.
// The function validates the token and attaches the token claims to the request context.
// If the token is valid, the next handler is called; otherwise, an "Unauthorized" response is sent.
func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Split the Authorization header value to extract the JWT token
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		if len(authHeader) != 2 {
			// The Authorization header is malformed, respond with an error
			AuthorizationResponse("Malformed JWT token", w)
		} else {
			// Extract the JWT token from the second element of authHeader
			jwtToken := authHeader[1]

			// Parse and verify the JWT token
			token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				// Verify the token's signing method is HMAC
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				// Retrieve the signing key from the configuration
				return []byte(config.Cfg.JwtSecret), nil
			})

			// Check if the token is valid and extract the claims
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				// Attach the token claims to the request context
				ctx := context.WithValue(r.Context(), "props", claims)
				// Call the next handler with the updated context
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				// The token is unauthorized or invalid, respond with an error
				AuthorizationResponse("Unauthorized", w)
			}
		}
	})
}
