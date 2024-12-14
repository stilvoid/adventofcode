package main

import (
	"fmt"
	"os"
	"strings"
)

type pos struct {
	x, y int
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

	fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &m.a.x, &m.a.y)
	fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &m.b.x, &m.b.y)
	fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &m.prize.x, &m.prize.y)

	return m
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (m machine) try() int {
	min_toks := -1

	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			x := a*m.a.x + b*m.b.x
			y := a*m.a.y + b*m.b.y

			if x == m.prize.x && y == m.prize.y {
				toks := 3*a + b
				if min_toks == -1 || toks < min_toks {
					min_toks = toks
				}
			}

			if x >= m.prize.x || y >= m.prize.y {
				break
			}
		}
	}

	return min_toks
}

func main() {
	data, _ := os.ReadFile("data")

	parts := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	machines := make([]machine, len(parts))

	for i, part := range parts {
		machines[i] = newMachine(part)
	}

	fmt.Println(machines)

	total := 0
	for _, machine := range machines {
		toks := machine.try()
		if toks != -1 {
			total += toks
		}
	}
	fmt.Println(total)
}
