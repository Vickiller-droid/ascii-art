package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Error: ")
	}
	banner := "standard.txt"
	banenrFile, _ := loadBanner(banner)
	if len(os.Args) == 3 {
		banenrFile = banner + ".txt"
	}
	inputFile := strings.Split(os.Args[1], `\n`)
	for i, input := range inputFile {
		if input == "" {
			if len(inputFile)-1 == i {
				continue
			}
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			for _, ch := range input {
				fmt.Print(banenrFile[ch][i])
			}
		}
	}
	
}