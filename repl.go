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
