package main

import (
	"net/http"
	"rsch/auth_service/app"
	"rsch/auth_service/app/config"
	"rsch/auth_service/controller"
	"rsch/auth_service/helper"
	"rsch/auth_service/repository"
	"rsch/auth_service/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	config := config.New()
	db := app.NewDB(config)
	httpClient := app.NewHttpClient()
	validate := validator.New()

	userRepository := repository.NewUser(httpClient)

	authenticationRepository := repository.NewAuthentication()
	authenticationService := service.NewAuthentication(authenticationRepository, userRepository, db, validate)
	authenticationController := controller.NewAuthentication(authenticationService)

	router := app.NewRouter(authenticationController)

	server := http.Server{
		Addr:    config.Server.Host + ":" + config.Server.Port,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
