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

func (p *intcode) get(pos int64) int64 {
	return p.mem[pos]
}

func (p *intcode) set(value, pos int64) {
	p.mem[pos] = value
}

func (p *intcode) getref() int64 {
	return p.mem[p.fetch()]
}

func (p *intcode) setref(value int64) {
	p.mem[p.fetch()] = value
}

func (p *intcode) run() {
	for {
		op := p.fetch()

		switch op {
		case ADD:
			p.setref(p.getref() + p.getref())
		case MULT:
			p.setref(p.getref() * p.getref())
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

	for noun := int64(0); noun < 99; noun++ {
		for verb := int64(0); verb < 99; verb++ {
			p := orig.clone()

			p.set(noun, 1)
			p.set(verb, 2)

			p.run()

			if p.get(0) == 19690720 {
				fmt.Println(100*noun + verb)
				return
			}
		}
	}
}
