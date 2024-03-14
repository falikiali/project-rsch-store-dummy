package web

type WebResponse struct {
	StatusCode    int         `json:"status_code"`
	StatudMessage string      `json:"status_message"`
	Data          interface{} `json:"data,omitempty"`
}
