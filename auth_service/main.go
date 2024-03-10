package main

import (
	"net/http"
	"os"
	"rsch/auth_service/app"
	"rsch/auth_service/controller"
	"rsch/auth_service/exception"
	"rsch/auth_service/helper"
	"rsch/auth_service/repository"
	"rsch/auth_service/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)

	db := app.NewDB()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	authenticationRepository := repository.NewAuthenticationRepository()
	authenticationService := service.NewAuthenticationService(authenticationRepository, userRepository, db, validate)
	authenticationController := controller.NewAuthenticationController(authenticationService)

	router := httprouter.New()

	router.POST("/api/v3/authentication/login", authenticationController.Login)
	router.POST("/api/v3/authentication/register", authenticationController.Register)
	router.DELETE("/api/v3/authentication/logout", authenticationController.Logout)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:" + os.Getenv("RUNNING_PORT"),
		Handler: router,
	}

	err = server.ListenAndServe()
	helper.PanicIfError(err)

	//Kurang API untuk Logout
}
