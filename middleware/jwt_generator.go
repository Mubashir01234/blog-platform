package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type claims struct {
	Username string `json:"username"`
	Dash     string `json:"dash"`
	// Email    string `json:"email"`
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (string, error) {
	claims := &claims{
		Username: user.Username,
		Dash:     user.Dash,
		// Email:    user.Email,
		UserID: user.ID.Hex(),
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

	tokenString, err := token.SignedString(JWT_SECRET)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
