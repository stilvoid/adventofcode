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
	}

	w, h := len(locs[0]), len(locs)

	// Find the guard
	var sx, sy int
	for i, line := range locs {
		for j, char := range line {
			if char == '^' {
				sx, sy = j, i
			}
		}
	}

	hasLoop := func(x, y, tx, ty int, ch chan bool) {
		seen := make(map[string]bool)

		var nx, ny int
		dx, dy := 0, -1
		for {
			id := fmt.Sprintf("%d,%d,%d,%d", x, y, dx, dy)

			if seen[id] {
				ch <- true
				return
			}

			seen[id] = true
			locs[y][x] = 'X'

			nx = x + dx
			ny = y + dy

			if nx < 0 || nx >= w || ny < 0 || ny >= h {
				break
			}

			if (nx == tx && ny == ty) || locs[ny][nx] == '#' {
				// Turn clockwise
				dx, dy = -dy, dx
			} else {
				x = nx
				y = ny
			}
		}

		ch <- false
	}

	ch := make(chan bool)
	for ty := 0; ty < h; ty++ {
		for tx := 0; tx < w; tx++ {
			go hasLoop(sx, sy, tx, ty, ch)
		}
	}

	found := 0
	count := 0
	for found < w*h {
		loop := <-ch
		found++
		if loop {
			count++
		}
	}
	fmt.Println(count)
}
