package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(filename string) (map[rune][]string, error) {
	content, err := os.ReadFile("standard.txt")
	if err != nil {
		return nil, fmt.Errorf("file not found")
	}

	if len(content) == 0 {
		return nil, fmt.Errorf("file not empty")
	}
	lines := strings.Split(string(content), "\n")
	mapList := make(map[rune][]string)
	for i := 0; i < 95; i++ {
		ch := rune(i+32)
		start := i*9+1
		w := lines[start : start+8]
		line := make([]string, 8)
		copy(line, w)
		mapList[ch]= line
	}
	return mapList, nil
}
