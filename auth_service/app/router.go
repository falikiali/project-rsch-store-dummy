package app

import (
	"rsch/auth_service/exception"
	"rsch/auth_service/middleware"
	"rsch/auth_service/model/domain"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(authenticationController domain.AuthenticationController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v3/authentication", middleware.BearerTokenMiddleware(authenticationController.ValidateToken))
	router.POST("/api/v3/authentication/login", authenticationController.Login)
	router.POST("/api/v3/authentication/register", authenticationController.Register)
	router.DELETE("/api/v3/authentication/logout", authenticationController.Logout)

	router.PanicHandler = exception.ErrorHandler

	return router
}
