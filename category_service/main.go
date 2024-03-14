package main

import (
	"net/http"
	"rsch/category_service/app"
	"rsch/category_service/app/config"
	"rsch/category_service/controller"
	"rsch/category_service/helper"
	"rsch/category_service/repository"
	"rsch/category_service/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	config := config.NewConfig()
	db := app.NewDB(config)
	validate := validator.New()

	categoryRepository := repository.NewCategory()
	categoryService := service.NewCategory(categoryRepository, db, validate)
	categoryController := controller.NewCategory(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    config.Server.Host + ":" + config.Server.Port,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
