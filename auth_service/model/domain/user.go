package domain

import (
	"context"
	"rsch/auth_service/model/web"
)

type User struct {
	Id          string
	Email       string
	Password    string
	Username    string
	PhoneNumber string
}

type UserRepository interface {
	CreateUser(ctx context.Context, request web.RegisterRequest) (User, error)
	FindUserByEmailAndPassword(ctx context.Context, request web.LoginRequest) (User, error)
}
