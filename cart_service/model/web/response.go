package web

type WebResponse struct {
	StatusCode    int         `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Data          interface{} `json:"data,omitempty"`
}

type ValidateTokenResponse struct {
	StatusCode    int                       `json:"status_code"`
	StatusMessage string                    `json:"status_message"`
	Data          DataValidateTokenResponse `json:"data"`
}

type DataValidateTokenResponse struct {
	Id string `json:"id"`
}

type FindProductsInCartByIdUserResponse struct {
	Id            string `json:"id"`
	IdProduct     string `json:"id_product"`
	IdProductSize string `json:"id_product_size"`
	IsSelected    bool   `json:"is_selected"`
	ProductName   string `json:"product_name"`
	ProductImage  string `json:"product_image"`
	ProductSize   string `json:"product_size"`
	ProductStock  int    `json:"product_stock"`
	Quantity      int    `json:"qty"`
	TotalPrice    int    `json:"total_price"`
}
