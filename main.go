package main

import (
	"bufio"
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex

`)
	return nil
}

type cliCommand struct {
	name     string
	desc     string
	callback func() error
}

var commandMap = map[string]cliCommand{
	"exit": {
		name:     "exit",
		desc:     "Exit the Pokedex",
		callback: commandExit,
	},
	"help": {
		name:     "help",
		desc:     "Return a tutorial",
		callback: commandHelp,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		// nothing left
		if scanner.Scan() == false {
			break
		}

		query := scanner.Text()
		cleanedQuery := cleanInput(query)

		if len(cleanedQuery) == 0 {
			fmt.Print("invalid input")
		}

		command, exists := commandMap[cleanedQuery[0]]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}
