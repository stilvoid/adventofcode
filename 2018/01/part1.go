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

	for _, line := range lines {
		fmt.Sscanf(line[1:], "%d", &n)

		if line[0] == '+' {
			total += n
		} else {
			total -= n
		}
	}

	fmt.Println(total)
}
