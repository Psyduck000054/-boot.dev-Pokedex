package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXIT
// ---------------------------------------------------------

func commandExit(c *config, useless0 string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

// ---------------------------------------------------------
// HELP
// ---------------------------------------------------------

func commandHelp(c *config, useless0 string) error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

#MISCELLANEOUS
help: Displays a help message
exit: Exit the Pokedex

#MAP
map: Shows the next page of in-game locations
mapb: Shows the previous page of in-game locations

#EXPLORE
explore [area]: Shows all possible Pokemon encounters in the area

`)
	return nil
}

// ---------------------------------------------------------
// MAP
// ---------------------------------------------------------

func commandMap(c *config, useless0 string) error {
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

func commandMapb(c *config, useless0 string) error {
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
// EXPLORE
// ---------------------------------------------------------

func commandExplore(c *config, area string) error {
	pokeSlice, err := c.pokeClient.ListPokemons(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s ...\n", area)
	fmt.Println("Found Pokemon:")

	for index, pokemon := range pokeSlice.Results {
		fmt.Printf("%d | %s\n", index+1, pokemon.Container.Name)
	}

	return nil
}
