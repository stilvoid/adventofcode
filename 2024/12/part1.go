package main

import (
	"fmt"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

func (p pos) add(other pos) pos {
	return pos{
		p.x + other.x,
		p.y + other.y,
	}
}

var dirs = []pos{
	pos{0, -1},
	pos{0, 1},
	pos{-1, 0},
	pos{1, 0},
}

func findContigs(layout map[pos]rune) [][]pos {
	out := make([][]pos, 0)

	var check func(pos, map[pos]bool)

	done := make(map[pos]bool)

	check = func(start pos, seen map[pos]bool) {
		done[start] = true

		for _, dir := range dirs {
			next := start.add(dir)

			if kind, ok := layout[next]; ok {
				// Exists
				if _, ok := seen[next]; !ok {
					// Not checked

					if kind == layout[start] {
						seen[next] = true
						check(next, seen)
					} else {
						seen[next] = false
					}
				}
			}
		}
	}

	for loc := range layout {
		if !done[loc] {
			seen := map[pos]bool{
				loc: true,
			}

			check(loc, seen)

			contig := make([]pos, 0)
			for loc, include := range seen {
				if include {
					contig = append(contig, loc)
				}
			}

			if len(contig) > 0 {
				out = append(out, contig)
			}
		}
	}

	return out
}

func perim(layout map[pos]rune, area []pos) int {
	total := 0

	for _, loc := range area {
		for _, dir := range dirs {
			if layout[loc] != layout[loc.add(dir)] {
				total++
			}
		}
	}

	return total
}

func main() {
	data, _ := os.ReadFile("data")

	layout := make(map[pos]rune)

	for y, row := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		for x, kind := range row {
			loc := pos{x, y}

			layout[loc] = kind
		}
	}

	fmt.Println(layout)

	plots := findContigs(layout)

	fmt.Println(plots)

	total := 0
	for _, plot := range plots {
		a := len(plot)
		p := perim(layout, plot)

		total += a * p

		fmt.Printf("%d x %d = %d\n", a, p, a*p)
	}

	fmt.Println(total)
}
