package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("data")

	locs := make([][]byte, 0)

	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		locs = append(locs, []byte(line))
		fmt.Println(line)
	}

	w, h := len(locs[0]), len(locs)
	fmt.Println(w, h)

	// Find the guard
	var x, y int
	dx, dy := 0, -1
	for i, line := range locs {
		for j, char := range line {
			if char == '^' {
				x, y = j, i
			}
		}
	}

	var nx, ny int
	for {
		locs[y][x] = 'X'

		nx = x + dx
		ny = y + dy

		if nx < 0 || nx >= w || ny < 0 || ny >= h {
			break
		}

		/*
			0,-1 -> 1,0
			1,0 -> 0,1
			0,1 -> -1,0
			-1,0 -> 0,-1
		*/
		if locs[ny][nx] == '#' {
			// Turn clockwise
			dx, dy = -dy, dx
		} else {
			x = nx
			y = ny
		}
	}

	count := 0
	for _, line := range locs {
		fmt.Println(string(line))
		for _, char := range line {
			if char == 'X' {
				count++
			}
		}
	}

	fmt.Println(count)
}
