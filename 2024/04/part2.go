package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data")

	rows := strings.Split(strings.TrimSpace(string(data)), "\n")

	count := 0

	for y := 1; y < len(rows)-1; y++ {
		for x := 1; x < len(rows[0])-1; x++ {
			if rows[y][x] == 'A' {
				ms, ss := 0, 0

				test := func(c byte) {
					switch c {
					case 'M':
						ms++
					case 'S':
						ss++
					}
				}

				test(rows[y-1][x-1])
				test(rows[y-1][x+1])
				test(rows[y+1][x-1])
				test(rows[y+1][x+1])

				if ms == 2 && ss == 2 && rows[y-1][x-1] != rows[y+1][x+1] {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
