package pokeapi

import (
	"net/http"
	"time"

	"github.com/dimadudin/pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	pokeCache  pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		pokeCache: pokecache.NewCache(cacheInterval),
	}
}
