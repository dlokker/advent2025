package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	runInput = true
	multCircuits = 3
)

type coord struct {
	x, y, z int
}

type connection struct {
	a, b coord
	dist float64
}

type DistHeap []connection

func (h DistHeap) Len() int { return len(h) }
func (h DistHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h DistHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *DistHeap) Push(x any) {
	*h = append(*h, x.(connection))
}

func (h *DistHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func computeDistance(a, b coord) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func makeConnection(distHeap *DistHeap, circuits *[]map[coord]bool) connection {
	// Get the minimum distance connection
	minDist := heap.Pop(distHeap).(connection)
	aC, bC := -1, -1
	// Find which circuits the coords belong to
	for c := range *circuits {
		if _, ok := (*circuits)[c][minDist.a]; ok {			
			aC = c
		}
		if _, ok := (*circuits)[c][minDist.b]; ok {
			bC = c
		}
	}
	// Add the connection appropriately
	if aC != -1 && bC == -1 {
		(*circuits)[aC][minDist.b] = true
	}
	if aC == -1 && bC != -1 {
		(*circuits)[bC][minDist.a] = true
	}
	if aC != -1 && bC != -1 && aC != bC {
		for b := range (*circuits)[bC] {
			(*circuits)[aC][b] = true
		}
		*circuits = append((*circuits)[:bC], (*circuits)[bC+1:]...)
	}
	if aC == -1 && bC == -1 {
		newCircuit := make(map[coord]bool)
		newCircuit[minDist.a] = true
		newCircuit[minDist.b] = true
		*circuits = append((*circuits), newCircuit)
	}
	return minDist
}

func playground(input []coord, connections, expected int, part string) {
	distHeap := &DistHeap{}
	heap.Init(distHeap)
	var result int
	// Build a min-heap of all distances
	for i, c := range input {
		for j := i + 1; j < len(input); j++ {
			heap.Push(distHeap, connection{a: c, b: input[j], dist: computeDistance(c, input[j])})
		}
	}
	var circuits []map[coord]bool
	if part == "p1"	{
		for range(connections) {
			makeConnection(distHeap, &circuits)
		}
		result = 1
		for range(multCircuits) {
			max := 0
			maxI := -1
			for c := range circuits {
				if len(circuits[c]) > max {
					max = len(circuits[c])
					maxI = c
				}
			}
			result *= max
			circuits = append(circuits[:maxI], circuits[maxI+1:]...)
		}
	}
	if part == "p2" {
		minDist := makeConnection(distHeap, &circuits)
		for len(circuits[0]) != len(input) {
			minDist = makeConnection(distHeap, &circuits)
		}
		result = minDist.a.x * minDist.b.x
	}
	fmt.Printf("Expected(%s): %d, Got: %d\n", part, expected, result)
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
		var input []coord
		var expected, connections int
		lines := strings.Split(string(file), "\n")
		for _, line := range lines {
			if len(line) >= 9 && line[:9] == "expected:" {
				fmt.Sscanf(line, "expected: %d", &expected)
				continue
			}
			if len(line) > 12 && line[:12] == "connections:" {
				fmt.Sscanf(line, "connections: %d", &connections)
				continue
			}
			parts := strings.Split(line, ",")
			var c coord
			for i, part := range parts {
				num, _ := strconv.Atoi(strings.TrimSpace(part))
				switch i {
				case 0:
					c.x = num
				case 1:
					c.y = num
				case 2:
					c.z = num
				}
			}
			input = append(input, c)
		}

		// Run stuff on the input
		fmt.Printf("=== Test Case: %s ===\n", filename)
		if strings.Contains(filename, "p1") {
			playground(input, connections, expected, "p1")
		} 
		if strings.Contains(filename, "p2") {
			playground(input, connections, expected, "p2")
		}
	}
}
