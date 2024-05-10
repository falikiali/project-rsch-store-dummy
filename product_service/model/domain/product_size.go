package domain

import (
	"context"
	"database/sql"
)

type ProductSize struct {
	Id        string
	IdProduct string
	Size      string
	Stock     int
}

type ProductSizeRepository interface {
	Create(ctx context.Context, tx *sql.Tx, productSizes []ProductSize) []ProductSize
	Update(ctx context.Context, tx *sql.Tx, productSize ProductSize) ProductSize
	FindByIdProduct(ctx context.Context, tx *sql.Tx, idProduct string) []ProductSize
}
