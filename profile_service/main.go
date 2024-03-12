package main

import (
	"net/http"
	"rsch/profile_service/app"
	"rsch/profile_service/app/config"
	"rsch/profile_service/controller"
	"rsch/profile_service/exception"
	"rsch/profile_service/helper"
	"rsch/profile_service/repository"
	"rsch/profile_service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	config := config.New()
	validate := validator.New()
	db := app.NewDB(config)

	userRepository := repository.NewUser()
	userService := service.NewUser(userRepository, db, validate)
	userController := controller.NewUser(userService)

	router := httprouter.New()

	router.POST("/api/v3/user", userController.Create)
	router.PUT("/api/v3/user/password", userController.ChangePassword)
	router.PUT("/api/v3/user/profile", userController.UpdateProfile)
	router.PUT("/api/v3/user/phone-number", userController.UpdatePhoneNumber)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    config.Server.Host + ":" + config.Server.Port,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
