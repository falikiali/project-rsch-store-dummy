package app

import (
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	client := &http.Client{Timeout: 12 * time.Second}
	return client
}
