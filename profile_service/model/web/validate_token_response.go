package web

type ValidateTokenResponse struct {
	StatusCode    int                       `json:"status_code"`
	StatusMessage string                    `json:"status_message"`
	Data          DataValidateTokenResponse `json:"data"`
}

type DataValidateTokenResponse struct {
	Id string `json:"id"`
}
