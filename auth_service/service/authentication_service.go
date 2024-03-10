package service

import (
	"context"
	"rsch/auth_service/model/web"
)

type AuthenticationService interface {
	Register(ctx context.Context, request web.RegisterRequest) web.AuthenticationResponse
	Login(ctx context.Context, request web.LoginRequest) web.AuthenticationResponse
	Logout(ctx context.Context, request web.LogoutRequest)
}
