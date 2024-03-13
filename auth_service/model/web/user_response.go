package web

type UserResponse struct {
	StatusCode    int              `json:"status_code"`
	StatusMessage string           `json:"status_message"`
	Data          DataUserResponse `json:"data"`
}

type DataUserResponse struct {
	Id string `json:"id"`
}
