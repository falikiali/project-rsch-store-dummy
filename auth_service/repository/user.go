package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"rsch/auth_service/helper"
	"rsch/auth_service/model/domain"
	"rsch/auth_service/model/web"
)

type User struct {
	Client *http.Client
}

func NewUser(client *http.Client) domain.UserRepository {
	return &User{
		Client: client,
	}
}

func (repository *User) CreateUser(ctx context.Context, request web.RegisterRequest) (domain.User, error) {
	endpoint := "http://192.168.1.9:3001/api/v3/user"
	requestBody, err := json.Marshal(request)
	helper.PanicIfError(err)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(requestBody))
	helper.PanicIfError(err)
	req.Header.Set("Content-Type", "application/json")

	response, err := repository.Client.Do(req)
	helper.PanicIfError(err)
	defer response.Body.Close()

	user := domain.User{}

	if response.StatusCode == 400 {
		webResponse := web.WebResponse{}
		helper.ReadFromResponseBody(response, &webResponse)

		data, ok := webResponse.Data.(string)
		if ok {
			return user, errors.New(data)
		}
	}

	if response.StatusCode == 200 {
		userResponse := web.UserResponse{}
		helper.ReadFromResponseBody(response, &userResponse)
		user.Id = userResponse.Data.Id

		return user, nil
	}

	return user, errors.New("")
}

func (repository *User) FindUserByEmailAndPassword(ctx context.Context, request web.LoginRequest) (domain.User, error) {
	endpoint := "http://192.168.1.9:3001/api/v3/user/validate"
	params := url.Values{}
	params.Add("email", request.Email)
	params.Add("password", request.Password)
	endpoint += "?" + params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	helper.PanicIfError(err)

	response, err := repository.Client.Do(req)
	helper.PanicIfError(err)
	defer response.Body.Close()

	user := domain.User{}

	if response.StatusCode == 401 {
		webResponse := web.WebResponse{}
		helper.ReadFromResponseBody(response, &webResponse)

		data, ok := webResponse.Data.(string)
		if ok {
			return user, errors.New(data)
		}
	}

	if response.StatusCode == 200 {
		userResponse := web.UserResponse{}
		helper.ReadFromResponseBody(response, &userResponse)
		user.Id = userResponse.Data.Id

		return user, nil
	}

	return user, errors.New("")
}
