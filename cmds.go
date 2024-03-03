package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func cmdHelp() error {
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

func cmdExit() error {
	os.Exit(0)
	return nil
}

type mapResponse struct {
	Next     string
	Previous string
	Results  []struct {
		Name string
	}
}

func cmdMap() error {
	res, reqErr := http.Get("https://pokeapi.co/api/v2/location-area/")
	if reqErr != nil {
		return reqErr
	}
	text, readErr := io.ReadAll(res.Body)
	res.Body.Close()
	if readErr != nil {
		return readErr
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("response failed with status code %d", res.StatusCode)
	}
	resp := mapResponse{}
	jsonErr := json.Unmarshal(text, &resp)
	if jsonErr != nil {
		return jsonErr
	}
	locations := resp.Results
	for _, loc := range locations {
		fmt.Println(loc.Name)
	}
	return nil
}

func cmdMapBack() error {
	return nil
}
