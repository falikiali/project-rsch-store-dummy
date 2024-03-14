package app

import (
	"rsch/category_service/exception"
	"rsch/category_service/middleware"
	"rsch/category_service/model/domain"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController domain.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v3/category", middleware.ApiKeyMiddleware(categoryController.Create))
	router.PUT("/api/v3/category/:categoryId", middleware.ApiKeyMiddleware(categoryController.Update))
	router.GET("/api/v3/category", categoryController.FindAll)

	router.PanicHandler = exception.ErrorHandler

	return router
}
