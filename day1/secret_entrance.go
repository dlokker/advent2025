// Build: go build secret_entrance.go
// Run: ./secret_entrance --file test3-6
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	inputFile = flag.String("file", "", "input file to use")
)

func main() {
	flag.Parse()
	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", *inputFile)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error occurred during scanning: %v\n", err)
	}
	secretEntrance(input)
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
