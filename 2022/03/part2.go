package main

import (
	"fmt"
	"os"
	"strings"
)

func findBadge(lines []string) byte {
	seen := make(map[byte]int)

	for _, line := range lines {
		thisLine := make(map[byte]bool)

		for _, b := range []byte(line) {
			thisLine[b] = true
		}

		for b := range thisLine {
			if _, ok := seen[b]; !ok {
				seen[b] = 0
			}
			seen[b]++

			if seen[b] == len(lines) {
				return b
			}
		}
	}

	panic("no badge found")
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
	input := strings.Split(string(f), "\n")

	badges := make([]byte, 0)

	for i := 0; i < len(input); {
		lines := make([]string, 3)
		for j := 0; j < 3; j++ {
			lines[j] = input[i]
			i++
		}
		badges = append(badges, findBadge(lines))
	}

	fmt.Println(string(badges))

	fmt.Println(score(badges))
}
