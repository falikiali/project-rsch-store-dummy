package web

type LogoutRequest struct {
	AccessToken string `json:"access_token" validate:"required,jwt"`
}
