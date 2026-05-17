package pokeapi

import (
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

type Client struct {
	userCache *pokecache.Cache
	cli       http.Client
}

func NewClient(cacheInterval time.Duration, timeout time.Duration) *Client {
	c := &Client{
		userCache: pokecache.NewCache(cacheInterval),
	}

	return c
}
