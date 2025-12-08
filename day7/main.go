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

func laboratories1(input [][]string, expected int) {
	tachyonManifolds := make(map[int]bool)
	var result int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "S" {
				tachyonManifolds[j] = true
			}
			if input[i][j] == "^" {
				if tachyonManifolds[j] {
					result++
					if j > 0 {
						tachyonManifolds[j-1] = true
					}
					if j < len(input[i])-1 {
						tachyonManifolds[j+1] = true
					}
					tachyonManifolds[j] = false
				}
			}
		}
	}
	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("      P1: %d\n", result)
}

func laboratories2(input [][]string, expected int) {
	var result int
	colValues := make([]int, len(input[0]))
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "S" {
				colValues[j] = 1
			}
			if input[i][j] == "^" {
				if colValues[j] > 0 {
					if j > 0 {
						colValues[j-1] += colValues[j]
					}
					if j < len(input[i])-1 {
						colValues[j+1] += colValues[j]
					}
					colValues[j] = 0
				}
			}
		}
	}
	for _, v := range colValues {
		result += v
	}
	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("      P2: %d\n", result)
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
		var input [][]string
		var expected int
		lines := strings.Split(string(file), "\n")
		for _, line := range lines {
			if len(line) >= 9 && line[:9] == "expected:" {
				fmt.Sscanf(line, "expected: %d", &expected)
				continue
			}
			input = append(input, strings.Split(line, ""))
		}

		// Run stuff on the input
		fmt.Printf("=== Test Case: %s ===\n", filename)
		if strings.Contains(filename, "p1") {
			laboratories1(input, expected)
		} 
		if strings.Contains(filename, "p2") {
			laboratories2(input, expected)
		}
	}
}
