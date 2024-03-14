package repository

import (
	"context"
	"database/sql"
	"errors"
	"rsch/category_service/helper"
	"rsch/category_service/model/domain"
	"strconv"
)

type Category struct{}

func NewCategory() domain.CategoryRepository {
	return &Category{}
}

func (repository *Category) Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories (name) VALUES(?)"
	res, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int16(id)

	return category
}

func (repository *Category) Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	SQL := "UPDATE categories SET name = ? WHERE id = ?"
	res, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	rowAffected, err := res.RowsAffected()
	helper.PanicIfError(err)

	if rowAffected > 0 {
		return category, nil
	}

	return category, errors.New("category not found")
}

func (repository *Category) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	categories := []domain.Category{}

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}

func (repository *Category) FindById(ctx context.Context, tx *sql.Tx, id int16) error {
	SQL := "SELECT id FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		return nil
	}

	return errors.New("category not found")
}

func (repository *Category) FindNameIsExist(ctx context.Context, tx *sql.Tx, name string) (domain.Category, error) {
	SQL := "SELECT id, name FROM categories WHERE name = ?"
	rows, err := tx.QueryContext(ctx, SQL, name)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		return category, errors.New(category.Name + " category already exist with ID " + strconv.Itoa(int(category.Id)))
	}

	category.Name = name
	return category, nil
}
