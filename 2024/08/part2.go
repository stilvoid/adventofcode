package main

import (
	"fmt"
	"math"
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
					antinodes[pos{locs[i].x, locs[i].y}] = true
					//theMap[locs[i].y][locs[i].x] = '#'

					antinodes[pos{locs[j].x, locs[j].y}] = true
					//theMap[locs[j].y][locs[j].x] = '#'

					dx := locs[i].x - locs[j].x
					dy := locs[i].y - locs[j].y

					a := math.Atan2(float64(dy), float64(dx))

					for y := 0; y < h; y++ {
						for x := 0; x < w; x++ {
							dx1 := locs[i].x - x
							dy1 := locs[i].y - y
							a1 := math.Atan2(float64(dy1), float64(dx1))

							dx2 := x - locs[i].x
							dy2 := y - locs[i].y
							a2 := math.Atan2(float64(dy2), float64(dx2))

							if a == a1 || a == a2 {
								antinodes[pos{x, y}] = true
								theMap[y][x] = '#'
							}
						}
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
