package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
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

		fmt.Printf("\n=== Test Case: %s ===\n", filename)

		// Read the file
		file, err := os.Open(filepath)
		if err != nil {
			fmt.Printf("Failed to open file %s: %v\n", filename, err)
			continue
		}

		scanner := bufio.NewScanner(file)
		var input []Range
		var expected int
		for scanner.Scan() {
			if scanner.Text()[:9] == "expected:" {
				fmt.Sscanf(scanner.Text(), "expected: %d", &expected)
				continue
			}
			line := scanner.Text()
			ranges := strings.Split(line, ",")
			for _, rg := range ranges {
				var r Range
				fmt.Sscanf(rg, "%d-%d", &r.start, &r.end)
				input = append(input, r)
			}
		}
		file.Close()

		if err := scanner.Err(); err != nil {
			fmt.Printf("Error occurred during scanning %s: %v\n", filename, err)
			continue
		}

		// Run giftShop on the input
		giftShop(input, expected)
	}
}

func giftShop(ranges []Range, expected int) {
	var res int
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			s := strconv.Itoa(i)
			if len(s)%2 != 0 {
				continue
			}
			mid := len(s) / 2
			if s[:mid] == s[mid:] {
				res += i
			}
		}
	}

	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("     Got: %d\n", res)
}