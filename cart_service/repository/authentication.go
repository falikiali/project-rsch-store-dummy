package repository

import (
	"context"
	"errors"
	"net/http"
	"rsch/cart_service/helper"
	"rsch/cart_service/model/domain"
	"rsch/cart_service/model/web"
)

type Authentication struct {
	Client *http.Client
}

func NewAuthentication(client *http.Client) domain.AuthenticationRepository {
	return &Authentication{
		Client: client,
	}
}

func (repository *Authentication) ValidateToken(ctx context.Context, accessToken string) (string, error) {
	endpoint := "http://192.168.1.9:3000/api/v3/authentication"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	helper.PanicIfError(err)

	req.Header.Set("Authorization", "Bearer "+accessToken)

	response, err := repository.Client.Do(req)
	helper.PanicIfError(err)
	defer response.Body.Close()

	var idUser string

	if response.StatusCode == 401 {
		webResponse := web.WebResponse{}
		helper.ReadFromResponseBody(response, &webResponse)

		data, ok := webResponse.Data.(string)
		if ok {
			return idUser, errors.New(data)
		}
	}

	if response.StatusCode == 200 {
		validateTokenResponse := web.ValidateTokenResponse{}
		helper.ReadFromResponseBody(response, &validateTokenResponse)
		idUser = validateTokenResponse.Data.Id

		return idUser, nil
	}

	return idUser, errors.New("")
}
