package repository

import (
	"context"
	"database/sql"
	"rsch/auth_service/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) string
	FindUserByEmail(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	FindUserIsExist(ctx context.Context, tx *sql.Tx, user domain.User) error
}
