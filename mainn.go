package main 

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	pieces := strings.Split(string(content), "\n\n")
	// fmt.Println(pieces[40])
	input := os.Args[1]
	// fmt.Println(ascii)
	// fmt.Println(string(content))
	// fmt.Println("Total pieces:", len(pieces))

	if len(os.Args) != 2 {
		fmt.Println("Usage : go run. \"Hello\"")
		return
	}
	// input := os.Args[1]

	content, err := os.ReadFile("standard.txt")
	for _, char := range input {
		fmt.Println(char - 32)
	}
}