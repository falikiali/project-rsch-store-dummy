package app

import (
	"rsch/auth_service/helper"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateNewJwt(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         id,
		"created_at": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte("ouval-research-est-1997"))
	helper.PanicIfError(err)
	return tokenString
}
