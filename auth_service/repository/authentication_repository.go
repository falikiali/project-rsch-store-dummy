package repository

import (
	"context"
	"database/sql"
	"rsch/auth_service/model/domain"
)

type AuthenticationRepository interface {
	Create(ctx context.Context, tx *sql.Tx, authentication domain.Authentication) string
	Delete(ctx context.Context, tx *sql.Tx, token string)
}
