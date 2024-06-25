package service

import (
	"context"
	"database/sql"
	"rsch/cart_service/exception"
	"rsch/cart_service/helper"
	"rsch/cart_service/model/domain"
	"rsch/cart_service/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Cart struct {
	CartRepository           domain.CartRepository
	AuthenticationRepository domain.AuthenticationRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewCart(cartRepository domain.CartRepository, authenticationRepository domain.AuthenticationRepository, DB *sql.DB, validate *validator.Validate) domain.CartService {
	return &Cart{
		CartRepository:           cartRepository,
		AuthenticationRepository: authenticationRepository,
		DB:                       DB,
		Validate:                 validate,
	}
}

func (service *Cart) AddProductToCart(ctx context.Context, token string, request web.AddProductToCartRequest) string {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	idUser, err := service.AuthenticationRepository.ValidateToken(ctx, token)
	if err != nil {
		if err.Error() == "" {
			panic(err)
		} else {
			panic(exception.NewUnauthorizedError("bearer token is invalid"))
		}
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	cart := domain.Cart{
		Id:            uuid.NewString(),
		IdUser:        idUser,
		IdProduct:     request.IdProduct,
		IdProductSize: request.IdProductSize,
		Quantity:      request.Quantity,
	}

	newCart, err := service.CartRepository.FindProductInCartIsExist(ctx, tx, cart)
	if err == nil {
		tx, err := service.DB.Begin()
		helper.PanicIfError(err)
		defer helper.CommitOrRollback(tx)

		err = service.CartRepository.Update(ctx, tx, newCart.Id, newCart.Quantity+1)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		return newCart.Id
	}

	s := service.CartRepository.Create(ctx, tx, cart)
	return s
}

func (service *Cart) UpdateProductInCart(ctx context.Context, token string, requests []web.UpdateProductInCartRequest) {
	_, err := service.AuthenticationRepository.ValidateToken(ctx, token)
	if err != nil {
		if err.Error() == "" {
			panic(err)
		} else {
			panic(exception.NewUnauthorizedError("bearer token is invalid"))
		}
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	for _, request := range requests {
		err := service.Validate.Struct(request)
		helper.PanicIfError(err)

		err = service.CartRepository.Update(ctx, tx, request.Id, request.Quantity)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
	}
}

func (service *Cart) UpdateSelectedProductInCart(ctx context.Context, token string, request web.UpdateSelectedProductInCartRequest) {
	_, err := service.AuthenticationRepository.ValidateToken(ctx, token)
	if err != nil {
		if err.Error() == "" {
			panic(err)
		} else {
			panic(exception.NewUnauthorizedError("bearer token is invalid"))
		}
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Validate.Struct(request)
	helper.PanicIfError(err)

	var isSelected int
	if request.IsSelected {
		isSelected = 1
	} else {
		isSelected = 0
	}

	err = service.CartRepository.UpdateSelected(ctx, tx, request.Id, isSelected)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
}

func (service *Cart) DeleteProductInCart(ctx context.Context, token string, requests []web.DeleteProductInCartRequest) {
	_, err := service.AuthenticationRepository.ValidateToken(ctx, token)
	if err != nil {
		if err.Error() == "" {
			panic(err)
		} else {
			panic(exception.NewUnauthorizedError("bearer token is invalid"))
		}
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	for _, request := range requests {
		err := service.Validate.Struct(request)
		helper.PanicIfError(err)

		err = service.CartRepository.Delete(ctx, tx, request.Id)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
	}
}

func (service *Cart) FindProductsInCartByIdUser(ctx context.Context, token string) []web.FindProductsInCartByIdUserResponse {
	idUser, err := service.AuthenticationRepository.ValidateToken(ctx, token)
	if err != nil {
		if err.Error() == "" {
			panic(err)
		} else {
			panic(exception.NewUnauthorizedError("bearer token is invalid"))
		}
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.CartRepository.FindProductsInCartByIdUser(ctx, tx, idUser)
	productResponses := []web.FindProductsInCartByIdUserResponse{}

	for _, product := range products {
		var isSelected bool

		if product.IsSelected == 0 {
			isSelected = false
		} else {
			isSelected = true
		}

		productResponse := web.FindProductsInCartByIdUserResponse{
			Id:            product.Id,
			IdProduct:     product.IdProduct,
			IdProductSize: product.IdProductSize,
			IsSelected:    isSelected,
			ProductName:   product.ProductName,
			ProductImage:  "http://192.168.1.9:3003/image/" + product.ProductImage,
			ProductSize:   product.ProductSize,
			ProductStock:  product.ProductStock,
			Quantity:      product.Quantity,
			TotalPrice:    product.TotalPrice,
		}

		productResponses = append(productResponses, productResponse)
	}

	return productResponses
}
