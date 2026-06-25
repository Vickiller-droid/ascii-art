package main

import (
	"strings"
	"fmt"
)
func generate(input string, lines []string) {
	if input == "" {
		return
	}

	if strings.Trim(input, "\\n") == "" {
		fmt.Print(strings.Repeat("\n", strings.Count(input, "\\n")))
		return
	}

	words := strings.Split(input, "\\n")

	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			for _, char := range word {
				fmt.Print(lines[(int(char)-32)*9+1+i])
			}
			fmt.Println()
		}
	}
}