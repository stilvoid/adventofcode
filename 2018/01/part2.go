package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	total := 0
	var n int

	reached := make(map[int]bool)

	for {
		for _, line := range lines {
			reached[total] = true

			fmt.Sscanf(line[1:], "%d", &n)

			if line[0] == '+' {
				total += n
			} else {
				total -= n
			}

			if reached[total] {
				fmt.Println(total)
				return
			}
		}
	}
}
