package utils

import (
	"fmt"
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

func ValidateToken(tokenJWT string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenJWT, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("method tanda tangan tidak valid")
		}
		return []byte(config.Env.SecretKey), nil
	})

	// cek token valid
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("token tidak valid")
	}

	// ambil claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("gagal mendapatkan claims")
	}

	return claims, nil
}
