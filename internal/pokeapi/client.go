package pokeapi

import (
	"net/http"
	"time"

	"github.com/Hien-Trinh/pokedex/internal/pokeapi"
	"github.com/Hien-Trinh/pokedex/internal/pokecache"
)

type Client struct {
	cache         pokecache.Cache
	caughtPokemon map[string]pokeapi.Pokemon
	httpClient    http.Client
}

func NewClient(cacheInterval time.Duration, timeout time.Duration) Client {
	return Client{
		cache:         pokecache.NewCache(cacheInterval),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
