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

	mul := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\))`)

	result := mul.FindAllStringSubmatch(string(data), -1)

	for _, part := range result {
		fmt.Printf("%#v\n", part)
	}

	total := 0
	enabled := true

	for _, part := range result {
		op := part[1]

		switch op {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				a, _ := strconv.Atoi(part[2])
				b, _ := strconv.Atoi(part[3])
				total += a * b
			}
		}
	}

	fmt.Println(total)
}
