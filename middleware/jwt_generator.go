package middleware

import (
	"blog/config"
	"time"

	"blog/models"

	"github.com/golang-jwt/jwt"
)

// Claims represents the custom claims for the JWT.
type claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserId   string `json:"user_id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT token based on the provided user information.
func GenerateJWT(user models.User) (*string, error) {
	// Create custom claims for the JWT
	claims := &claims{
		Username: user.Username,
		Email:    user.Email,
		UserId:   user.Id.Hex(),
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token expiration time (24 hours from now)
		},
	}

	// Create a new JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the JWT secret
	tokenString, err := token.SignedString([]byte(config.Cfg.JwtSecret))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
