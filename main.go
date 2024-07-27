package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-reverse/asciiArt"
	"ascii-art-reverse/utils"
)

func main() {
	// Check if an argument is provided
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Printf("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>\n")
		return
	}
	standardFile := "banners/standard.txt"
	utils.ParseFlag()
	sampleFile := *utils.ReversePtr
	// if no flag is provided, switch to accept printable ascii input and display their banner files in the terminal
	if sampleFile == "" {
		// If no flag is provided, use the Filename function to decide the file to be used
		sampleFile = asciiArt.BannerFile()
	}

	// special handling for new line , since we are now dealing with the normal  ascii
	// print a new line and exit in case argument is a new line character only
	if os.Args[1] == "\\n" {
		fmt.Println()
		return
	}

	// Load the banner map from the file
	bannerMap, err := asciiArt.LoadBannerMap(sampleFile)
	if err != nil {
		fmt.Println("error loading banner map:", err)
		return
	}

	// Process the provided argument
	args := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	args = strings.ReplaceAll(args, "\\t", "    ")
	lines := strings.Split(args, "\n")

	// Generate the ASCII art for each line
	for _, line := range lines {
		asciiArt.PrintLineBanner(line, bannerMap)
	}

	// Create a mapping from standard.txt
	artToChar := utils.ParseBannerFile((standardFile))

	// Decode the sample.txt using the mapping
	decodedString := utils.DecodeFile(sampleFile, artToChar)

	// Print the result
	fmt.Println(decodedString)
}
