package repository

import (
	"context"
	"database/sql"
	"errors"
	"rsch/cart_service/helper"
	"rsch/cart_service/model/domain"
)

type Cart struct{}

func NewCart() domain.CartRepository {
	return &Cart{}
}

func (repository *Cart) Create(ctx context.Context, tx *sql.Tx, cart domain.Cart) string {
	SQL := "INSERT INTO carts (id, id_user, id_product, id_product_size, quantity) VALUES (?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, cart.Id, cart.IdUser, cart.IdProduct, cart.IdProductSize, cart.Quantity)
	helper.PanicIfError(err)

	return cart.Id
}

func (repository *Cart) Update(ctx context.Context, tx *sql.Tx, idCart string, qty int) error {
	SQL := "UPDATE carts SET quantity = ? WHERE id = ?"
	res, err := tx.ExecContext(ctx, SQL, qty, idCart)
	helper.PanicIfError(err)

	rows, err := res.RowsAffected()
	helper.PanicIfError(err)

	if rows > 0 {
		return nil
	}

	return errors.New("product in cart not found")
}

func (repository *Cart) UpdateSelected(ctx context.Context, tx *sql.Tx, idCart string, isSelected int) error {
	SQL := "UPDATE carts SET is_selected = ? WHERE id = ?"
	res, err := tx.ExecContext(ctx, SQL, isSelected, idCart)
	helper.PanicIfError(err)

	rows, err := res.RowsAffected()
	helper.PanicIfError(err)

	if rows > 0 {
		return nil
	}

	return errors.New("product in cart not found")
}

func (repository *Cart) Delete(ctx context.Context, tx *sql.Tx, idCart string) error {
	SQL := "DELETE FROM carts WHERE id = ?"
	res, err := tx.ExecContext(ctx, SQL, idCart)
	helper.PanicIfError(err)

	rows, err := res.RowsAffected()
	helper.PanicIfError(err)

	if rows > 0 {
		return nil
	}

	return errors.New("product in cart not found")
}

func (repository *Cart) FindProductsInCartByIdUser(ctx context.Context, tx *sql.Tx, idUser string) []domain.Cart {
	SQL := `
		SELECT c.id, c.id_product, c.id_product_size, c.is_selected, p.name, p.id_image, ps.size, ps.stock, c.quantity, p.price * c.quantity FROM carts c
		INNER JOIN products p ON p.id = c.id_product
		INNER JOIN product_size ps ON ps.id = c.id_product_size
		WHERE c.id_user = ?
		ORDER BY c.created_at DESC
	`
	rows, err := tx.QueryContext(ctx, SQL, idUser)
	helper.PanicIfError(err)

	carts := []domain.Cart{}
	for rows.Next() {
		cart := domain.Cart{}
		err := rows.Scan(&cart.Id, &cart.IdProduct, &cart.IdProductSize, &cart.IsSelected, &cart.ProductName, &cart.ProductImage, &cart.ProductSize, &cart.ProductStock, &cart.Quantity, &cart.TotalPrice)
		helper.PanicIfError(err)

		carts = append(carts, cart)
	}

	return carts
}

func (repository *Cart) FindProductInCartIsExist(ctx context.Context, tx *sql.Tx, cart domain.Cart) (domain.Cart, error) {
	SQL := "SELECT id, quantity FROM carts WHERE id_user = ? AND id_product = ? AND id_product_size = ?"
	rows, err := tx.QueryContext(ctx, SQL, cart.IdUser, cart.IdProduct, cart.IdProductSize)
	helper.PanicIfError(err)

	if rows.Next() {
		err := rows.Scan(&cart.Id, &cart.Quantity)
		helper.PanicIfError(err)
		return cart, nil
	}

	return cart, errors.New("product in cart doesn't exist")
}
