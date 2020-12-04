package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	counts := make(map[int]int)

	for _, line := range lines {
		chars := make(map[rune]int)
		for _, char := range line {
			chars[char]++
		}

		lineCounts := make(map[int]bool)
		for _, count := range chars {
			lineCounts[count] = true
		}

		for count, _ := range lineCounts {
			counts[count]++
		}
	}

	fmt.Println(counts[2] * counts[3])
}
