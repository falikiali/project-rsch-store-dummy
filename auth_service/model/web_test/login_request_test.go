package web_test

import (
	"rsch/auth_service/model/web"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestLoginRequest(t *testing.T) {
	validate := validator.New()
	loginRequest := web.LoginRequest{
		Email:    "ali@gmail.com",
		Password: "asdasdasd",
	}

	err := validate.Struct(loginRequest)
	assert.Nil(t, err)
}
