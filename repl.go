package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	temp0 := strings.ToLower(text)
	temp1 := strings.TrimSpace(temp0)

	var fArray []string

	fArray = strings.Fields(temp1)

	return fArray
}

// start the repl system
func replInit(c config) {
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
			fmt.Println("invalid input")
			continue
		}

		command, exists := commandList[cleanedQuery[0]]
		if exists {
			if command.name != "explore" && command.name != "catch" {
				err := command.callback(&c, "")
				if err != nil {
					fmt.Println(err)
				}
			} else {
				// explore take 2 params so we have to make sure the input has a length of at least 2
				if len(cleanedQuery) < 2 {
					fmt.Println("invalid input")
					continue
				}

				area := cleanedQuery[1]
				err := command.callback(&c, area)
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
