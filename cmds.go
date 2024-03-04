package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    cmdHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the programm",
			callback:    cmdExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 next locations",
			callback:    cmdMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays 20 previous locations",
			callback:    cmdMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Takes an area name as argument, displays all pokemon found in area",
			callback:    cmdExplore,
		},
		"catch": {
			name:        "catch",
			description: "Takes the pokemon name as argument, attempts to catch the pokemon",
			callback:    cmdCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Takes the pokemon name as argument, displays all stats if the pokemon was caught",
			callback:    cmdInspect,
		},
	}
}

func cmdHelp(cfg *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("too many arguments")
	}
	fmt.Println("")
	fmt.Println("Usage: <command>")
	fmt.Println("All commands:")
	availableCmds := getCommands()
	for _, cmd := range availableCmds {
		fmt.Printf("- %s : %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}

func cmdExit(cfg *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("too many arguments")
	}
	os.Exit(0)
	return nil
}

func cmdMap(cfg *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("too many arguments")
	}
	resp, err := cfg.pokeapiClient.ListAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, loc := range resp.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}
	cfg.nextLocationURL = resp.Next
	cfg.previousLocationURL = resp.Previous
	return nil
}

func cmdMapBack(cfg *config, args ...string) error {
	if len(args) > 0 {
		return errors.New("too many arguments")
	}
	if cfg.previousLocationURL == nil {
		return errors.New("can't go back, you're on the first page")
	}
	areaList, err := cfg.pokeapiClient.ListAreas(cfg.previousLocationURL)
	if err != nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, loc := range areaList.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}
	cfg.nextLocationURL = areaList.Next
	cfg.previousLocationURL = areaList.Previous
	return nil
}

func cmdExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}
	if len(args) > 1 {
		return errors.New("too many arguments")
	}
	locName := args[0]
	area, err := cfg.pokeapiClient.GetArea(locName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locName)
	fmt.Println("Found pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func cmdCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}
	if len(args) > 1 {
		return errors.New("too many arguments")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	exp := pokemon.BaseExperience
	fmt.Printf("Throwing pokeball at %s with exp %d\n", pokemonName, exp)
	threshold := 50
	roll := rand.Intn(exp)
	if roll < threshold {
		cfg.caughtPokemon[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

func cmdInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}
	if len(args) > 1 {
		return errors.New("too many arguments")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Println("Name: ", pokemonName)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" -%s\n", t.Type.Name)
	}
	return nil
}
