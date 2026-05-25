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
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	pieces := strings.Split(string(content), "\n\n")
	_ = pieces

	for _, char := range input {
		fmt.Println(pieces[char-32],"\n\n")
	}
}