package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, _ := os.ReadFile("data")
	fmt.Println(string(data))

	mul := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	result := mul.FindAllStringSubmatch(string(data), -1)

	total := 0
	for _, part := range result {
		a, _ := strconv.Atoi(part[1])
		b, _ := strconv.Atoi(part[2])

		total += a * b
	}

	fmt.Println(total)
}
