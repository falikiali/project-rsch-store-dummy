package app

import (
	"rsch/cart_service/exception"
	"rsch/cart_service/middleware"
	"rsch/cart_service/model/domain"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(cartController domain.CartController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v3/cart", middleware.BearerTokenMiddleware(cartController.AddProductToCart))
	router.PUT("/api/v3/cart", middleware.BearerTokenMiddleware(cartController.UpdateProductInCart))
	router.DELETE("/api/v3/cart", middleware.BearerTokenMiddleware(cartController.DeleteProductInCart))
	router.GET("/api/v3/cart", middleware.BearerTokenMiddleware(cartController.FindProductsInCartByIdUser))
	router.PUT("/api/v3/cart/selected", middleware.BearerTokenMiddleware(cartController.UpdateSelectedProductInCart))

	router.PanicHandler = exception.ErrorHandler

	return router
}
