package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	ADD     = 1
	MULT    = 2
	READ    = 3
	WRITE   = 4
	JPTRUE  = 5
	JPFALSE = 6
	LESS    = 7
	EQU     = 8
	END     = 99
)

type intcode struct {
	mem []int64
	pc  int64
}

func (p *intcode) String() string {
	out := strings.Builder{}

	for i, n := range p.mem {
		out.WriteString(fmt.Sprintf("%d:%d ", i, n))
	}

	return out.String()
}

func (p *intcode) jump(pos int64) {
	p.pc = pos
}

func (p *intcode) fetch(immediate bool) int64 {
	n := p.mem[p.pc]
	p.pc++

	if immediate {
		return n
	}

	return p.get(n)
}

func (p *intcode) get(pos int64) int64 {
	return p.mem[pos]
}

func (p *intcode) set(value, pos int64) {
	p.mem[pos] = value
}

func (p *intcode) run(inputs ...int64) []int64 {
	var modes []bool

	outputs := make([]int64, 0)

	for {
		n := p.fetch(true)
		modes = make([]bool, 4)

		op := n % 100

		n = n / 100
		for i := 0; i < 4; i++ {
			c := n % 10

			if c > 0 {
				modes[i] = true
			}

			n = (n - c) / 10
		}

		switch op {
		case ADD:
			p.set(p.fetch(modes[0])+p.fetch(modes[1]), p.fetch(true))
		case MULT:
			p.set(p.fetch(modes[0])*p.fetch(modes[1]), p.fetch(true))
		case READ:
			var input int64

			if len(inputs) > 0 {
				input, inputs = inputs[0], inputs[1:]
			} else {
				fmt.Print("? ")
				_, err := fmt.Scanln(&input)
				if err != nil {
					panic(err)
				}
			}

			p.set(input, p.fetch(true))
		case WRITE:
			out := p.fetch(modes[0])
			outputs = append(outputs, out)
			fmt.Println(">", out)
		case JPTRUE:
			cmp := p.fetch(modes[0])
			pos := p.fetch(modes[1])

			if cmp != 0 {
				p.jump(pos)
			}
		case JPFALSE:
			cmp := p.fetch(modes[0])
			pos := p.fetch(modes[1])

			if cmp == 0 {
				p.jump(pos)
			}
		case LESS:
			if p.fetch(modes[0]) < p.fetch(modes[1]) {
				p.set(1, p.fetch(true))
			} else {
				p.set(0, p.fetch(true))
			}
		case EQU:
			if p.fetch(modes[0]) == p.fetch(modes[1]) {
				p.set(1, p.fetch(true))
			} else {
				p.set(0, p.fetch(true))
			}
		case END:
			return outputs
		}
	}
}

func (p intcode) clone() intcode {
	out := intcode{
		mem: make([]int64, len(p.mem)),
		pc:  0,
	}

	for i, n := range p.mem {
		out.mem[i] = n
	}

	return out
}

func main() {
	orig := intcode{}

	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(input), ",")

	orig.mem = make([]int64, len(parts))
	for i, part := range parts {
		n, err := strconv.ParseInt(strings.TrimSpace(part), 0, 64)
		if err != nil {
			panic(err)
		}

		orig.set(n, int64(i))
	}

	var (
		max           *int64
		i, j, k, l, m int64
	)

	for i = 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			if j == i {
				continue
			}
			for k = 0; k < 5; k++ {
				if k == i || k == j {
					continue
				}
				for l = 0; l < 5; l++ {
					if l == i || l == j || l == k {
						continue
					}
					for m = 0; m < 5; m++ {
						if m == i || m == j || m == k || m == l {
							continue
						}

						settings := []int64{i, j, k, l, m}

						out := []int64{0}
						for _, s := range settings {
							p := orig.clone()
							out = p.run(int64(s), out[0])
						}

						if max == nil || out[0] > *max {
							o := out[0]
							max = &o
						}
					}
				}
			}
		}
	}

	fmt.Println(*max)
}
