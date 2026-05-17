package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	c := config{
		pokeClient: pokeapi.NewClient(5*time.Second, 5*time.Second),
		previous:   nil,
		next:       nil,
	}

	replInit(c)
}
