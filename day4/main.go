package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	runInput = true
)

func checkAdjacent(i, j int, input [][]byte) bool {
	directions := [][2]int{
		{-1, 0},  // Up
		{1, 0},   // Down
		{0, -1},  // Left
		{0, 1},   // Right
		{-1, 1},  // north east
		{-1, -1}, // north west
		{1, 1},   // south east
		{1, -1},  // south west
	}
	var count int
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if ni >= 0 && ni < len(input) && nj >= 0 && nj < len(input[0]) {
			if input[ni][nj] == '@' {
				count++
			}
		}
	}
	return count < 4
}

func printingDepartment2(input [][]byte, expected int) {
	var res, total int
	var purge [][]int
	for {
		for i, row := range input {
			// fmt.Printf("%s\n", row)
			for j, cell := range row {
				if cell == '@' {
					if checkAdjacent(i, j, input) {
						res++
						purge = append(purge, []int{i, j})
					}
				}
			}
		}
		for _, p := range purge {
			input[p[0]][p[1]] = '.'
		}
		purge = [][]int{}
		if res == 0 {
			break
		}
		total += res
		res = 0
	}

	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("  Part 2: %d\n", total)
}

func printingDepartment1(input [][]byte, expected int) {
	var res int
	for i, row := range input {
		// fmt.Printf("%s\n", row)
		for j, cell := range row {
			if cell == '@' {
				if checkAdjacent(i, j, input) {
					res++
				}
			}
		}
	}

	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("  Part 1: %d\n", res)
}




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

		if !runInput && filename == "p1p2input" {
			continue
		}

		fmt.Printf("\n=== Test Case: %s ===\n", filename)

		// Read the file
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Printf("Failed to open file %s: %v\n", filename, err)
			continue
		}

		scanner := bufio.NewScanner(file)
		var input [][]byte
		var expected int
		for scanner.Scan() {
			if scanner.Text()[:9] == "expected:" {
				fmt.Sscanf(scanner.Text(), "expected: %d", &expected)
				continue
			}
			input = append(input, []byte(scanner.Text()))
		}
		file.Close()

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error occurred during scanning %s: %v\n", filename, err)
			continue
		}

		// Run stuff on the input
		if strings.Contains(filename, "p1") {
			printingDepartment1(input, expected)
		} 
		if strings.Contains(filename, "p2") {
			printingDepartment2(input, expected)
		}
	}
}
