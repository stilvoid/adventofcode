package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

const (
	SPACE = '.'
	ROID  = '#'
	BASE  = 'X'
)

type point struct {
	x, y int
}

type smap struct {
	width, height int
	pos           []row
}

type row []rune

func newSmap(w, h int) smap {
	out := smap{
		width:  w,
		height: h,
		pos:    make([]row, h),
	}

	for i, _ := range out.pos {
		out.pos[i] = make(row, w)
	}

	return out
}

func (sm smap) String() string {
	out := strings.Builder{}

	for _, row := range sm.pos {
		for _, p := range row {
			out.WriteRune(p)
		}
		out.WriteString("\n")
	}

	return out.String()
}

func (sm smap) countSee(from point) int {
	angles := make(map[float64]bool)

	for y, row := range sm.pos {
		for x, p := range row {
			if p == ROID && (y != from.y || x != from.x) {
				angle := math.Atan2(float64(from.y-y), float64(from.x-x))
				angles[angle] = true
			}
		}
	}

	return len(angles)
}

func main() {
	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	sm := newSmap(len(lines[0]), len(lines))

	var base point

	for y, line := range lines {
		for x, c := range line {
			sm.pos[y][x] = c

			if c == BASE {
				base = point{x, y}
			}
		}
	}

	targets := make(map[float64]map[float64]point)

	for y, row := range sm.pos {
		for x, p := range row {
			if p == ROID {
				dx := float64(x - base.x)
				dy := float64(y - base.y)

				angle := math.Atan2(dy, dx) + math.Pi/2
				dist := math.Hypot(dx, dy)

				for angle < 0 {
					angle += math.Pi * 2
				}

				if _, ok := targets[angle]; !ok {
					targets[angle] = make(map[float64]point)
				}

				targets[angle][dist] = point{x, y}
			}
		}
	}

	fmt.Println("Base:", base)

	count := 0
	for len(targets) > 0 {
		angles := make([]float64, 0)
		for angle, _ := range targets {
			angles = append(angles, angle)
		}
		sort.Float64s(angles)

		for _, angle := range angles {
			dists := make([]float64, 0)
			for dist, _ := range targets[angle] {
				dists = append(dists, dist)
			}
			sort.Float64s(dists)

			if len(dists) == 0 {
				delete(targets, angle)
				continue
			}

			count++
			dead := targets[angle][dists[0]]

			fmt.Println(count, dead)

			delete(targets[angle], dists[0])
		}
	}
}
