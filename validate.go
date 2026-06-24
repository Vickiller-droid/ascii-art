package main 

import (
	"fmt"
	"os"
	"strings"
)

func validate() (string, string) {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . \"Hello\"")
		os.Exit(1)
	}

	input := os.Args[1]
	banner := "standard" 

	if len(os.Args) == 3 {
		banner = os.Args[2]
	}

	banner = strings.TrimSuffix(banner, ".txt")

	if banner != "standard" &&
	banner != "shadow" &&
	banner != "thinkertoy" {
		fmt.Println("Invalid font: Use standard, shadow or thinkertoy with \".txt\", e.g standard.txt or standard alone")
		os.Exit(1)
	}
	return input, banner
}