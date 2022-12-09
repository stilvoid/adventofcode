package main

import (
	"fmt"
	"os"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

type point struct {
	x, y int
}

func (p *point) follow(o point) {
	dx := o.x - p.x
	dy := o.y - p.y

	if dx == 0 && abs(dy) > 1 {
		p.y += sgn(dy)
	} else if dy == 0 && abs(dx) > 1 {
		p.x += sgn(dx)
	} else if abs(dx)+abs(dy) > 2 {
		p.x += sgn(dx)
		p.y += sgn(dy)
	}
}

type rope struct {
	parts   []point
	visited map[point]bool
}

func newRope(n int) *rope {
	return &rope{
		parts: make([]point, n),
		visited: map[point]bool{
			{0, 0}: true,
		},
	}
}

func (r *rope) move(dir string, dist int) {
	for i := 0; i < dist; i++ {
		switch dir {
		case "U":
			r.parts[0].y++
		case "D":
			r.parts[0].y--
		case "R":
			r.parts[0].x++
		case "L":
			r.parts[0].x--
		}

		for i := 1; i < len(r.parts); i++ {
			r.parts[i].follow(r.parts[i-1])
		}

		r.visited[r.parts[len(r.parts)-1]] = true
	}
}

func main() {
	input, _ := os.ReadFile("input")

	r1 := newRope(2)
	r2 := newRope(10)

	var dir string
	var dist int

	for _, line := range strings.Split(string(input), "\n") {
		fmt.Sscanf(line, "%s %d", &dir, &dist)
		r1.move(dir, dist)
		r2.move(dir, dist)
	}

	fmt.Println("Part1:", len(r1.visited))
	fmt.Println("Part2:", len(r2.visited))
}
