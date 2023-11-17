package utils

import (
	"errors"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtSecret     = []byte(os.Getenv("SECRET_KEY")) // Replace with your actual secret key
	revokedTokens = make(map[string]bool)
	mutex         sync.Mutex
)

func InvalidateToken(tokenString string) {
	mutex.Lock()
	defer mutex.Unlock()

	revokedTokens[tokenString] = true
}

// ParseToken parses a JWT token and returns the claims if the token is valid and not revoked
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	if revokedTokens[tokenString] {
		return nil, errors.New("Token has been revoked")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid token")
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func StringToUint(strValue string) (uint, error) {
	uintValue, err := strconv.ParseUint(strValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(uintValue), nil
}
