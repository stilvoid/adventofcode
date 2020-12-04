package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type claim struct {
	id string
	x  int
	y  int
	w  int
	h  int
}

func main() {
	data, _ := ioutil.ReadFile("input")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	claims := make([]claim, 0)

	maxX, maxY := 0, 0

	// Load the claims
	for _, line := range lines {
		var c claim
		fmt.Sscanf(line[1:], "%s @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.w, &c.h)

		if c.x+c.w > maxX {
			maxX = c.x + c.w
		}

		if c.y+c.h > maxY {
			maxY = c.y + c.y
		}

		claims = append(claims, c)
	}

	// Find the one with no overlaps
	for i, c1 := range claims {
		overlaps := 0
		for j, c2 := range claims {

			if i != j {
				if c1.x >= c2.x+c2.w || c2.x >= c1.x+c1.w {
					continue
				}

				if c1.y >= c2.y+c2.h || c2.y >= c1.y+c1.h {
					continue
				}

				overlaps++
			}
		}
		if overlaps == 0 {
			fmt.Println(c1.id)
			return
		}
	}
}
