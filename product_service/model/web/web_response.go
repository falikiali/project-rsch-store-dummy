package web

type WebResponse struct {
	StatusCode    int         `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Data          interface{} `json:"data,omitempty"`
	Pagination    interface{} `json:"pagination,omitempty"`
}
