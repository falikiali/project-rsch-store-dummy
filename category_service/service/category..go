package service

import (
	"context"
	"database/sql"
	"rsch/category_service/exception"
	"rsch/category_service/helper"
	"rsch/category_service/model/domain"
	"rsch/category_service/model/web"

	"github.com/go-playground/validator/v10"
)

type Category struct {
	CategoryRepository domain.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategory(categoryRepository domain.CategoryRepository, db *sql.DB, validate *validator.Validate) domain.CategoryService {
	return &Category{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *Category) Create(ctx context.Context, request web.CreateCategoryRequest) web.CreateCategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category, err = service.CategoryRepository.FindNameIsExist(ctx, tx, category.Name)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	category = service.CategoryRepository.Create(ctx, tx, category)

	return web.CreateCategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func (service *Category) Update(ctx context.Context, request web.UpdateCategoryRequest) web.UpdateCategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Id:   request.Id,
		Name: request.Name,
	}

	err = service.CategoryRepository.FindById(ctx, tx, category.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category, err = service.CategoryRepository.FindNameIsExist(ctx, tx, category.Name)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	category.Id = request.Id

	category, err = service.CategoryRepository.Update(ctx, tx, category)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.UpdateCategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func (service *Category) FindAll(ctx context.Context) []web.FindAllCategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	categoryResponses := []web.FindAllCategoryResponse{}
	for _, category := range categories {
		categoryResponses = append(categoryResponses, web.FindAllCategoryResponse(category))
	}
	return categoryResponses
}
