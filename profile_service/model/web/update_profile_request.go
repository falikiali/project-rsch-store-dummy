package web

type UpdateProfileRequest struct {
	Fullname string `json:"fullname" validate:"required"`
	Username string `json:"username" validate:"required,excludes= "`
}
