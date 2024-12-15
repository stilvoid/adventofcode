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
			switch char {
			case '#':
				out[y] = append(out[y], '#', '#')
			case 'O':
				out[y] = append(out[y], '[', ']')
			case '.':
				out[y] = append(out[y], '.', '.')
			case '@':
				out[y] = append(out[y], '@', '.')
			default:
				panic(fmt.Sprint("Unexpected: %r", char))
			}
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

func (w world) check(start, dir loc) bool {
	next := start.add(dir)
	nextWhat := w.get(next)

	if nextWhat == '#' {
		return false
	}

	if nextWhat == '.' {
		return true
	}

	if nextWhat == '[' {
		if dir == left || dir == right {
			return w.check(next, dir)
		}

		return w.check(next, dir) && w.check(next.add(right), dir)
	}

	if nextWhat == ']' {
		if dir == left || dir == right {
			return w.check(next, dir)
		}

		return w.check(next, dir) && w.check(next.add(left), dir)
	}

	panic(fmt.Sprint("Unexpected: %s", nextWhat))
}

func (w world) move(what rune, start, dir loc) loc {
	next := start.add(dir)
	nextWhat := w.get(next)

	if nextWhat == '#' {
		return start
	}

	if nextWhat == '.' {
		w.set(start, '.')
		w.set(next, what)
		return next
	}

	if nextWhat == '[' || nextWhat == ']' {
		if dir == left || dir == right {
			if w.move(nextWhat, next, dir) != next {
				return w.move(what, start, dir)
			}
			return start
		}

		if w.check(start, dir) {
			w.move(nextWhat, next, dir)

			if nextWhat == '[' {
				w.move(']', next.add(right), dir)
			} else {
				w.move('[', next.add(left), dir)
			}

			return w.move(what, start, dir)
		}

		return start
	}

	panic(fmt.Sprintf("Unexpected: %c", nextWhat))
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
			if char == '[' {
				total += 100*y + x
			}
		}
	}

	fmt.Println(total)
}
