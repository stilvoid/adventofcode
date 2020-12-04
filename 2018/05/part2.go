package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input")
	line := []byte(strings.TrimSpace(string(data)))

	// Get the list of types
	types := make(map[string]bool)
	for _, c := range line {
		types[strings.ToLower(string(c))] = true
	}

	var a, b string
	shortestCount := len(line)

	originalLine := append(line)

	for t, _ := range types {
		line = append(originalLine)

		// Remove the offender
		line = []byte(strings.Replace(string(line), t, "", -1))
		line = []byte(strings.Replace(string(line), strings.ToUpper(t), "", -1))

		for i := 0; i < len(line); i++ {
			if i < 0 {
				i = 0
			}

			if i < len(line)-1 {
				a = string(line[i])
				b = string(line[i+1])

				if a != b && strings.ToLower(a) == strings.ToLower(b) {
					// Anihilate
					line = append(line[:i], line[i+2:]...)

					// Step back so we can check the previous character again
					i -= 2
				}
			}
		}

		if len(line) < shortestCount {
			shortestCount = len(line)
		}
	}

	fmt.Println(shortestCount)
}
