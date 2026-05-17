package main

import (
	"pokedex/internal/pokeapi"
)

type config struct {
	pokeClient *pokeapi.Client
	previous   *string
	next       *string
}
