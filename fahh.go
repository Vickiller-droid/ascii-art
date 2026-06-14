package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"Hello\"")
		return
	}

	input := os.Args[1]
	content, err := os.ReadFile("standard.txt")

	lines := strings.Split(string(content), "\n")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	words := strings.Split(input, "\\n")

	for _, word := range words {

		for i := 0; i < 8; i++ {
			for _, char := range word {
				fmt.Print(lines[(int(char)-32)*9+1+i])
			}
			fmt.Println()
		}
	}
}
