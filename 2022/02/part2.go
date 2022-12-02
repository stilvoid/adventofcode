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
}

func scoreGame(line string) int {
	parts := strings.Fields(line)

	left := shapes[parts[0]]

	if parts[1] == "X" {
		if left == 1 {
			return 3
		}
		return left - 1
	} else if parts[1] == "Y" {
		return 3 + left
	} else {
		if left == 3 {
			return 7
		}
		return 6 + left + 1
	}
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
