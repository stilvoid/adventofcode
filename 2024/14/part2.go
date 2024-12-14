package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	w = 101
	h = 103
)

type vec struct {
	x, y int
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func (v vec) add(o vec) vec {
	x := mod(v.x+o.x, w)
	y := mod(v.y+o.y, h)

	return vec{x, y}
}

type robot struct {
	p vec
	v vec
}

func newRobot(in string) *robot {
	var px, py, vx, vy int
	fmt.Sscanf(in, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)

	return &robot{
		p: vec{px, py},
		v: vec{vx, vy},
	}
}

func draw(robots []*robot) string {
	out := strings.Builder{}
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			n := 0
			for _, robot := range robots {
				if robot.p.x == x && robot.p.y == y {
					n++
				}
			}

			if n > 0 {
				out.WriteString("#")
			} else {
				out.WriteString(" ")
			}
		}
		out.WriteString("\n")
	}
	return out.String()
}

func main() {
	data, _ := os.ReadFile("data")

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	robots := make([]*robot, len(lines))

	for i, line := range lines {
		robots[i] = newRobot(line)
	}

	draw(robots)

	for i := 0; i < 100000; i++ {
		if i%1000 == 0 {
			fmt.Println(i)
		}

		for _, robot := range robots {
			robot.p = robot.p.add(robot.v)
		}

		card := draw(robots)
		if strings.Contains(card, strings.Repeat("#", 20)) {
			fmt.Println(card)
			fmt.Println(i)
			fmt.Println()
		}
	}
}
