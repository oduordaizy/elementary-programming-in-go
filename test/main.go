package main

import (
    "fmt"

    "piscine"
)

func main() {
    fmt.Print(piscine.FirstWord("hello there"))
    fmt.Print(piscine.FirstWord(""))
    fmt.Print(piscine.FirstWord("hello   .........  bye"))
}

func FisrtWord(s string) string {

	if s == "" {
		return "\n"
	}

	for i = 0; i < len(s); i++ {
	if s[i] == " "{
		return s[:i] + "\n"
	}

}