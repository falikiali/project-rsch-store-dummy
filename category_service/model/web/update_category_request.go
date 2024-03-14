package web

type UpdateCategoryRequest struct {
	Id   int16  `validate:"required"`
	Name string `json:"name" validate:"required"`
}
