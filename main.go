package main

import (
	"time"

	"github.com/dimadudin/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	previousLocationURL *string
	nextLocationURL     *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(30*time.Minute, 5*time.Minute),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
