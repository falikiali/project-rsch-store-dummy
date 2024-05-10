package service

import (
	"context"
	"database/sql"
	"rsch/product_service/exception"
	"rsch/product_service/helper"
	"rsch/product_service/model/domain"
	"rsch/product_service/model/web"

	"github.com/google/uuid"
)

type Product struct {
	ProductRepository      domain.ProductRepository
	ProductSizeRepository  domain.ProductSizeRepository
	ProductImageRepository domain.ProductImageRepository
	DB                     *sql.DB
}

func NewProduct(
	productRepository domain.ProductRepository,
	productSizeRepository domain.ProductSizeRepository,
	productImageRepository domain.ProductImageRepository,
	db *sql.DB,
) domain.ProductService {
	return &Product{
		ProductRepository:      productRepository,
		ProductSizeRepository:  productSizeRepository,
		ProductImageRepository: productImageRepository,
		DB:                     db,
	}
}

func (service *Product) Create(ctx context.Context, product domain.Product, productImage domain.ProductImage) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	newProductImage := domain.ProductImage{
		Id:   productImage.Id,
		Data: productImage.Data,
	}
	newProductImage = service.ProductImageRepository.Create(ctx, tx, newProductImage)

	newProduct := domain.Product{
		Id:          uuid.NewString(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		Purpose:     product.Purpose,
		Image:       newProductImage.Id,
		DetailSize:  product.DetailSize,
	}
	_ = service.ProductRepository.Create(ctx, tx, newProduct)

	productSizes := []domain.ProductSize{}
	for _, detailSize := range newProduct.DetailSize {
		detailSize.Id = uuid.NewString()
		detailSize.IdProduct = newProduct.Id

		productSizes = append(productSizes, detailSize)
	}
	productSizes = service.ProductSizeRepository.Create(ctx, tx, productSizes)

	countStock := 0
	detailSizeResponses := []web.DetailSizeResponse{}
	for _, productSize := range productSizes {
		countStock += productSize.Stock

		detailSizeResponse := web.DetailSizeResponse{
			Size:  productSize.Size,
			Stock: productSize.Stock,
		}
		detailSizeResponses = append(detailSizeResponses, detailSizeResponse)
	}

	return web.ProductResponse{
		Id:          newProduct.Id,
		Name:        newProduct.Name,
		Price:       newProduct.Price,
		Description: newProduct.Description,
		Purpose:     newProduct.Purpose,
		Category:    newProduct.Category,
		Stock:       countStock,
		Image:       "/image/" + newProductImage.Id,
		DetailSize:  detailSizeResponses,
	}
}

func (service *Product) Image(ctx context.Context, idImage string) []byte {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	productImage, err := service.ProductImageRepository.FindById(ctx, tx, idImage)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return productImage.Data
}

func (service *Product) FindProducts(ctx context.Context, page int, filters map[string]interface{}) ([]web.ProductResponse, web.PaginationResponse) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindProducts(ctx, tx, page, filters)
	if len(products) == 0 {
		panic(exception.NewNotFoundError("data is empty"))
	}

	productResponses := []web.ProductResponse{}
	for _, product := range products {
		productResponse := web.ProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Purpose:     product.Purpose,
			Category:    product.Category,
			Stock:       product.Stock,
			Image:       "/image/" + product.Image,
		}
		productResponses = append(productResponses, productResponse)
	}

	totalData, totalPage := service.ProductRepository.CountProducts(ctx, tx, filters)
	paginationResponse := web.PaginationResponse{
		Page:      page,
		TotalPage: totalPage,
		TotalData: totalData,
	}

	return productResponses, paginationResponse
}

func (service *Product) FindProductById(ctx context.Context, idProduct string) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindProductById(ctx, tx, idProduct)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	detailSizeResponses := []web.DetailSizeResponse{}
	for _, ps := range product.DetailSize {
		detailSizeResponses = append(detailSizeResponses, web.DetailSizeResponse{
			Id:    ps.Id,
			Size:  ps.Size,
			Stock: ps.Stock,
		})
	}

	productResponse := web.ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Purpose:     product.Purpose,
		Category:    product.Category,
		Image:       "/image/" + product.Image,
		Stock:       product.Stock,
		DetailSize:  detailSizeResponses,
	}

	return productResponse
}
