package web

type UpdateProductInCartRequest struct {
	Id       string `json:"id" validate:"required"`
	Quantity int    `json:"qty" validate:"required,number"`
}

type UpdateSelectedProductInCartRequest struct {
	Id         string `json:"id" validate:"required"`
	IsSelected bool   `json:"is_selected" validate:"boolean"`
}

type AddProductToCartRequest struct {
	IdProduct     string `json:"id_product" validate:"required"`
	IdProductSize string `json:"id_product_size" validate:"required"`
	Quantity      int    `json:"qty" validate:"required,number"`
}

type DeleteProductInCartRequest struct {
	Id string `json:"id" validate:"required"`
}
