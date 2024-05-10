package repository

import (
	"context"
	"database/sql"
	"rsch/product_service/helper"
	"rsch/product_service/model/domain"
)

type ProductSize struct{}

func NewProductSize() domain.ProductSizeRepository {
	return &ProductSize{}
}

func (repository *ProductSize) Create(ctx context.Context, tx *sql.Tx, productSizes []domain.ProductSize) []domain.ProductSize {
	SQL := "INSERT INTO product_size (id, id_product, size, stock) VALUES(?, ?, ?, ?)"

	statment, err := tx.PrepareContext(ctx, SQL)
	helper.PanicIfError(err)
	defer statment.Close()

	for _, productSize := range productSizes {
		_, err := statment.ExecContext(ctx, productSize.Id, productSize.IdProduct, productSize.Size, productSize.Stock)
		helper.PanicIfError(err)
	}

	return productSizes
}

func (repository *ProductSize) Update(ctx context.Context, tx *sql.Tx, productSize domain.ProductSize) domain.ProductSize {
	SQL := "UPDATE product_size SET size = ?, stock = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, productSize.Size, productSize.Stock, productSize.Id)
	helper.PanicIfError(err)

	return productSize
}

func (repository *ProductSize) FindByIdProduct(ctx context.Context, tx *sql.Tx, idProduct string) []domain.ProductSize {
	SQL := "SELECT id, id_product, size, stock FROM product_size WHERE id_product = ?"
	r, err := tx.QueryContext(ctx, SQL, idProduct)
	helper.PanicIfError(err)

	productSizes := []domain.ProductSize{}
	for r.Next() {
		productSize := domain.ProductSize{}
		err := r.Scan(&productSize.Id, &productSize.IdProduct, &productSize.Size, &productSize.Stock)
		helper.PanicIfError(err)

		productSizes = append(productSizes, productSize)
	}

	return productSizes
}
