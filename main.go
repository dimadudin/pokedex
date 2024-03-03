package main

import (
	"log"

	"github.com/dimadudin/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	previousLocationURL *string
	nextLocationURL     *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	err := startRepl(&cfg)
	if err != nil {
		log.Fatal(err)
	}
}
