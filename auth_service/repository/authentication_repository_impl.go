package repository

import (
	"context"
	"database/sql"
	"rsch/auth_service/helper"
	"rsch/auth_service/model/domain"
)

type AuthenticationRepositoryImpl struct{}

func NewAuthenticationRepository() AuthenticationRepository {
	return &AuthenticationRepositoryImpl{}
}

func (repository *AuthenticationRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, authentication domain.Authentication) string {
	querySql := "INSERT INTO authentications(id, token) VALUES(?, ?)"
	_, err := tx.ExecContext(ctx, querySql, authentication.Id, authentication.Token)
	helper.PanicIfError(err)

	return authentication.Token
}

func (repository *AuthenticationRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, token string) {
	querySql := "DELETE FROM authentications WHERE token = ?"
	_, err := tx.ExecContext(ctx, querySql, token)
	helper.PanicIfError(err)
}
