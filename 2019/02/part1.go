package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	ADD  = 1
	MULT = 2
	END  = 99
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

func (p *intcode) fetch() int64 {
	op := p.mem[p.pc]
	p.pc++

	return op
}

func (p *intcode) run() {
	for {
		op := p.fetch()

		switch op {
		case ADD:
			l := p.mem[p.fetch()]
			r := p.mem[p.fetch()]
			dst := p.fetch()

			p.mem[dst] = l + r
		case MULT:
			l := p.mem[p.fetch()]
			r := p.mem[p.fetch()]
			dst := p.fetch()

			p.mem[dst] = l * r
		case END:
			return
		}

		fmt.Println(p.String())
	}
}

func main() {
	p := intcode{}

	input, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(string(input), ",")

	p.mem = make([]int64, len(parts))

	for i, part := range parts {
		n, err := strconv.ParseInt(strings.TrimSpace(part), 0, 64)
		if err != nil {
			panic(err)
		}

		p.mem[i] = n
	}

	fmt.Println(p.String())

	p.run()

	fmt.Println(p.String())
}
