package main 

import(
	"fmt"
	"strings"
)

func generate(input string, lines []string) {
	words := strings.Split(input, "\\n")

	for _, word := range words {
		if strings.TrimSpace(word) == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8;i++ {
			for _, char := range word {
				fmt.Print(lines[(int(char)-32)*9+1+i])
			}
			fmt.Println()
		}
	}
}