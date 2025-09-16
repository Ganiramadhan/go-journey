package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func ttlFromEnv(key string, def time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return def
}

func GenerateTokenPair(userID string) (TokenPair, error) {
	secret := os.Getenv("JWT_SECRET")
	accessTTL := ttlFromEnv("ACCESS_TOKEN_TTL", 15*time.Minute)
	refreshTTL := ttlFromEnv("REFRESH_TOKEN_TTL", 7*24*time.Hour)

	now := time.Now()

	access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"type": "access",
		"exp":  now.Add(accessTTL).Unix(),
		"iat":  now.Unix(),
	})

	// Refresh token
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"type": "refresh",
		"exp":  now.Add(refreshTTL).Unix(),
		"iat":  now.Unix(),
	})

	// Sign token
	accessStr, err := access.SignedString([]byte(secret))
	if err != nil {
		return TokenPair{}, err
	}
	refreshStr, err := refresh.SignedString([]byte(secret))
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{AccessToken: accessStr, RefreshToken: refreshStr}, nil
}

func ParseToken(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {
	secret := os.Getenv("JWT_SECRET")
	t, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return t, nil, err
	}

	return t, claims, nil
}
