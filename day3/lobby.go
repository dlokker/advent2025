package main

import (
	"fmt"
	"strconv"
)

func lobbyP1(input []string, expected int) {
	var res, total int
	for _, line := range input {
		for i := range line {
			out := findMaxPair(i, line)
			if out > res {
				res = out
			}
		}
		// fmt.Printf("Max for line %s: %d\n", line, res)
		total += res
		res = 0
	}

	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("  Part 1: %v\n", total)
}

func findMaxPair(idx int, line string) int {
	cur := string(line[idx])
	var maxPair int
	for j := idx + 1; j < len(line); j++ {
		combined := cur + string(line[j])
		intVal, _ := strconv.Atoi(combined)
		if intVal > maxPair {
			// fmt.Printf("Found new max pair at index %d: %s -> %d\n", idx, combined, intVal)
			maxPair = intVal
		}
	}
	return maxPair
}

func lobbyP2(input []string, expected int) {
	var res, total int
	for _, line := range input {
		// fmt.Printf("%s\n", line)
		for i := range line {
			out := findMaxDozen(i, 12, line)
			outInt, _ := strconv.Atoi(out)
			if outInt > res {
				res = outInt
			}
		}
		// fmt.Printf("Max for line %s: %d\n", line, res)
		total += res
		res = 0
	}

	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("  Part 2: %v\n", total)
}

func findMaxDozen(start, digits int, line string) string {
	if digits == 0 || start >= len(line) {
		return ""
	}
	// Find largest number with at least 11 digits to the right of it.
	var max int32
	var idx int
	for i := start; i <= len(line)-digits; i++ {
		if rune(line[i]) > max {
			max = rune(line[i])
			idx = i
		}
	}
	// fmt.Printf("Max for place %d: %c\n", digits, max)
	return string(max) + findMaxDozen(idx+1, digits-1, line)	
}
