package repository

import (
	"context"
	"database/sql"
	"errors"
	"rsch/auth_service/helper"
	"rsch/auth_service/model/domain"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) string {
	SQL := "INSERT INTO users (id, email, password) VALUES(?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Email, user.Password)
	helper.PanicIfError(err)

	return user.Id
}

func (repository *UserRepositoryImpl) FindUserByEmail(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "SELECT id FROM users WHERE password = ? AND email = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.Password, user.Email)
	helper.PanicIfError(err)
	defer rows.Close()

	user = domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id)
		helper.PanicIfError(err)

		return user, nil
	}

	return user, errors.New("incorrect email or password")
}

func (repository *UserRepositoryImpl) FindUserIsExist(ctx context.Context, tx *sql.Tx, user domain.User) error {
	SQL := "SELECT id FROM users WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.Email)
	helper.PanicIfError(err)
	defer rows.Close()

	user = domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id)
		helper.PanicIfError(err)

		return errors.New("email has already been registered")
	}

	return nil
}
