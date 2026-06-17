package main

import (
	"fmt"
	"os"
)

func validate() string {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"Hello\"")
		os.Exit(1)
	}
	return os.Args[1]
}
