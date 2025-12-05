package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	runInput = true
)

type Range struct {
	start int
	end   int
}

func cafeteria1(ranges []Range, available []int, expected int) {
	var count int
	for _, a := range available {
		// fmt.Printf("Checking %d against ranges:\n", a)
		for _, r := range ranges {
			// fmt.Printf("  Range: %d-%d\n", r.start, r.end)
			if a >= r.start && a <= r.end {
				// fmt.Printf("  %d is within range %d-%d\n", a, r.start, r.end)
				count++
				break
			}
		}
	}
	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("      P1: %d\n", count)
}

func combineRanges(rangehist *[]Range) {
	for i := range(len(*rangehist)) {
		for j := i + 1; j < len(*rangehist); j++ {
			ri := (*rangehist)[i]
			rj := (*rangehist)[j]
			if ri.start <= rj.end && ri.end >= rj.start {
				// They overlap, combine them
				fmt.Printf("  Combining ranges %d-%d and %d-%d\n", ri.start, ri.end, rj.start, rj.end)
				if rj.start < ri.start {
					(*rangehist)[i].start = rj.start
				}
				if rj.end > ri.end {
					(*rangehist)[i].end = rj.end
				}
				// Remove
				*rangehist = append((*rangehist)[:j], (*rangehist)[j+1:]...)
				j--
			}
		}
	}
}

func cafeteria2(ranges []Range, expected int) {
	var count int
	var rangehist []Range
	for _, r := range ranges {
		fmt.Printf("Processing range: %d-%d\n", r.start, r.end)
		count += (r.end - r.start + 1)
		addedToExisting := false
		for i, rh := range rangehist {
			if r.start > rh.start && r.start <= rh.end && r.end >= rh.end {
				overlap := rh.end - r.start + 1
				fmt.Printf("  rightlap with %d-%d: %d\n", rh.start, rh.end, overlap)
				rangehist[i].end = r.end
				addedToExisting = true
				count -= overlap
			} else if r.end >= rh.start && r.end < rh.end && r.start < rh.start {
				overlap := r.end - rh.start + 1
				fmt.Printf("  leftlap with %d-%d: %d\n", rh.start, rh.end, overlap)
				rangehist[i].start = r.start
				addedToExisting = true
				count -= overlap
			} else if r.start <= rh.start && r.end >= rh.end {
				overlap := rh.end - rh.start + 1
				fmt.Printf("  Overlap with %d-%d: %d\n", rh.start, rh.end, overlap)
				rangehist[i].start = r.start
				rangehist[i].end = r.end
				addedToExisting = true
				count -= overlap
			} else if r.start >= rh.start && r.end <= rh.end {
				overlap := r.end - r.start + 1
				fmt.Printf("  Innerlap with %d-%d: %d\n", rh.start, rh.end, overlap)
				addedToExisting = true
				count -= overlap
			}
		}
		if !addedToExisting {
			fmt.Printf("  No overlap, adding new range %d-%d\n", r.start, r.end)
			rangehist = append(rangehist, r)
		}
		combineRanges(&rangehist)
		fmt.Printf("count: %d\n", count)
	}
	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("      P2: %d\n", count)
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
		var ranges []Range
		var available []int
		var expected int
		lines := strings.Split(string(file), "\n")
		for _, line := range lines {
			if len(line) >= 9 && line[:9] == "expected:" {
				fmt.Sscanf(line, "expected: %d", &expected)
				continue
			}
			if strings.Contains(line, "-") {
				var r Range
				fmt.Sscanf(line, "%d-%d", &r.start, &r.end)
				ranges = append(ranges, r)
			}
			if line != "" {
				num, _ := strconv.Atoi(line)
				available = append(available, num)
			}
		}

		// Run stuff on the input
		fmt.Printf("=== Test Case: %s ===\n", filename)
		if strings.Contains(filename, "p1") {
			cafeteria1(ranges, available, expected)
		} 
		if strings.Contains(filename, "p2") {
			cafeteria2(ranges, expected)
		}
	}
}
