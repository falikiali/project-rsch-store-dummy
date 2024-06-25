package main

import (
	"net/http"
	"rsch/cart_service/app"
	"rsch/cart_service/app/config"
	"rsch/cart_service/controller"
	"rsch/cart_service/helper"
	"rsch/cart_service/repository"
	"rsch/cart_service/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	client := app.NewHttpClient()
	config := config.New()
	db := app.NewDB(config)
	validate := validator.New()

	authenticationRepository := repository.NewAuthentication(client)
	cartRepository := repository.NewCart()
	cartService := service.NewCart(cartRepository, authenticationRepository, db, validate)
	cartController := controller.NewCart(cartService)

	router := app.NewRouter(cartController)

	server := http.Server{
		Addr:    config.Server.Host + ":" + config.Server.Port,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
