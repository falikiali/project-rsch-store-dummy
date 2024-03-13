package service

import (
	"context"
	"database/sql"
	"rsch/profile_service/exception"
	"rsch/profile_service/helper"
	"rsch/profile_service/model/domain"
	"rsch/profile_service/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type User struct {
	UserRepository           domain.UserRepository
	AuthenticationRepository domain.AuthenticationRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewUser(userRepository domain.UserRepository, authenticationRepository domain.AuthenticationRepository, DB *sql.DB, validate *validator.Validate) domain.UserService {
	return &User{
		UserRepository:           userRepository,
		AuthenticationRepository: authenticationRepository,
		DB:                       DB,
		Validate:                 validate,
	}
}

func (service *User) Create(ctx context.Context, createUserRequest web.CreateUserRequest) web.CreateUserResponse {
	err := service.Validate.Struct(createUserRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Id:       uuid.NewString(),
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

func (service *User) FindUserById(ctx context.Context, token string) web.FindUserById {
	user, err := service.AuthenticationRepository.ValidateToken(ctx, token)
	if err != nil {
		if err.Error() == "" {
			panic(err)
		} else {
			panic(exception.NewUnauthorizedError("bearer token is invalid"))
		}
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err = service.UserRepository.FindUserById(ctx, tx, user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.FindUserById{
		Email:       user.Email,
		Fullname:    user.Fullname,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
	}
}

func (service *User) FindUserByEmailAndPassword(ctx context.Context, email string, password string) web.FindUserEmailPasswordResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Email:    email,
		Password: password,
	}

	user, err = service.UserRepository.FindUserByEmailAndPassword(ctx, tx, user)
	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	return web.FindUserEmailPasswordResponse{
		Id: user.Id,
	}
}

func (service *User) ChangePassword(ctx context.Context, token string, changePasswordRequest web.ChangePasswordRequest) {
	user, err := service.AuthenticationRepository.ValidateToken(ctx, token)
	if err != nil {
		if err.Error() == "" {
			panic(err)
		} else {
			panic(exception.NewUnauthorizedError("bearer token is invalid"))
		}
	}

	err = service.Validate.Struct(changePasswordRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user = domain.User{
		Id: user.Id,
	}

	user, err = service.UserRepository.FindOldPassword(ctx, tx, user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = service.Validate.VarWithValue(changePasswordRequest.OldPassword, user.Password, "eqfield")
	if err != nil {
		panic(exception.NewBadRequestError("Your old password is wrong"))
	}

	user = domain.User{
		Id:       user.Id,
		Password: changePasswordRequest.NewPassword,
	}

	err = service.UserRepository.UpdatePassword(ctx, tx, user)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	//Lakukan komunikasi dengan Auth Service untuk melakukan penghapusan seluruh session user tsb setelah update password.
	//Dengan tujuan melakukan relogin ketika berhasil mengubah password
}

func (service *User) UpdateProfile(ctx context.Context, token string, updateProfileRequest web.UpdateProfileRequest) web.UpdateProfileResponse {
	user, err := service.AuthenticationRepository.ValidateToken(ctx, token)
	if err != nil {
		if err.Error() == "" {
			panic(err)
		} else {
			panic(exception.NewUnauthorizedError("bearer token is invalid"))
		}
	}

	err = service.Validate.Struct(updateProfileRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user = domain.User{
		Id:       user.Id,
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

func (service *User) UpdatePhoneNumber(ctx context.Context, token string, updatePhoneNumberRequest web.UpdatePhoneNumberRequest) web.UpdatePhoneNumberResponse {
	user, err := service.AuthenticationRepository.ValidateToken(ctx, token)
	if err != nil {
		if err.Error() == "" {
			panic(err)
		} else {
			panic(exception.NewUnauthorizedError("bearer token is invalid"))
		}
	}

	err = service.Validate.Struct(updatePhoneNumberRequest)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user = domain.User{
		Id:          user.Id,
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
