package main

import (
	"fmt"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

func main() {
	data, _ := os.ReadFile("data")

	antennas := make(map[rune][]pos)

	var w, h int

	theMap := make([][]rune, 0)

	for y, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		row := make([]rune, 0)

		for x, char := range line {
			if char != '.' {
				if _, ok := antennas[char]; !ok {
					antennas[char] = make([]pos, 0)
				}

				antennas[char] = append(antennas[char], pos{x, y})
			}
			w = x + 1
			row = append(row, char)
		}
		h = y + 1
		theMap = append(theMap, row)
	}

	fmt.Println(antennas)

	antinodes := make(map[pos]bool)

	for _, locs := range antennas {
		if len(locs) > 1 {
			for i := 0; i < len(locs); i++ {
				for j := i + 1; j < len(locs); j++ {
					dx := locs[i].x - locs[j].x
					dy := locs[i].y - locs[j].y

					left := pos{x: locs[i].x + dx, y: locs[i].y + dy}
					right := pos{x: locs[j].x - dx, y: locs[j].y - dy}

					if left.x >= 0 && left.x < w && left.y >= 0 && left.y < h {
						antinodes[left] = true
						theMap[left.y][left.x] = '#'
					}

					if right.x >= 0 && right.x < w && right.y >= 0 && right.y < h {
						antinodes[right] = true
						theMap[right.y][right.x] = '#'
					}
				}
			}
		}
	}

	for _, row := range theMap {
		fmt.Println(string(row))
	}

	fmt.Println(len(antinodes))
}
