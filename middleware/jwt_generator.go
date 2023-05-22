package middleware

import (
	"blog/config"
	"time"

	"blog/models"

	"github.com/golang-jwt/jwt"
)

type claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserId   string `json:"user_id"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (*string, error) {
	claims := &claims{
		Username: user.Username,
		Email:    user.Email,
		UserId:   user.Id.Hex(),
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Cfg.JwtSecret))

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
