package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile := flag.String("file", "", "input file to use")
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
	count := 0
	cur := 50
	for _, i := range input {
		dir := i[:1]
		rotate, _ := strconv.Atoi(i[1:])
		if dir == "R" {
			cur += rotate
		}
		if dir == "L" {
			cur -= rotate
		}
		if cur%100 == 0 {
			count++
		}
		cur = cur % 100
	}
	fmt.Printf("Password: %v\n", count)
}
