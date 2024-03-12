package domain

import (
	"context"
	"database/sql"
	"net/http"
	"rsch/profile_service/model/web"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	Id          string
	Email       string
	Password    string
	Fullname    string
	PhoneNumber string
	Username    string
}

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user User) string
	UpdateFullname(ctx context.Context, tx *sql.Tx, user User) (User, error)
	UpdatePassword(ctx context.Context, tx *sql.Tx, user User) error
	UpdateUsername(ctx context.Context, tx *sql.Tx, user User) (User, error)
	UpdatePhoneNumber(ctx context.Context, tx *sql.Tx, user User) (User, error)
	FindUserByEmailAndPassword(ctx context.Context, tx *sql.Tx, user User) (User, error)
	FindEmailIsExist(ctx context.Context, tx *sql.Tx, user User) error
	FindUsernameIsExist(ctx context.Context, tx *sql.Tx, user User) error
	FindPhoneNumberIsExist(ctx context.Context, tx *sql.Tx, user User) error
}

type UserService interface {
	Create(ctx context.Context, createUserRequest web.CreateUserRequest) web.CreateUserResponse
	ChangePassword(ctx context.Context, changePasswordRequest web.ChangePasswordRequest)
	UpdateProfile(ctx context.Context, updateProfileRequest web.UpdateProfileRequest) web.UpdateProfileResponse
	UpdatePhoneNumber(ctx context.Context, updatePhoneNumberRequest web.UpdatePhoneNumberRequest) web.UpdatePhoneNumberResponse
}

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	ChangePassword(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdatePhoneNumber(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
