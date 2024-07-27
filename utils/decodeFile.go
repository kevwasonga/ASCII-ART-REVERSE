package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DecodeFile(filename string, artToChar map[string]string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	artHeight := 8
	var decodedString strings.Builder
	lineLength := len(lines[0])
	for i := 0; i < lineLength; {
		maxWidth := 0
		for j := 0; j < artHeight; j++ {
			if i < len(lines[j]) {
				if maxWidth < len(lines[j][i:]) {
					maxWidth = len(lines[j][i:])
				}
			}
		}
		found := false
		for width := 1; width <= maxWidth; width++ {
			subArt := []string{}
			for j := 0; j < artHeight; j++ {
				if i+width <= len(lines[j]) {
					subArt = append(subArt, lines[j][i:i+width])
				} else {
					subArt = append(subArt, lines[j][i:])
				}
			}
			charArt := strings.Join(subArt, "\n")
			if val, exists := artToChar[charArt]; exists {
				decodedString.WriteString(val)
				i += width
				found = true
				break
			}
		}
		if !found {
			i++
		}
	}

	return decodedString.String()
}
