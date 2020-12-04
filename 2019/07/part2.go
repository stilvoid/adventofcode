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
	name  string
	debug bool
	mem   []int64
	pc    int64
}

func (p *intcode) Debug(pattern string, parts ...interface{}) {
	if p.debug {
		fmt.Printf(pattern, parts...)
	}
}

func (p *intcode) String() string {
	out := strings.Builder{}

	out.WriteString(fmt.Sprintf("%s) @%d ", p.name, p.pc))

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

func (p *intcode) run(input <-chan int64, output chan<- int64) <-chan int64 {
	var modes []bool

	var last int64

	finish := make(chan int64)

	go func() {
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
				p.Debug("%s READ\n", p.name)
				p.set(<-input, p.fetch(true))
			case WRITE:
				last = p.fetch(modes[0])
				p.Debug("%s WRITE %d\n", p.name, last)
				output <- last
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
				p.Debug("Exit %s: %d\n", p.name, last)
				finish <- last
				close(output)
				close(finish)

				return
			}
		}
	}()

	return finish
}

func (p intcode) clone(name string) intcode {
	out := intcode{
		mem: make([]int64, len(p.mem)),
		pc:  0,
	}

	for i, n := range p.mem {
		out.mem[i] = n
	}

	out.name = name

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

	for i = 5; i < 10; i++ {
		for j = 5; j < 10; j++ {
			if j == i {
				continue
			}

			for k = 5; k < 10; k++ {
				if k == i || k == j {
					continue
				}

				for l = 5; l < 10; l++ {
					if l == i || l == j || l == k {
						continue
					}

					for m = 5; m < 10; m++ {
						if m == i || m == j || m == k || m == l {
							continue
						}

						a := orig.clone("a")
						ac := make(chan int64, 100)

						b := orig.clone("b")
						bc := make(chan int64, 100)

						c := orig.clone("c")
						cc := make(chan int64, 100)

						d := orig.clone("d")
						dc := make(chan int64, 100)

						e := orig.clone("e")
						ec := make(chan int64, 100)

						a.run(ac, bc)
						b.run(bc, cc)
						c.run(cc, dc)
						d.run(dc, ec)
						done := e.run(ec, ac)

						ac <- i
						bc <- j
						cc <- k
						dc <- l
						ec <- m

						// Trigger a
						ac <- 0

						out := <-done

						if max == nil || out > *max {
							max = &out
						}
					}
				}
			}
		}
	}

	fmt.Println(*max)
}
