package controller

import (
	"net/http"
	"rsch/profile_service/exception"
	"rsch/profile_service/helper"
	"rsch/profile_service/model/domain"
	"rsch/profile_service/model/web"
	"strings"

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

func (controller *User) FindUserById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	data := controller.UserService.FindUserById(r.Context(), token)

	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          data,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *User) FindUserByEmailAndPassword(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")

	if email == "" || password == "" {
		panic(exception.NewBadRequestError("email or password is cannot be empty"))
	}

	data := controller.UserService.FindUserByEmailAndPassword(r.Context(), email, password)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          data,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *User) ChangePassword(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	changePasswordRequest := web.ChangePasswordRequest{}
	helper.ReadFromRequestBody(r, &changePasswordRequest)

	controller.UserService.ChangePassword(r.Context(), token, changePasswordRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *User) UpdateProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	updateProfileRequest := web.UpdateProfileRequest{}
	helper.ReadFromRequestBody(r, &updateProfileRequest)

	data := controller.UserService.UpdateProfile(r.Context(), token, updateProfileRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          data,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *User) UpdatePhoneNumber(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	updatePhoneNumberRequest := web.UpdatePhoneNumberRequest{}
	helper.ReadFromRequestBody(r, &updatePhoneNumberRequest)

	data := controller.UserService.UpdatePhoneNumber(r.Context(), token, updatePhoneNumberRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          data,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}
