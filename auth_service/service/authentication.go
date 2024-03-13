package service

import (
	"context"
	"database/sql"
	"rsch/auth_service/app"
	"rsch/auth_service/exception"
	"rsch/auth_service/helper"
	"rsch/auth_service/model/domain"
	"rsch/auth_service/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Authentication struct {
	AuthenticationRepository domain.AuthenticationRepository
	UserRepository           domain.UserRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewAuthentication(authenticationRepository domain.AuthenticationRepository, userRepository domain.UserRepository, DB *sql.DB, validate *validator.Validate) domain.AuthenticationService {
	return &Authentication{
		authenticationRepository,
		userRepository,
		DB,
		validate,
	}
}

func (service *Authentication) ValidateToken(ctx context.Context, token string) web.ValidateTokenResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	authentication, err := service.AuthenticationRepository.FindByToken(ctx, tx, token)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	user, err := app.ParseJWT(authentication.Token)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	return web.ValidateTokenResponse{
		Id: user.Id,
	}
}

func (service *Authentication) Register(ctx context.Context, request web.RegisterRequest) web.AuthenticationResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.CreateUser(ctx, request)
	if err != nil {
		if err.Error() == "" {
			panic(err)
		} else {
			panic(exception.NewBadRequestError(err.Error()))
		}
	}

	authentication := domain.Authentication{
		Id:    uuid.NewString(),
		Token: app.GenerateNewJwt(user.Id),
	}
	token := service.AuthenticationRepository.Create(ctx, tx, authentication)

	return web.AuthenticationResponse{
		AccessToken: token,
	}
}

func (service *Authentication) Login(ctx context.Context, request web.LoginRequest) web.AuthenticationResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	user, err := service.UserRepository.FindUserByEmailAndPassword(ctx, request)
	if err != nil {
		if err.Error() == "" {
			panic(err)
		} else {
			panic(exception.NewUnauthorizedError(err.Error()))
		}
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	authentication := domain.Authentication{
		Id:    uuid.NewString(),
		Token: app.GenerateNewJwt(user.Id),
	}
	token := service.AuthenticationRepository.Create(ctx, tx, authentication)

	return web.AuthenticationResponse{
		AccessToken: token,
	}
}

func (service *Authentication) Logout(ctx context.Context, request web.LogoutRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.AuthenticationRepository.Delete(ctx, tx, request.AccessToken)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}
}
