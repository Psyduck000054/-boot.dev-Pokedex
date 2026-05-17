package main

import (
	"fmt"
	"os"
)

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
