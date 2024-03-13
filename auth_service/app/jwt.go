package app

import (
	"errors"
	"os"
	"rsch/auth_service/helper"
	"rsch/auth_service/model/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type JWTClaim struct {
	Id        string `json:"id"`
	CreatedAt int64  `json:"created_at"`
}

func GenerateNewJwt(id string) string {
	err := godotenv.Load()
	helper.PanicIfError(err)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         id,
		"created_at": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	helper.PanicIfError(err)
	return tokenString
}

func ParseJWT(tokenString string) (domain.User, error) {
	err := godotenv.Load()
	helper.PanicIfError(err)

	user := domain.User{}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	helper.PanicIfError(err)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user.Id = claims["id"].(string)
		return user, nil
	}

	return user, errors.New("token is invalid")
}
