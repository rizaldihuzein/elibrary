package restapi

import (
	"net/http"
	"time"
)

type api struct {
	http    *http.Client
	timeout time.Duration
	baseURL string
}

func New() *api {
	timeout := 15 * time.Second
	return &api{
		timeout: timeout,
		http: &http.Client{
			Timeout: timeout,
		},
		baseURL: "https://openlibrary.org",
	}
}
