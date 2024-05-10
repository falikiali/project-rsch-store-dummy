package domain

import (
	"context"
	"database/sql"
)

type ProductImage struct {
	Id   string
	Data []byte
}

type ProductImageRepository interface {
	Create(ctx context.Context, tx *sql.Tx, productImage ProductImage) ProductImage
	Update(ctx context.Context, tx *sql.Tx, productImage ProductImage) ProductImage
	FindById(ctx context.Context, tx *sql.Tx, idProductImage string) (ProductImage, error)
}
