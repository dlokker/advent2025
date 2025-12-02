// Build: go build secret_entrance.go
// Run: ./secret_entrance
package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

		fmt.Printf("\n=== Processing %s ===\n", filename)

		// Read the file
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Printf("Failed to open file %s: %v\n", filename, err)
			continue
		}

		scanner := bufio.NewScanner(file)
		var input []string
		for scanner.Scan() {
			input = append(input, scanner.Text())
		}
		file.Close()

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error occurred during scanning %s: %v\n", filename, err)
			continue
		}

		// Run secretEntrance on the input
		secretEntrance(input)
	}
}

func secretEntrance(input []string) {
	var c1, c2 int
	cur := 50
	for _, i := range input {
		dir := i[:1]
		rotate, _ := strconv.Atoi(i[1:])
		if dir == "R" {
			c2 += numPasses(cur, rotate)
			cur += rotate
		}
		if dir == "L" {
			c2 += numPasses(cur, -rotate)
			cur -= rotate
		}
		if cur%100 == 0 {
			c1++
		}
		cur = cur % 100
	}
	fmt.Printf("Part 1: %v\n", c1)
	fmt.Printf("Part 2: %v\n", c2)
}

func numPasses(start, rotate int) int {
	if start <= 0 && rotate < 0 {
		if start+rotate <= -100 {
			return abs((start + rotate) / 100)
		}
	}
	if start < 0 && rotate > 0 {
		if start+rotate >= 0 {
			return 1 + abs((start+rotate)/100)
		}
	}
	if start >= 0 && rotate > 0 {
		if start+rotate >= 100 {
			return abs((start + rotate) / 100)
		}
	}
	if start > 0 && rotate < 0 {
		if start+rotate <= 0 {
			return 1 + abs((start+rotate)/100)
		}
	}
	return 0
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
