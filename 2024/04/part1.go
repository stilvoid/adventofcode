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

	test := func(s string) bool {
		if s == "XMAS" || s == "SAMX" {
			count++
			return true
		}
		return false
	}

	for y, row := range rows {
		for x, _ := range row {
			if x < len(row)-3 {
				// Right
				test(row[x : x+4])
			}

			if y < len(rows)-3 {
				// Down
				down := make([]byte, 4)
				for i := 0; i < 4; i++ {
					down[i] = rows[y+i][x]
				}
				test(string(down))
			}

			if x < len(row)-3 && y < len(rows)-3 {
				// Diag down-right
				diag := make([]byte, 4)
				for i := 0; i < 4; i++ {
					diag[i] = rows[y+i][x+i]
				}
				test(string(diag))
			}

			if x >= 3 && y < len(rows)-3 {
				// Diag down-left
				diag := make([]byte, 4)
				for i := 0; i < 4; i++ {
					diag[i] = rows[y+i][x-i]
				}
				test(string(diag))
			}
		}
	}

	fmt.Println(count)
}
