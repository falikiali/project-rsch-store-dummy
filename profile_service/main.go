package main

import (
	"net/http"
	"rsch/profile_service/app"
	"rsch/profile_service/app/config"
	"rsch/profile_service/controller"
	"rsch/profile_service/helper"
	"rsch/profile_service/repository"
	"rsch/profile_service/service"

	"github.com/go-playground/validator/v10"
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

	router := app.NewRouter(userController)

	server := http.Server{
		Addr:    config.Server.Host + ":" + config.Server.Port,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
