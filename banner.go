package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(banner string) []string {
	content, err := os.ReadFile(banner + ".txt")

	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}

	lines := strings.Split(string(content), "\n")

	if len(content) == 0 {
		fmt.Println("Error: Empty File")
		os.Exit(1)
	}

	if len(lines) != 855 {
		fmt.Println("Error: Incomplete or corrupt file")
		os.Exit(1)
	}
	return lines
}