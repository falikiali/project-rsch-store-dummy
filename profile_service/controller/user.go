package controller

import (
	"net/http"
	"rsch/profile_service/helper"
	"rsch/profile_service/model/domain"
	"rsch/profile_service/model/web"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	UserService domain.UserService
}

func NewUser(userService domain.UserService) domain.UserController {
	return &User{
		UserService: userService,
	}
}

func (controller *User) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	createUserRequest := web.CreateUserRequest{}
	helper.ReadFromRequestBody(r, &createUserRequest)

	data := controller.UserService.Create(r.Context(), createUserRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          data,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *User) ChangePassword(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	changePasswordRequest := web.ChangePasswordRequest{}
	helper.ReadFromRequestBody(r, &changePasswordRequest)

	controller.UserService.ChangePassword(r.Context(), changePasswordRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *User) UpdateProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	updateProfileRequest := web.UpdateProfileRequest{}
	helper.ReadFromRequestBody(r, &updateProfileRequest)

	data := controller.UserService.UpdateProfile(r.Context(), updateProfileRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          data,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *User) UpdatePhoneNumber(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	updatePhoneNumberRequest := web.UpdatePhoneNumberRequest{}
	helper.ReadFromRequestBody(r, &updatePhoneNumberRequest)

	data := controller.UserService.UpdatePhoneNumber(r.Context(), updatePhoneNumberRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          data,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}
