package domain

import (
	"context"
	"database/sql"
	"net/http"
	"rsch/category_service/model/web"

	"github.com/julienschmidt/httprouter"
)

type Category struct {
	Id   int16
	Name string
}

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category Category) Category
	Update(ctx context.Context, tx *sql.Tx, category Category) (Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []Category
	FindById(ctx context.Context, tx *sql.Tx, id int16) error
	FindNameIsExist(ctx context.Context, tx *sql.Tx, name string) (Category, error)
}

type CategoryService interface {
	Create(ctx context.Context, request web.CreateCategoryRequest) web.CreateCategoryResponse
	Update(ctx context.Context, request web.UpdateCategoryRequest) web.UpdateCategoryResponse
	FindAll(ctx context.Context) []web.FindAllCategoryResponse
}

type CategoryController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
