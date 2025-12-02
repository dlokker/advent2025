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
		if strings.Contains(filename, "p1") {
			giftShopP1(input, expected)
		} 
		if strings.Contains(filename, "p2") {
			giftShopP2(input, expected)
		}
	}
}

func giftShopP1(ranges []Range, expected int) {
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
	fmt.Printf("  P1 Got: %d\n", res)
}

func giftShopP2(ranges []Range, expected int) {
	var res int
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			s := strconv.Itoa(i)
			if len(s) == 1 {
				continue
			}
			// Build substrings up to half the length of s and check if that repeats
			for j := 1; j <= len(s) / 2; j++ {
				var found bool
				// If the length of substring does not divide evenly into length of s, skip
				if len(s)%len(s[:j]) != 0 {
					continue
				}
				for k := len(s[:j]); k < len(s); k += len(s[:j]) {
					if s[:j] != s[k:k+len(s[:j])] {
						break
					}
					if k+len(s[:j]) == len(s) {
						res += i
						found = true
						break
					}
				}
				if found {
					break
				}
			}
		}
	}

	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("  P2 Got: %d\n", res)
}