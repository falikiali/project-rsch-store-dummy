package app

import (
	"rsch/profile_service/exception"
	"rsch/profile_service/middleware"
	"rsch/profile_service/model/domain"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController domain.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v3/user", userController.Create)
	router.GET("/api/v3/user/validate", userController.FindUserByEmailAndPassword)
	router.GET("/api/v3/user", middleware.BearerTokenMiddleware(userController.FindUserById))
	router.PUT("/api/v3/user/password", middleware.BearerTokenMiddleware(userController.ChangePassword))
	router.PUT("/api/v3/user/profile", middleware.BearerTokenMiddleware(userController.UpdateProfile))
	router.PUT("/api/v3/user/phone-number", middleware.BearerTokenMiddleware(userController.UpdatePhoneNumber))

	router.PanicHandler = exception.ErrorHandler

	return router
}
