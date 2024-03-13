package controller

import (
	"net/http"
	"rsch/auth_service/helper"
	"rsch/auth_service/model/domain"
	"rsch/auth_service/model/web"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Authentication struct {
	AuthenticationService domain.AuthenticationService
}

func NewAuthentication(authenticationService domain.AuthenticationService) domain.AuthenticationController {
	return &Authentication{
		AuthenticationService: authenticationService,
	}
}

func (controller *Authentication) ValidateToken(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	dataValidateToken := controller.AuthenticationService.ValidateToken(r.Context(), token)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          dataValidateToken,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *Authentication) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	registerRequest := web.RegisterRequest{}
	helper.ReadFromRequestBody(r, &registerRequest)

	dataRegisterResponse := controller.AuthenticationService.Register(r.Context(), registerRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          dataRegisterResponse,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *Authentication) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	loginRequest := web.LoginRequest{}
	helper.ReadFromRequestBody(r, &loginRequest)

	dataLoginResponse := controller.AuthenticationService.Login(r.Context(), loginRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          dataLoginResponse,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *Authentication) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	logoutRequest := web.LogoutRequest{}
	helper.ReadFromRequestBody(r, &logoutRequest)

	controller.AuthenticationService.Logout(r.Context(), logoutRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}
