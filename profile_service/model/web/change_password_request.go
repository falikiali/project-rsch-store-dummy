package web

type ChangePasswordRequest struct {
	OldPassword        string `json:"old_password" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required,min=8,nefield=OldPassword"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required,min=8,eqfield=NewPassword"`
}
