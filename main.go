package main

import (
	"log"
	"time"

	"github.com/dimadudin/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	previousLocationURL *string
	nextLocationURL     *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(10 * time.Second),
	}
	err := startRepl(&cfg)
	if err != nil {
		log.Fatal(err)
	}
}
