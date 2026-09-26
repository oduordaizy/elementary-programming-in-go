package main

import (
	"fmt"
	"os"
)

func main() {
	userArguments := os.Args[1:]

	if len(userArguments) != 1 {
		return
	}

	words := []string{}

	newstring := ""
	for _, val := range userArguments[0] {
		if val == ' ' {
			if newstring != "" {
				words = append(words, newstring)
				newstring = ""
			}
			continue
		}
		newstring += string(val)
	}

	for idx, word := range words {
		fmt.Print(word)
		if idx != len(words)-1 {
			fmt.Print("   ")
		}
	}
	fmt.Println()
}
