package web_test

import (
	"rsch/auth_service/model/web"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestRegisterRequest(t *testing.T) {
	validate := validator.New()
	registerRequest := web.RegisterRequest{
		Email:           "alifaliki@gmail.com",
		Password:        "12345678",
		ConfirmPassword: "12345678",
	}

	err := validate.Struct(registerRequest)
	assert.Nil(t, err)
}
