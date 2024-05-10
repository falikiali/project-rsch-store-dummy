package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"rsch/product_service/helper"
	"rsch/product_service/model/domain"
	"strconv"
	"strings"
)

type Product struct{}

func NewProduct() domain.ProductRepository {
	return &Product{}
}

func (repository *Product) Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO products (id, name, price, id_category, purpose, id_image, description) VALUES(?, ?, ?, ?, ?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, product.Id, product.Name, product.Price, product.Category, product.Purpose, product.Image, product.Description)
	helper.PanicIfError(err)

	return product
}

func (repository *Product) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE products SET name = ?, description = ?, price = ?, id_category = ?, purpose = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Description, product.Price, product.Category, product.Purpose, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *Product) FindProducts(ctx context.Context, tx *sql.Tx, page int, filters map[string]interface{}) []domain.Product {
	SQL := `SELECT p.id, p.name, p.description, p.price, p.purpose, p.id_category, p.id_image, SUM(ps.stock) as stock FROM products p
		INNER JOIN product_size ps ON ps.id_product = p.id
	`

	var args []interface{}

	whereClause := ""
	for key, value := range filters {
		if whereClause != "" {
			whereClause += " AND "
		}

		if key == "name" {
			whereClause += fmt.Sprintf("%s LIKE ?", key)
		} else {
			whereClause += fmt.Sprintf("%s = ?", key)
		}
		args = append(args, value)
	}

	if whereClause != "" {
		SQL += " WHERE " + whereClause + " GROUP BY p.id ORDER BY p.created_at DESC LIMIT 2 OFFSET ?"
	} else {
		SQL += " GROUP BY p.id ORDER BY p.created_at DESC LIMIT 2 OFFSET ?"
	}

	args = append(args, (page-1)*2)

	rows, err := tx.QueryContext(ctx, SQL, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	products := []domain.Product{}

	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Purpose, &product.Category, &product.Image, &product.Stock)
		helper.PanicIfError(err)

		products = append(products, product)
	}

	return products
}

func (repository *Product) CountProducts(ctx context.Context, tx *sql.Tx, filters map[string]interface{}) (int, int) {
	SQL := "SELECT COUNT(id) FROM products"

	var args []interface{}

	whereClause := ""
	for key, value := range filters {
		if whereClause != "" {
			whereClause += " AND "
		}

		if key == "name" {
			whereClause += fmt.Sprintf("%s LIKE ?", key)
		} else {
			whereClause += fmt.Sprintf("%s = ?", key)
		}
		args = append(args, value)
	}

	if whereClause != "" {
		SQL += " WHERE " + whereClause
	}

	rows, err := tx.QueryContext(ctx, SQL, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		var totalData int
		var totalPage int

		err := rows.Scan(&totalData)
		helper.PanicIfError(err)

		totalPage = totalData / 2
		if totalData%2 != 0 {
			totalPage += 1
		}

		return totalData, totalPage
	}

	return 0, 0
}

func (repository *Product) FindProductById(ctx context.Context, tx *sql.Tx, idProduct string) (domain.Product, error) {
	SQL := `SELECT p.id, p.name, p.description, p.price, p.purpose, p.id_category, p.id_image, GROUP_CONCAT(ps.id), GROUP_CONCAT(ps.size), GROUP_CONCAT(ps.stock), SUM(ps.stock) as stock FROM products p
		INNER JOIN product_size ps ON ps.id_product = p.id
		WHERE p.id = ?
		GROUP BY p.id
	`

	rows, err := tx.QueryContext(ctx, SQL, idProduct)
	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.Product{}

	if rows.Next() {
		var idProductSizeStr, sizeStr, stockStr string

		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Purpose, &product.Category, &product.Image, &idProductSizeStr, &sizeStr, &stockStr, &product.Stock)
		helper.PanicIfError(err)

		idProductSizes := strings.Split(idProductSizeStr, ",")
		sizes := strings.Split(sizeStr, ",")
		stocks := strings.Split(stockStr, ",")

		for i := range idProductSizes {
			stock, err := strconv.Atoi(stocks[i])
			helper.PanicIfError(err)

			product.DetailSize = append(product.DetailSize, domain.ProductSize{
				Id:    idProductSizes[i],
				Size:  sizes[i],
				Stock: stock,
			})
		}

		return product, nil
	}

	return product, errors.New("product not found")
}
