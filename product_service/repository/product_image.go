package repository

import (
	"context"
	"database/sql"
	"errors"
	"rsch/product_service/helper"
	"rsch/product_service/model/domain"
)

type ProductImage struct{}

func NewProductImage() domain.ProductImageRepository {
	return &ProductImage{}
}

func (repository *ProductImage) Create(ctx context.Context, tx *sql.Tx, productImage domain.ProductImage) domain.ProductImage {
	SQL := "INSERT INTO product_image (id, data) VALUES(?, ?)"
	_, err := tx.ExecContext(ctx, SQL, productImage.Id, productImage.Data)
	helper.PanicIfError(err)

	return productImage
}

func (repository *ProductImage) Update(ctx context.Context, tx *sql.Tx, productImage domain.ProductImage) domain.ProductImage {
	SQL := "UPDATE product_image SET data = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, productImage.Data, productImage.Id)
	helper.PanicIfError(err)

	return productImage
}

func (repository *ProductImage) FindById(ctx context.Context, tx *sql.Tx, idProductImage string) (domain.ProductImage, error) {
	SQL := "SELECT id, data FROM product_image WHERE id = ?"
	r, err := tx.QueryContext(ctx, SQL, idProductImage)
	helper.PanicIfError(err)

	productImage := domain.ProductImage{}
	if r.Next() {
		err := r.Scan(&productImage.Id, &productImage.Data)
		helper.PanicIfError(err)

		return productImage, nil
	}

	return productImage, errors.New("product image not found")
}
