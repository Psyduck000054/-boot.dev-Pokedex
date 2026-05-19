package main

import (
	"fmt"
	"math/rand"
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

func commandExplore(u *config, area string) error {
	pokeSlice, err := u.pokeClient.ListPokemons(area)
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

// ---------------------------------------------------------
// CATCH
// ---------------------------------------------------------

func commandCatch(c *config, p string) error {
	pokemon, err := c.pokeClient.PropertiesRetrieval(p)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", p)

	// BaseXP of a Pokemon runs from 36-300
	// The higher the BaseXP, the harder it is to catch
	// Min catch rate i guess is 25%?

	// f(36) = 1, f(300) = 0.25
	// f(x) = ax + b

	var catchRate float32
	catchRate = float32(-1.0/352.0)*float32(pokemon.BaseXP) + float32(97.0/88.0)

	fmt.Printf("Base XP: %d\n", pokemon.BaseXP)
	fmt.Printf("[Catch rate: %f]\n", catchRate)

	temp0 := rand.Float32()

	if catchRate >= temp0 {
		// caught
		fmt.Printf("%s was caught!\n", p)
	} else {
		// escape
		fmt.Printf("%s escaped!\n", p)
	}

	return nil
}
