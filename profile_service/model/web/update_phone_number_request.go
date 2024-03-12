package web

type UpdatePhoneNumberRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required,number,startsnotwith=0"`
}
