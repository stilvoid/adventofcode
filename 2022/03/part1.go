package main

import (
	"fmt"
	"os"
	"strings"
)

func findDups(line string) []byte {
	in := []byte(line)

	seen := make(map[byte]bool)
	dups := make(map[byte]bool)
	out := make([]byte, 0)

	for i, b := range in {
		if i < len(in)/2 {
			seen[b] = true
		} else if _, ok := seen[b]; ok {
			dups[b] = true
		}
	}

	for b := range dups {
		out = append(out, b)
	}

	return out
}

func score(in []byte) int {
	out := 0
	for _, b := range in {
		if b >= 'a' && b <= 'z' {
			out += int(b - 'a' + 1)
		} else if b >= 'A' && b <= 'Z' {
			out += int(b - 'A' + 27)
		} else {
			panic(fmt.Sprintf("invalid byte: %v", b))
		}
	}

	return out
}

func main() {
	f, _ := os.ReadFile("input")

	dups := make([]byte, 0)

	for _, line := range strings.Split(string(f), "\n") {
		dups = append(dups, findDups(line)...)
	}

	fmt.Println(string(dups))

	fmt.Println(score(dups))
}
