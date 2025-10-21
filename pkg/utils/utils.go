package utils

import (
	"fmt"
	"time"
	"tutorial/pkg/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUUID() string {
	return uuid.New().String()
}

func GenerateJWTToken(userID string) (string, error) {
	cfg := config.GetConfig()
	if cfg == nil {
		return "", fmt.Errorf("config not loaded")
	}

	secretKey := cfg.JWT.SecretKey

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * time.Duration(cfg.JWT.ExpirationMinutes)).Unix(),
		"iat":     time.Now().Unix(), // Issued at
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
