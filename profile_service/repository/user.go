package repository

import (
	"context"
	"database/sql"
	"errors"
	"rsch/profile_service/helper"
	"rsch/profile_service/model/domain"
)

type User struct{}

func NewUser() domain.UserRepository {
	return &User{}
}

func (repository *User) Create(ctx context.Context, tx *sql.Tx, user domain.User) string {
	SQL := "INSERT INTO users (id, email, password) VALUES(?, ?, ?)"
	_, err := tx.ExecContext(ctx, SQL, user.Id, user.Email, user.Password)
	helper.PanicIfError(err)

	return user.Id
}

func (repository *User) UpdateFullname(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "UPDATE users SET fullname = ? WHERE id = ?"
	res, err := tx.ExecContext(ctx, SQL, user.Fullname, user.Id)
	helper.PanicIfError(err)

	rows, err := res.RowsAffected()
	helper.PanicIfError(err)

	if rows > 0 {
		return user, nil
	}

	return user, errors.New("user not found")
}

func (repository *User) UpdatePassword(ctx context.Context, tx *sql.Tx, user domain.User) error {
	SQL := "UPDATE users SET password = ? WHERE id = ?"
	res, err := tx.ExecContext(ctx, SQL, user.Password, user.Id)
	helper.PanicIfError(err)

	rows, err := res.RowsAffected()
	helper.PanicIfError(err)

	if rows > 0 {
		return nil
	}

	return errors.New("user not found")
}

func (repository *User) UpdateUsername(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "UPDATE users SET username = ? WHERE id = ?"
	res, err := tx.ExecContext(ctx, SQL, user.Username, user.Id)
	helper.PanicIfError(err)

	rows, err := res.RowsAffected()
	helper.PanicIfError(err)

	if rows > 0 {
		return user, nil
	}

	return user, errors.New("user not found")
}

func (repository *User) UpdatePhoneNumber(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	SQL := "UPDATE users SET fullname = ? WHERE id = ?"
	res, err := tx.ExecContext(ctx, SQL, user.PhoneNumber, user.Id)
	helper.PanicIfError(err)

	rows, err := res.RowsAffected()
	helper.PanicIfError(err)

	if rows > 0 {
		return user, nil
	}

	return user, errors.New("user not found")
}

func (repository *User) FindUserByEmailAndPassword(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
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

func (repository *User) FindEmailIsExist(ctx context.Context, tx *sql.Tx, user domain.User) error {
	SQL := "SELECT id FROM users WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.Email)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		return errors.New("email has already been registered")
	}

	return nil
}

func (repository *User) FindUsernameIsExist(ctx context.Context, tx *sql.Tx, user domain.User) error {
	SQL := "SELECT id FROM users WHERE username = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.Username)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		return errors.New("username has already been used")
	}

	return nil
}

func (repository *User) FindPhoneNumberIsExist(ctx context.Context, tx *sql.Tx, user domain.User) error {
	SQL := "SELECT id FROM users WHERE phone_number = ?"
	rows, err := tx.QueryContext(ctx, SQL, user.PhoneNumber)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		return errors.New("phone number has already been used")
	}

	return nil
}
