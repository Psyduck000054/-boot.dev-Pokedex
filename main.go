package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type config struct {
	previous string
	next     string
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
	// 1. build URL
	var url string

	if c.next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = c.next
	}

	// 2. send GET request
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// 3. read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// 4. unmarshal json into go structs
	var mapRes mapResponse
	err = json.Unmarshal(body, &mapRes)

	// 5. extract the fields the command needs
	for _, item := range mapRes.Results {
		fmt.Println(item.Name)
	}

	// no c.previous case
	if mapRes.Previous != nil {
		c.previous = *mapRes.Previous
	} else {
		c.previous = ""
	}

	c.next = mapRes.Next

	return nil
}

func commandMapb(c *config) error {
	// 1. build URL
	var url string

	if c.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		url = c.previous
	}

	// 2. send GET request
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// 3. read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// 4. unmarshal json into go structs
	var mapRes mapResponse
	err = json.Unmarshal(body, &mapRes)

	// 5. extract the fields the command needs
	for _, item := range mapRes.Results {
		fmt.Println(item.Name)
	}

	// no c.previous case
	if mapRes.Previous != nil {
		c.previous = *mapRes.Previous
	} else {
		c.previous = ""
	}

	c.next = mapRes.Next

	return nil
}

type result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type mapResponse struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous *string  `json:"previous"`
	Results  []result `json:"results"`
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
	var c config

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
