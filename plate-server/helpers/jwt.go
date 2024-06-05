package helpers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var err error
var tokenExpirationDuration = time.Hour * 8

func GenerateToken(id uuid.UUID, nama string) string {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	claims := jwt.MapClaims{
		"id":   id,
		"nama": nama,
		"time": time.Now().Add(tokenExpirationDuration).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("SECRET")

	signToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return ""
	}

	return signToken
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	errRes := errors.New("token is invalid")
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer ")

	if !bearer {
		return nil, errRes
	}

	secretKey := os.Getenv("SECRET")

	tokenString := strings.TrimPrefix(headerToken, "Bearer ")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := int64(claims["time"].(float64))
		if time.Unix(exp, 0).Before(time.Now()) {
			return nil, errors.New("token has expired")
		}
		return claims, nil
	}

	return nil, errRes
}
