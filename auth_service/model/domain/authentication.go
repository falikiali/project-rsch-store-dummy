package domain

import (
	"context"
	"database/sql"
	"net/http"
	"rsch/auth_service/model/web"

	"github.com/julienschmidt/httprouter"
)

type Authentication struct {
	Id    string
	Token string
}

type AuthenticationRepository interface {
	FindByToken(ctx context.Context, tx *sql.Tx, token string) (Authentication, error)
	Create(ctx context.Context, tx *sql.Tx, authentication Authentication) string
	Delete(ctx context.Context, tx *sql.Tx, token string) error
}

type AuthenticationService interface {
	ValidateToken(ctx context.Context, token string) web.ValidateTokenResponse
	Register(ctx context.Context, request web.RegisterRequest) web.AuthenticationResponse
	Login(ctx context.Context, request web.LoginRequest) web.AuthenticationResponse
	Logout(ctx context.Context, request web.LogoutRequest)
}

type AuthenticationController interface {
	ValidateToken(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Register(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Logout(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
