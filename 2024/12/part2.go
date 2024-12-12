package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
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

type graph map[pos]rune

func (g graph) String() string {
	w, h := 0, 0

	for loc := range g {
		if loc.x > w {
			w = loc.x
		}

		if loc.y > h {
			h = loc.y
		}
	}

	out := strings.Builder{}

	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			r, ok := g[pos{x, y}]
			if !ok {
				out.WriteRune('o')
			} else {
				out.WriteRune(r)
			}
		}
		out.WriteRune('\n')
	}

	return out.String()
}

var up = pos{0, -1}
var down = pos{0, 1}
var left = pos{-1, 0}
var right = pos{1, 0}

var dirs = []pos{up, down, left, right}

func findContigs(layout graph) []graph {
	out := make([]graph, 0)

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

			contig := make(graph)
			for loc, include := range seen {
				if include {
					contig[loc] = layout[loc]
				}
			}

			if len(contig) > 0 {
				out = append(out, contig)
			}
		}
	}

	return out
}

func getEdges(kind rune, dir pos, layout graph) int {
	out := make(graph)

	for loc := range layout {
		next := loc.add(dir)
		if layout[loc] == kind && layout[next] != kind {
			out[loc] = kind
		} else {
			out[loc] = ' '
		}
	}

	total := 0
	for _, plot := range findContigs(out) {
		if plot[slices.Collect(maps.Keys(plot))[0]] == kind {
			total += 1
		}
	}

	return total
}

func getAllEdges(kind rune, layout graph) int {
	total := 0

	for _, dir := range dirs {
		total += getEdges(kind, dir, layout)
	}

	return total
}

func main() {
	data, _ := os.ReadFile("data")

	layout := make(graph)

	for y, row := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		for x, kind := range row {
			loc := pos{x, y}

			layout[loc] = kind
		}
	}

	fmt.Println(layout)
	fmt.Println()

	plots := findContigs(layout)

	total := 0
	for _, plot := range plots {
		fmt.Println(plot)
		fmt.Println()

		edges := getAllEdges(slices.Collect(maps.Values(plot))[0], plot)

		fmt.Println(edges)
		total += len(plot) * edges
	}

	fmt.Println(total)
}
