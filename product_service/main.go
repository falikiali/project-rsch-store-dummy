package main

import (
	"net/http"
	"rsch/product_service/app"
	"rsch/product_service/app/config"
	"rsch/product_service/controller"
	"rsch/product_service/helper"
	"rsch/product_service/repository"
	"rsch/product_service/service"
)

func main() {
	config := config.New()
	db := app.NewDB(config)

	productRepository := repository.NewProduct()
	productSizeRepository := repository.NewProductSize()
	productImageRepository := repository.NewProductImage()

	productService := service.NewProduct(productRepository, productSizeRepository, productImageRepository, db)
	productController := controller.NewProduct(productService)
	router := app.NewRouter(productController)

	server := http.Server{
		Addr:    config.Server.Host + ":" + config.Server.Port,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
