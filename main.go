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
	file := strings.Split(string(content), "\n")

	charMap := map[rune][]string{}

	for i := 32; i < 126; i++ {
		start := (i-32)*9 + 1
		end := start + 8
		charMap[rune(i)] = file[start:end]
	}
	input := "hello"
	for z := 0; z < 8; z++ {
		for _, f := range input {

			fmt.Print(charMap[f][z])
		}
		fmt.Println()
	}
	for _, v := range charMap['h'] {
		fmt.Println(v)
	}
}
