package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lCase := strings.ToLower(text)
	words := strings.Fields(lCase)
	return words
}

func startRepl(cfg *config) {
	availableCmds := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}
		cmdName := cleaned[0]
		command, ok := availableCmds[cmdName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg, cleaned[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}
