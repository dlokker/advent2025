package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	runInput = false
)

func trashCompactor1(input [][]string, expected int) {
	ops := make([][]string, len(input[len(input)-1]))
	for i, op := range input[len(input)-1] {
		ops[i] = make([]string, len(input)-2)
		ops[i][0] = strings.TrimSpace(op)
	}
	for row := 0; row < len(input)-1; row++ {
		for col := 0; col < len(input[row]); col++ {
			ops[col] = append(ops[col], input[row][col])
		}
	}
	var result int
	for _, op := range ops {
		var total int
		first := true
		for i := 1; i < len(op); i++ {
			if op[i] == "" {
				continue
			}
			num, _ := strconv.Atoi(op[i])
			if op[0] == "*" {
				if first {
					total = 1
					first = false
				}
				total = total * num
			}
			if op[0] == "+" {
				total = total + num
			}
		}
		fmt.Printf("  Total for op %v: %d\n", op, total)
		result += total
	}
	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("      P1: %d\n", result)
}

func trashCompactor2(input [][]string, expected int) {
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
			l := strings.Split(line, " ")
			var clean []string
			for i := range l {
				if strings.TrimSpace(l[i]) == "" {
					continue
				}
				clean = append(clean, strings.TrimSpace(l[i]))
			}
			input = append(input, clean)
		}

		// Run stuff on the input
		fmt.Printf("=== Test Case: %s ===\n", filename)
		if strings.Contains(filename, "p1") {
			trashCompactor1(input, expected)
		} 
		if strings.Contains(filename, "p2") {
			trashCompactor2(input, expected)
		}
	}
}
