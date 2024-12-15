package main

import (
	"fmt"
	"os"
	"strings"
)

type loc struct {
	x, y int
}

func (l loc) add(o loc) loc {
	return loc{
		l.x + o.x,
		l.y + o.y,
	}
}

var (
	up    = loc{0, -1}
	down  = loc{0, 1}
	left  = loc{-1, 0}
	right = loc{1, 0}
)

type world [][]rune

func newWorld(in string) world {
	out := make([][]rune, 0)

	for y, line := range strings.Split(in, "\n") {
		out = append(out, make([]rune, 0))

		for _, char := range line {
			out[y] = append(out[y], char)
		}
	}

	return out
}

func (w world) Print() {
	for _, row := range w {
		fmt.Println(string(row))
	}
}

func (w world) get(pos loc) rune {
	return w[pos.y][pos.x]
}

func (w world) set(pos loc, what rune) {
	w[pos.y][pos.x] = what
}

func (w world) move(what rune, start, dir loc) loc {
	next := start.add(dir)
	nextWhat := w.get(next)

	switch nextWhat {
	case '#':
		return start
	case 'O':
		if w.move(nextWhat, next, dir) != next {
			return w.move(what, start, dir)
		}
		return start
	case '.':
		w.set(start, '.')
		w.set(next, what)
		return next
	default:
		panic(fmt.Sprintf("Unexpected: %r", nextWhat))
	}
}

func main() {
	data, _ := os.ReadFile("data")

	parts := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	w := newWorld(parts[0])

	w.Print()

	// Find robot
	robot := loc{}

	for y, row := range w {
		for x, char := range row {
			if char == '@' {
				robot.x = x
				robot.y = y
				break
			}
		}
	}

	fmt.Println(robot)

	for _, inst := range parts[1] {
		switch inst {
		case '^':
			robot = w.move('@', robot, up)
		case 'v':
			robot = w.move('@', robot, down)
		case '<':
			robot = w.move('@', robot, left)
		case '>':
			robot = w.move('@', robot, right)
		}

		fmt.Println(robot)

		w.Print()
	}

	total := 0
	for y, row := range w {
		for x, char := range row {
			if char == 'O' {
				total += 100*y + x
			}
		}
	}

	fmt.Println(total)
}
