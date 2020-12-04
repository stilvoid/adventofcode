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

	// Find the overlaps
	overlaps := 0
	for x := 0; x <= maxX; x++ {
		for y := 0; y <= maxY; y++ {
			count := 0

			for _, c := range claims {
				if x >= c.x && x < c.x+c.w && y >= c.y && y < c.y+c.h {
					count++
				}
			}

			if count >= 2 {
				overlaps++
			}
		}
	}

	fmt.Println(overlaps)
}
