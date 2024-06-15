package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("Available commands:")
	fmt.Println("help - Displays a help message")
	fmt.Println("exit - Exit the Pokedex")
	return nil
}

func commandExit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}

func startPokedex() {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Type 'help' to see a list of commands")

	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		command, ok := commands[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Println("Error:", err)
		}

	}

}

func main() {
	startPokedex()

}
