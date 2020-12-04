package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type dir string

const (
	U dir = "U"
	D dir = "D"
	L dir = "L"
	R dir = "R"
)

type move struct {
	dir dir
	len int
}

type wire []move

func (w *wire) plot(g *grid) {
	x := 0
	y := 0
	steps := 0

	for _, m := range *w {
		for i := 0; i < m.len; i++ {

			switch m.dir {
			case U:
				y++
			case D:
				y--
			case L:
				x--
			case R:
				x++
			}

			steps++

			g.plot(x, y, w, steps)
		}
	}
}

type point map[*wire]int

type grid map[int]map[int]point

func (g grid) plot(x, y int, w *wire, steps int) {
	if _, ok := g[y]; !ok {
		g[y] = make(map[int]point)
	}

	if _, ok := g[y][x]; !ok {
		g[y][x] = make(point)
	}

	if g[y][x][w] == 0 {
		g[y][x][w] = steps
	}
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	wires := make([]*wire, 0)

	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}

		w := make(wire, 0)

		for _, instruction := range strings.Split(line, ",") {
			d := dir(instruction[0])

			l, err := strconv.Atoi(instruction[1:])
			if err != nil {
				panic(err)
			}

			w = append(w, move{
				dir: d,
				len: l,
			})
		}

		wires = append(wires, &w)
	}
	fmt.Println(wires)

	// Plot the wires on the grid
	grid := make(grid)
	for _, w := range wires {
		w.plot(&grid)
	}

	// Find intersections
	var min *int
	for _, xs := range grid {
		for _, point := range xs {
			if len(point) > 1 {
				dist := 0

				for _, steps := range point {
					dist += steps
				}

				if min == nil || dist < *min {
					min = &dist
				}
			}
		}
	}

	fmt.Println(*min)
}
