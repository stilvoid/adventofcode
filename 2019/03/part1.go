package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

func (w wire) plot(g *grid) {
	x := 0
	y := 0

	for _, m := range w {
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

			g.plot(x, y)
		}
	}
}

type grid map[int]map[int]int

func (g grid) plot(x, y int) {
	if _, ok := g[y]; !ok {
		g[y] = make(map[int]int)
	}
	g[y][x]++
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	wires := make([]wire, 0)

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

		wires = append(wires, w)
	}
	fmt.Println(wires)

	// Plot the wires on the grid
	grid := make(grid)
	for _, w := range wires {
		w.plot(&grid)
	}
	fmt.Println(grid)

	// Find intersections
	var min *float64
	for y, xs := range grid {
		for x, count := range xs {
			if count > 1 {
				dist := math.Abs(float64(x)) + math.Abs(float64(y))
				fmt.Println(dist)

				if min == nil || dist < *min {
					min = &dist
				}
			}
		}
	}

	fmt.Println(*min)
}
