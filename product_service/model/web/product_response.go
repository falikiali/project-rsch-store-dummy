package web

type ProductResponse struct {
	Id          string               `json:"id"`
	Name        string               `json:"name"`
	Price       int64                `json:"price"`
	Description string               `json:"description"`
	Purpose     string               `json:"purpose"`
	Category    int16                `json:"category"`
	Stock       int                  `json:"stock"`
	Image       string               `json:"image"`
	DetailSize  []DetailSizeResponse `json:"detail_size,omitempty"`
}

type DetailSizeResponse struct {
	Id    string `json:"id,omitempty"`
	Size  string `json:"size"`
	Stock int    `json:"stock"`
}
