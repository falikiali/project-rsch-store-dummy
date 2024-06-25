package app

import (
	"rsch/product_service/exception"
	"rsch/product_service/middleware"
	"rsch/product_service/model/domain"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller domain.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.PanicHandler = exception.ErrorHandler
	router.POST("/api/v3/product", middleware.ApiKeyMiddleware(controller.Create))
	router.GET("/api/v3/product", controller.FindProducts)
	router.GET("/api/v3/product/:idProduct", controller.FindProductById)
	router.GET("/image/:idImage", controller.Image)

	router.PanicHandler = exception.ErrorHandler

	return router
}
