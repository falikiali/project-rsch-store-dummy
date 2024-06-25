package web

type PaginationResponse struct {
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
}
