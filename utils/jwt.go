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
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	SignedToken, err := token.SignedString([]byte(config.Env.SecretKey))

	if err != nil {
		return "", err
	}

	return SignedToken, nil
}

func ValidateToken(tokenJWT string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenJWT, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("ValidateToken: invalid signature method")
		}
		return []byte(config.Env.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("ValidateToken: invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("ValidateToken: failed to get claims")
	}

	return claims, nil
}

func AssertTypeClaims(claims any) (jwt.MapClaims, error) {
	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("AssertTypeClaims: claims format is invalid")
	}
	return claimsMap, nil
}
