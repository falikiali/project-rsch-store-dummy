package domain

import (
	"context"
	"database/sql"
	"net/http"
	"rsch/product_service/model/web"

	"github.com/julienschmidt/httprouter"
)

type Product struct {
	Id          string
	Name        string
	Price       int64
	Description string
	Category    int16
	Purpose     string
	Image       string
	Stock       int
	DetailSize  []ProductSize
}

type ProductRepository interface {
	Create(ctx context.Context, tx *sql.Tx, product Product) Product
	Update(ctx context.Context, tx *sql.Tx, product Product) Product
	FindProducts(ctx context.Context, tx *sql.Tx, page int, filters map[string]interface{}) []Product
	CountProducts(ctx context.Context, tx *sql.Tx, filters map[string]interface{}) (int, int)
	FindProductById(ctx context.Context, tx *sql.Tx, idProduct string) (Product, error)
}

type ProductService interface {
	Create(ctx context.Context, product Product, productImage ProductImage) web.ProductResponse
	// Update(ctx context.Context) web.ProductResponse
	FindProducts(ctx context.Context, page int, filters map[string]interface{}) ([]web.ProductResponse, web.PaginationResponse)
	FindProductById(ctx context.Context, idProduct string) web.ProductResponse
	Image(Ctx context.Context, idImage string) []byte
}

type ProductController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Image(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindProducts(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindProductById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
