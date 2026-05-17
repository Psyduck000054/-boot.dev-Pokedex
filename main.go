package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"time"
)

type config struct {
	pokeClient *pokeapi.Client
	previous   *string
	next       *string
}

// ---------------------------------------------------------
// COMMANDS
// ---------------------------------------------------------

func commandExit(c *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

#MISCELLANEOUS
help: Displays a help message
exit: Exit the Pokedex

#MAP
map: Shows the next page of in-game locations
mapb: Shows the previous page of in-game locations

`)
	return nil
}

func commandMap(c *config) error {
	res, err := c.pokeClient.ListLocations(c.next)
	if err != nil {
		return err
	}

	c.next = res.Next
	c.previous = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(c *config) error {
	if c.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := c.pokeClient.ListLocations(c.previous)
	if err != nil {
		return err
	}

	c.next = res.Next
	c.previous = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

// ---------------------------------------------------------
// COMMAND MAP
// ---------------------------------------------------------

type cliCommand struct {
	name     string
	desc     string
	callback func(c *config) error
}

var commandList = map[string]cliCommand{
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
	"map": {
		name:     "map",
		desc:     "List the locations of the next page",
		callback: commandMap,
	},
	"mapb": {
		name:     "mapb",
		desc:     "List the locations of the previous page",
		callback: commandMapb,
	},
}

// ---------------------------------------------------------
// MAIN FUNCTION
// ---------------------------------------------------------

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := config{
		pokeClient: pokeapi.NewClient(5*time.Second, 5*time.Second),
		previous:   nil,
		next:       nil,
	}

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

		command, exists := commandList[cleanedQuery[0]]
		if exists {
			err := command.callback(&c)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}
