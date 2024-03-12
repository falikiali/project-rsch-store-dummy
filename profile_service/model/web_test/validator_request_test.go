package webtest

import (
	"rsch/profile_service/model/web"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestUpdateProfileRequest(t *testing.T) {
	req := web.UpdateProfileRequest{
		Fullname: "Ali Faliki",
		Username: "falikiali",
	}

	validate := validator.New()
	err := validate.Struct(req)

	assert.Nil(t, err)
}

func TestUpdatePhoneNumberRequest(t *testing.T) {
	req := web.UpdatePhoneNumberRequest{
		PhoneNumber: "2681282085511",
	}

	validate := validator.New()
	err := validate.Struct(req)

	assert.Nil(t, err)
}

func TestChangePasswordRequest(t *testing.T) {
	req := web.ChangePasswordRequest{
		OldPassword:        "password12",
		NewPassword:        "password123",
		ConfirmNewPassword: "password123",
	}

	validate := validator.New()
	err := validate.Struct(req)

	assert.Nil(t, err)
}
