package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	runInput = true
)

func checkAdjacent(i, j int, input [][]byte) bool {
	directions := [][]int{
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
	var total int
	for {
		var res int
		var purge [][]int
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
		if res == 0 {
			break
		}
		total += res
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
	files, _ := os.ReadDir(testsDir)
	for _, fileInfo := range files {
		filename := fileInfo.Name()
		filepath := filepath.Join(testsDir, filename)
		if !runInput && filename == "p1p2input" {
			continue
		}
		file, _ := os.ReadFile(filepath)
		var input [][]byte
		var expected int
		lines := strings.Split(string(file), "\n")
		for _, line := range lines {
			if len(line) >= 9 && line[:9] == "expected:" {
				fmt.Sscanf(line, "expected: %d", &expected)
				continue
			}
			input = append(input, []byte(line))
		}

		// Run stuff on the input
		fmt.Printf("=== Test Case: %s ===\n", filename)
		if strings.Contains(filename, "p1") {
			printingDepartment1(input, expected)
		} 
		if strings.Contains(filename, "p2") {
			printingDepartment2(input, expected)
		}
	}
}
