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
		pokeapiClient: pokeapi.NewClient(30*time.Minute, 5*time.Minute),
	}
	err := startRepl(&cfg)
	if err != nil {
		log.Fatal(err)
	}
}
