package main

import (
	"fmt"
	"os"
)

func countUnique(in []byte) int {
	counter := make(map[byte]bool)

	for _, b := range in {
		counter[b] = true
	}

	return len(counter)
}

func findMarker(in []byte, length int) int {
	for i := 0; i < len(in)-length; i++ {
		t := in[i : i+length]
		if countUnique(t) == length {
			return i + length
		}
	}

	panic("marker not found")
}

func main() {
	start, _ := os.ReadFile("input")

	fmt.Println("Part 1:", findMarker(start, 4))
	fmt.Println("Part 2:", findMarker(start, 14))
}
