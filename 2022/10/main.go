package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	x     int
	t     int
	total int
}

type Op interface {
	Run(c *CPU)
}

type Noop struct{}

func (n Noop) Run(c *CPU) {
}

type StartAddX struct {
	v int
}

func (s StartAddX) Run(c *CPU) {
}

type EndAddX struct {
	v int
}

func (e EndAddX) Run(c *CPU) {
	c.x += e.v
}

type Program []Op

func newProgram(input []string) Program {
	out := make(Program, 0)

	for _, line := range input {
		parts := strings.Fields(line)

		switch parts[0] {
		case "noop":
			out = append(out, Noop{})
		case "addx":
			v, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			out = append(out, StartAddX{v}, EndAddX{v})
		}
	}

	return out
}

type Trace []int

var blankLine = "                                        "

func (t Trace) Render() {
	line := []byte(blankLine)

	for i, n := range t {
		x := i % 40
		if n-1 <= x && x <= n+1 {
			line[x] = '#'
		}

		if i > 0 && i%40 == 0 {
			fmt.Println(string(line))
			line = []byte(blankLine)
		}
	}
}

func (p Program) Run(c *CPU) Trace {
	out := make(Trace, 0)
	out = append(out, c.x)

	for _, op := range p {
		op.Run(c)
		out = append(out, c.x)
	}

	return out
}

func newCPU() *CPU {
	return &CPU{
		x: 1,
	}
}

var interesting = []int{20, 60, 100, 140, 180, 220}

func main() {
	input, _ := os.ReadFile("input")

	c := newCPU()
	p := newProgram(strings.Split(string(input), "\n"))

	out := p.Run(c)

	total := 0
	for _, n := range interesting {
		total += n * out[n-1]
	}
	fmt.Println("Part 1:", total)
	fmt.Println()

	fmt.Println("Part 2")
	out.Render()
}
