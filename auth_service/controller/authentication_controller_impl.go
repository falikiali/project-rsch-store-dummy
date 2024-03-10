package controller

import (
	"net/http"
	"rsch/auth_service/helper"
	"rsch/auth_service/model/web"
	"rsch/auth_service/service"

	"github.com/julienschmidt/httprouter"
)

type AuthenticationControllerImpl struct {
	AuthenticationService service.AuthenticationService
}

func NewAuthenticationController(authenticationService service.AuthenticationService) AuthenticationController {
	return &AuthenticationControllerImpl{
		AuthenticationService: authenticationService,
	}
}

func (controller *AuthenticationControllerImpl) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	registerRequest := web.RegisterRequest{}
	helper.ReadFromRequestBody(r, &registerRequest)

	dataRegisterResponse := controller.AuthenticationService.Register(r.Context(), registerRequest)
	webResponse := web.WebResponse{
		StatusCode:    200,
		StatusMessage: "OK",
		Data:          dataRegisterResponse,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *AuthenticationControllerImpl) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	loginRequest := web.LoginRequest{}
	helper.ReadFromRequestBody(r, &loginRequest)

	dataLoginResponse := controller.AuthenticationService.Login(r.Context(), loginRequest)
	webResponse := web.WebResponse{
		StatusCode:    200,
		StatusMessage: "OK",
		Data:          dataLoginResponse,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *AuthenticationControllerImpl) Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	logoutRequest := web.LogoutRequest{}
	helper.ReadFromRequestBody(r, &logoutRequest)

	controller.AuthenticationService.Logout(r.Context(), logoutRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}
