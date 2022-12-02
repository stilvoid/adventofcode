package main

import (
	"fmt"
	"os"
	"strings"
)

var shapes = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func scoreGame(line string) int {
	parts := strings.Fields(line)

	left := shapes[parts[0]]
	right := shapes[parts[1]]

	if (left%3)+1 == right {
		return 6 + right
	} else if left == right {
		return 3 + right
	}

	return right
}

func main() {
	f, _ := os.ReadFile("input")

	total := 0

	for _, line := range strings.Split(string(f), "\n") {
		score := scoreGame(line)
		total += score
		fmt.Println(score)
	}

	fmt.Println(total)
}
