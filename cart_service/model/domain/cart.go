package domain

import (
	"context"
	"database/sql"
	"net/http"
	"rsch/cart_service/model/web"

	"github.com/julienschmidt/httprouter"
)

type Cart struct {
	Id            string
	IdUser        string
	IdProduct     string
	IdProductSize string
	IsSelected    int
	ProductName   string
	ProductImage  string
	ProductSize   string
	ProductStock  int
	Quantity      int
	TotalPrice    int
}

type CartRepository interface {
	Create(ctx context.Context, tx *sql.Tx, cart Cart) string
	Update(ctx context.Context, tx *sql.Tx, idCart string, qty int) error
	UpdateSelected(ctx context.Context, tx *sql.Tx, idCart string, isSelected int) error
	Delete(ctx context.Context, tx *sql.Tx, idCart string) error
	FindProductsInCartByIdUser(ctx context.Context, tx *sql.Tx, idUser string) []Cart
	FindProductInCartIsExist(ctx context.Context, tx *sql.Tx, cart Cart) (Cart, error)
}

type CartService interface {
	AddProductToCart(ctx context.Context, token string, request web.AddProductToCartRequest) string
	UpdateProductInCart(ctx context.Context, token string, requests []web.UpdateProductInCartRequest)
	UpdateSelectedProductInCart(ctx context.Context, token string, request web.UpdateSelectedProductInCartRequest)
	DeleteProductInCart(ctx context.Context, token string, requests []web.DeleteProductInCartRequest)
	FindProductsInCartByIdUser(ctx context.Context, token string) []web.FindProductsInCartByIdUserResponse
}

type CartController interface {
	AddProductToCart(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateProductInCart(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateSelectedProductInCart(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	DeleteProductInCart(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindProductsInCartByIdUser(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
