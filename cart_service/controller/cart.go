package controller

import (
	"net/http"
	"rsch/cart_service/helper"
	"rsch/cart_service/model/domain"
	"rsch/cart_service/model/web"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Cart struct {
	CartService domain.CartService
}

func NewCart(cartService domain.CartService) domain.CartController {
	return &Cart{
		CartService: cartService,
	}
}

func (controller *Cart) AddProductToCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	addProductRequest := web.AddProductToCartRequest{}
	helper.ReadFromRequestBody(r, &addProductRequest)

	s := controller.CartService.AddProductToCart(r.Context(), token, addProductRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          s,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *Cart) UpdateProductInCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	updateProductRequests := []web.UpdateProductInCartRequest{}
	helper.ReadFromRequestBody(r, &updateProductRequests)

	controller.CartService.UpdateProductInCart(r.Context(), token, updateProductRequests)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *Cart) UpdateSelectedProductInCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	updateProductRequest := web.UpdateSelectedProductInCartRequest{}
	helper.ReadFromRequestBody(r, &updateProductRequest)

	controller.CartService.UpdateSelectedProductInCart(r.Context(), token, updateProductRequest)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *Cart) DeleteProductInCart(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	deleteProductRequests := []web.DeleteProductInCartRequest{}
	helper.ReadFromRequestBody(r, &deleteProductRequests)

	controller.CartService.DeleteProductInCart(r.Context(), token, deleteProductRequests)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}

func (controller *Cart) FindProductsInCartByIdUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authorizationHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	dataResponse := controller.CartService.FindProductsInCartByIdUser(r.Context(), token)
	webResponse := web.WebResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
		Data:          dataResponse,
	}

	helper.WriteToResponseBody(w, http.StatusOK, webResponse)
}
