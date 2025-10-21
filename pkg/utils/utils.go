package utils

import (
	"encoding/hex"
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

var decodedSecretKey []byte

func InitializeJWT() error {
	cfg := config.GetConfig()
	if cfg == nil {
		return fmt.Errorf("config not loaded")
	}

	secretKeyHex := cfg.JWT.SecretKey
	var err error
	decodedSecretKey, err = hex.DecodeString(secretKeyHex)
	if err != nil {
		return fmt.Errorf("invalid JWT secret key format: %w", err)
	}

	if len(decodedSecretKey) < 32 {
		return fmt.Errorf("JWT secret key must be at least 256 bits (32 bytes)")
	}

	return nil
}

func GenerateJWTToken(memberID string) (string, error) {
	cfg := config.GetConfig()
	if cfg == nil {
		return "", fmt.Errorf("config not loaded")
	}

	claims := jwt.MapClaims{
		"member_id": memberID,
		"exp":       time.Now().Add(time.Minute * time.Duration(cfg.JWT.ExpirationMinutes)).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(decodedSecretKey)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return decodedSecretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("token parsing failed: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims or signature")
	}

	return claims, nil
}
