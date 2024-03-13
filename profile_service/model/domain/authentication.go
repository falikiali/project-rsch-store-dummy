package domain

import (
	"context"
)

type AuthenticationRepository interface {
	ValidateToken(ctx context.Context, accessToken string) (User, error)
}
