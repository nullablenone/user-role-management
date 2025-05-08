package utils

import (
	"manajemen-user/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Valid 1 hari
	}

	// Membuat Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tanda tangani Token
	SignedToken, err := token.SignedString([]byte(config.Env.SecretKey))

	if err != nil {
		return "", err
	}

	return SignedToken, nil
}
