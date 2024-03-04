package main

import (
	"errors"
	"fmt"
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
	}
}

func cmdHelp(cfg *config, args ...string) error {
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
	os.Exit(0)
	return nil
}

func cmdMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
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
	if cfg.previousLocationURL == nil {
		return errors.New("can't go back, you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationURL)
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

func cmdExplore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("not enough arguments")
	}
	if len(args) > 1 {
		return errors.New("too many arguments")
	}
	locName := args[0]
	resp, err := cfg.pokeapiClient.ListPokemonInArea(locName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locName)
	fmt.Println("Found pokemon:")
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
