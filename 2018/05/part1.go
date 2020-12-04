package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input")
	line := []byte(strings.TrimSpace(string(data)))

	var a, b string

	for i := 0; i < len(line); i++ {
		if i < 0 {
			i = 0
		}

		if i < len(line)-1 {
			a = string(line[i])
			b = string(line[i+1])

			// Kaboom
			if a != b && strings.ToLower(a) == strings.ToLower(b) {
				// Anihilate
				line = append(line[:i], line[i+2:]...)

				// Step back so we can check the previous character again
				i -= 2
			}
		}
	}

	fmt.Println(len(line))
}
