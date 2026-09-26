package pokeapi

import (
	"net/http"
	"time"

	"github.com/pokemon/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(cacheInterval, time time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time,
		},
	}
}
