package main

import "strings"

func cleanInput(text string) []string {
	temp0 := strings.ToLower(text)
	temp1 := strings.TrimSpace(temp0)

	var fArray []string

	fArray = strings.Fields(temp1)

	return fArray
}
