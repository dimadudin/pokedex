package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
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

// fetches GET response from cache if possible,
// otherwise makes a network GET request and adds the response to cache
func (c *Client) get(URL string) ([]byte, error) {
	dat, ok := c.pokeCache.Get(URL)
	if ok {
		return dat, nil
	}

	resp, err := c.httpClient.Get(URL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("response failed with status code: %d", resp.StatusCode)
	}
	dat, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	c.pokeCache.Add(URL, dat)
	return dat, err
}

func unmarshalTo[T any](dat []byte) (T, error) {
	var target T
	err := json.Unmarshal(dat, &target)
	return target, err
}
