package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type smap struct {
	width, height int
	pos           []row
}

type row []bool

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
			if p {
				out.WriteString("#")
			} else {
				out.WriteString(".")
			}
		}
		out.WriteString("\n")
	}

	return out.String()
}

func (sm smap) countSee(fromX, fromY int) int {
	angles := make(map[float64]bool)

	for y, row := range sm.pos {
		for x, p := range row {
			if p && (y != fromY || x != fromX) {
				angle := math.Atan2(float64(fromY-y), float64(fromX-x))
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

	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				sm.pos[y][x] = true
			}
		}
	}

	fmt.Println(sm.String())

	var best *int
	var bestX, bestY int
	for y, row := range sm.pos {
		for x, p := range row {
			if p {
				count := sm.countSee(x, y)
				fmt.Print(count)

				if best == nil || count > *best {
					best = &count
					bestX = x
					bestY = y
				}
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println(bestX, bestY, *best)
}
