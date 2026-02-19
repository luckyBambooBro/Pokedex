package pokeapi

import (
	"net/http"
	"time"
	"github.com/luckybamboobro/pokedex/internal/cache"
)

type Client struct {
	httpClient http.Client
	cache *cache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client {
		httpClient: http.Client {
		Timeout: timeout,
		},
		cache: cache.NewCache(cacheInterval),
	}
}


