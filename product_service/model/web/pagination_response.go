package web

type PaginationResponse struct {
	Page      int `json:"page,omitempty"`
	TotalPage int `json:"total_page,omitempty"`
	TotalData int `json:"total_data,omitempty"`
}
