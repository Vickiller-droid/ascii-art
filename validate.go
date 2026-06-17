package main

import (
	"fmt"
	"os"
)

func validate() (string, string){
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . \"Hello\"" )
		os.Exit(1)
	}

	input := os.Args[1]
	banner := "standard"

	if len(os.Args) == 3 {
		banner = os.Args[2]
	}

	return  input, banner
}
