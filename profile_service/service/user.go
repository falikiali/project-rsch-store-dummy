package service

import (
	"context"
	"database/sql"
	"rsch/profile_service/exception"
	"rsch/profile_service/helper"
	"rsch/profile_service/model/domain"
	"rsch/profile_service/model/web"

	"github.com/go-playground/validator/v10"
)

type User struct {
	UserRepository domain.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUser(userRepository domain.UserRepository, DB *sql.DB, validate *validator.Validate) domain.UserService {
	return &User{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *User) Create(ctx context.Context, createUserRequest web.CreateUserRequest) web.CreateUserResponse {
	err := service.Validate.Struct(createUserRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	err = service.UserRepository.FindEmailIsExist(ctx, tx, user)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	user.Id = service.UserRepository.Create(ctx, tx, user)

	return web.CreateUserResponse{
		Id: user.Id,
	}
}

func (service *User) ChangePassword(ctx context.Context, changePasswordRequest web.ChangePasswordRequest) {
	err := service.Validate.Struct(changePasswordRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Id:       "id",
		Password: changePasswordRequest.NewPassword,
	}

	err = service.UserRepository.UpdatePassword(ctx, tx, user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	//Lakukan pemanggilan api logout dari auth service setelah update password.
	//Dengan tujuan melakukan relogin ketika berhasil mengubah password
}

func (service *User) UpdateProfile(ctx context.Context, updateProfileRequest web.UpdateProfileRequest) web.UpdateProfileResponse {
	err := service.Validate.Struct(updateProfileRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Id:       "id",
		Username: updateProfileRequest.Username,
		Fullname: updateProfileRequest.Fullname,
	}

	err = service.UserRepository.FindUsernameIsExist(ctx, tx, user)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	user, err = service.UserRepository.UpdateUsername(ctx, tx, user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user, err = service.UserRepository.UpdateFullname(ctx, tx, user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.UpdateProfileResponse{
		Username: user.Username,
		Fullname: user.Fullname,
	}
}

func (service *User) UpdatePhoneNumber(ctx context.Context, updatePhoneNumberRequest web.UpdatePhoneNumberRequest) web.UpdatePhoneNumberResponse {
	err := service.Validate.Struct(updatePhoneNumberRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Id:          "id",
		PhoneNumber: updatePhoneNumberRequest.PhoneNumber,
	}

	err = service.UserRepository.FindPhoneNumberIsExist(ctx, tx, user)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	user, err = service.UserRepository.UpdatePhoneNumber(ctx, tx, user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.UpdatePhoneNumberResponse{
		NewPhoneNumber: user.PhoneNumber,
	}
}
