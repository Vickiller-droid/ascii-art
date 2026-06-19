package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(banner string) []string {
	banner = strings.TrimSuffix(banner, ".txt")
	content, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("Error reading file:")
		os.Exit(1)
	}
	lines := strings.Split(string(content), "\n")
	if len(content) == 0 {
		fmt.Println("Error: Empty file")
		os.Exit(1)
	}

	if len(lines) != 855 {
		fmt.Println("Error: Incomplete or corrupted file")
		os.Exit(1)
	}
	return lines
}

