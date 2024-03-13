package repository

import (
	"context"
	"database/sql"
	"errors"
	"rsch/auth_service/helper"
	"rsch/auth_service/model/domain"
)

type Authentication struct{}

func NewAuthentication() domain.AuthenticationRepository {
	return &Authentication{}
}

func (repository *Authentication) FindByToken(ctx context.Context, tx *sql.Tx, token string) (domain.Authentication, error) {
	SQL := "SELECT token FROM authentications WHERE token = ?"
	rows, err := tx.QueryContext(ctx, SQL, token)
	helper.PanicIfError(err)
	defer rows.Close()

	authentication := domain.Authentication{}

	if rows.Next() {
		err := rows.Scan(&authentication.Token)
		helper.PanicIfError(err)

		return authentication, nil
	}

	return authentication, errors.New("bearer token is invalid")
}

func (repository *Authentication) Create(ctx context.Context, tx *sql.Tx, authentication domain.Authentication) string {
	SQL := "INSERT INTO authentications(id, token) VALUES(?, ?)"
	_, err := tx.ExecContext(ctx, SQL, authentication.Id, authentication.Token)
	helper.PanicIfError(err)

	return authentication.Token
}

func (repository *Authentication) Delete(ctx context.Context, tx *sql.Tx, token string) error {
	SQL := "DELETE FROM authentications WHERE token = ?"
	res, err := tx.ExecContext(ctx, SQL, token)
	helper.PanicIfError(err)

	countRowAffected, err := res.RowsAffected()
	helper.PanicIfError(err)

	if countRowAffected > 0 {
		return nil
	}

	return errors.New("your session has expired")
}
