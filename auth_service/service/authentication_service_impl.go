package service

import (
	"context"
	"database/sql"
	"rsch/auth_service/app"
	"rsch/auth_service/exception"
	"rsch/auth_service/helper"
	"rsch/auth_service/model/domain"
	"rsch/auth_service/model/web"
	"rsch/auth_service/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AuthenticationServiceImpl struct {
	AuthenticationRepository repository.AuthenticationRepository
	UserRepository           repository.UserRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewAuthenticationService(authenticationRepository repository.AuthenticationRepository, userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		authenticationRepository,
		userRepository,
		DB,
		validate,
	}
}

func (service *AuthenticationServiceImpl) Register(ctx context.Context, request web.RegisterRequest) web.AuthenticationResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Id:       uuid.NewString(),
		Email:    request.Email,
		Password: request.Password,
	}
	err = service.UserRepository.FindUserIsExist(ctx, tx, user)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	id := service.UserRepository.Create(ctx, tx, user)

	authentication := domain.Authentication{
		Id:    uuid.NewString(),
		Token: app.GenerateNewJwt(id),
	}
	token := service.AuthenticationRepository.Create(ctx, tx, authentication)

	return web.AuthenticationResponse{
		AccessToken: token,
	}
}

func (service *AuthenticationServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.AuthenticationResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err = service.UserRepository.FindUserByEmail(ctx, tx, user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
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

func (service *AuthenticationServiceImpl) Logout(ctx context.Context, request web.LogoutRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.AuthenticationRepository.Delete(ctx, tx, request.AccessToken)
}
