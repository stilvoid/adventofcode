package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type pos struct {
	x, y float64
}

func (p pos) add(o pos) pos {
	return pos{
		p.x + o.x,
		p.y + o.y,
	}
}

type machine struct {
	a, b, prize pos
}

func newMachine(input string) machine {
	lines := strings.Split(input, "\n")

	m := machine{}

	fmt.Sscanf(lines[0], "Button A: X+%f, Y+%f", &m.a.x, &m.a.y)
	fmt.Sscanf(lines[1], "Button B: X+%f, Y+%f", &m.b.x, &m.b.y)
	fmt.Sscanf(lines[2], "Prize: X=%f, Y=%f", &m.prize.x, &m.prize.y)

	m.prize.x += 10000000000000
	m.prize.y += 10000000000000

	return m
}

func solve(px, py, ax, ay, bx, by float64) int64 {
	var a, b float64

	a = (px*by - py*bx) / (ax*by - ay*bx)
	b = (py - a*ay) / by

	if a == math.Trunc(a) && b == math.Trunc(b) {
		return 3*int64(a) + int64(b)
	}

	return 0
}

func (m machine) try() int64 {
	return solve(m.prize.x, m.prize.y, m.a.x, m.a.y, m.b.x, m.b.y)
}

func main() {
	data, _ := os.ReadFile("data")

	parts := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	machines := make([]machine, len(parts))

	for i, part := range parts {
		machines[i] = newMachine(part)
	}

	var total int64 = 0
	for _, machine := range machines {
		toks := machine.try()
		if toks != -1 {
			fmt.Println(toks)
			total += toks
		}
	}
	fmt.Println(total)
}
