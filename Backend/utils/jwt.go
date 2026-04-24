package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SECRET = []byte("SUPER_SECRET_KEY")

func GenerateAccessToken(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SECRET)
}

func GenerateRefreshToken(userID int) (string, time.Time, error) {
	exp := time.Now().Add(7 * 24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     exp.Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SECRET)
	return token, exp, err
}