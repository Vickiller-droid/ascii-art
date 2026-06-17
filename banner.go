package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(banner string) []string {
	content, err := os.ReadFile(banner + ".txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	return strings.Split(string(content), "\n")
}
