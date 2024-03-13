package web

type FindUserById struct {
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
}
