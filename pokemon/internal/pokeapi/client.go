package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(time time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time,
		},
	}
}
