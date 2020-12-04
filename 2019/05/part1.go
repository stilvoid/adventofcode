package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	ADD   = 1
	MULT  = 2
	READ  = 3
	WRITE = 4
	END   = 99
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

func (p *intcode) run() {
	var modes []bool

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
			fmt.Print("? ")
			var input int64
			_, err := fmt.Scanln(&input)
			if err != nil {
				panic(err)
			}

			p.set(input, p.fetch(true))
		case WRITE:
			fmt.Println(">", p.fetch(modes[0]))
		case END:
			return
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

	p := orig.clone()
	p.run()
}
