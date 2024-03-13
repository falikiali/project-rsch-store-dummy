package main

import (
	"net/http"
	"rsch/profile_service/app"
	"rsch/profile_service/app/config"
	"rsch/profile_service/controller"
	"rsch/profile_service/exception"
	"rsch/profile_service/helper"
	"rsch/profile_service/middleware"
	"rsch/profile_service/repository"
	"rsch/profile_service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	client := app.NewHttpClient()
	config := config.New()
	validate := validator.New()
	db := app.NewDB(config)

	authenticationRepository := repository.NewAuthentication(client)

	userRepository := repository.NewUser()
	userService := service.NewUser(userRepository, authenticationRepository, db, validate)
	userController := controller.NewUser(userService)

	router := httprouter.New()

	router.POST("/api/v3/user", userController.Create)
	router.GET("/api/v3/user/validate", userController.FindUserByEmailAndPassword)
	router.GET("/api/v3/user", middleware.BearerTokenMiddleware(userController.FindUserById))
	router.PUT("/api/v3/user/password", middleware.BearerTokenMiddleware(userController.ChangePassword))
	router.PUT("/api/v3/user/profile", middleware.BearerTokenMiddleware(userController.UpdateProfile))
	router.PUT("/api/v3/user/phone-number", middleware.BearerTokenMiddleware(userController.UpdatePhoneNumber))

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    config.Server.Host + ":" + config.Server.Port,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
