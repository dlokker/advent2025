package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	testsDir := "tests"

	// Read all files in the tests directory
	files, err := os.ReadDir(testsDir)
	if err != nil {
		fmt.Printf("Failed to read tests directory: %v\n", err)
		return
	}

	// Process each file
	for _, fileInfo := range files {
		filename := fileInfo.Name()
		filepath := filepath.Join(testsDir, filename)

		fmt.Printf("\n=== Test Case: %s ===\n", filename)

		// Read the file
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Printf("Failed to open file %s: %v\n", filename, err)
			continue
		}

		scanner := bufio.NewScanner(file)
		var input []string
		var expected int
		for scanner.Scan() {
			if scanner.Text()[:9] == "expected:" {
				fmt.Sscanf(scanner.Text(), "expected: %d", &expected)
				continue
			}
			input = append(input, scanner.Text())
		}
		file.Close()

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error occurred during scanning %s: %v\n", filename, err)
			continue
		}

		// Run giftShop on the input
		if strings.Contains(filename, "p1") {
			lobbyP1(input, expected)
		} 
		if strings.Contains(filename, "p2") {
			lobbyP2(input, expected)
		}
	}
}
