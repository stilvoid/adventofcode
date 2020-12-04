package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	// Instructions
	ADD     = 1
	MULT    = 2
	READ    = 3
	WRITE   = 4
	JPTRUE  = 5
	JPFALSE = 6
	LESS    = 7
	EQU     = 8
	BASE    = 9
	END     = 99
)

const (
	// Modes
	POS = 0
	IMM = 1
	REL = 2
)

type intcode struct {
	mem   []int64
	pc    int
	base  int
	debug bool
}

func (p intcode) Debug(pattern string, parts ...interface{}) {
	if p.debug {
		fmt.Printf(pattern, parts...)
		fmt.Println()
	}
}

func (p *intcode) String() string {
	out := strings.Builder{}

	for i, n := range p.mem {
		out.WriteString(fmt.Sprintf("%d:%d ", i, n))
	}

	return out.String()
}

func (p *intcode) jump(pos int) {
	p.pc = pos
}

func (p *intcode) get(pos int) int64 {
	if pos < 0 {
		panic(fmt.Errorf("Attempt to read invalid address: %d", pos))
	}

	if pos >= len(p.mem) {
		return 0
	}

	return p.mem[pos]
}

func (p *intcode) set(value int64, pos int) {
	p.grow(pos)

	p.mem[pos] = value
}

func (p *intcode) next() int64 {
	n := p.mem[p.pc]
	p.pc++

	return n
}

func (p *intcode) fetch(mode int) int64 {
	n := p.next()

	switch mode {
	case POS:
		return p.get(int(n))
	case IMM:
		return n
	case REL:
		return p.get(p.base + int(n))
	}

	panic(fmt.Errorf("Unsupported mode: %d", mode))
}

func (p *intcode) value(mode int) int64 {
	n := p.next()

	switch mode {
	case POS:
		return n
	case REL:
		return int64(p.base) + n
	}

	panic(fmt.Errorf("Unsupported mode: %d", mode))
}

func (p *intcode) grow(size int) {
	if size >= len(p.mem) {
		new := make([]int64, size-len(p.mem)+1)

		p.mem = append(p.mem, new...)
	}
}

func (p *intcode) run(input chan int64, output chan<- int64) <-chan bool {
	var last int64
	var inst int64

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Died at %d: %d\n", p.pc, inst)
			panic(r)
		}
	}()

	done := make(chan bool)

	go func() {
		for {
			inst = p.fetch(IMM)
			n := inst
			modes := make([]int, 3)

			op := n % 100

			n = n / 100
			for i := 0; i < 3; i++ {
				mode := n % 10

				if mode > 2 {
					panic(fmt.Errorf("Invalid mode '%d' in %d at %d", mode, inst, p.pc-1))
				}

				modes[i] = int(mode)

				n = (n - mode) / 10
			}

			switch op {
			case ADD:
				a, b := p.fetch(modes[0]), p.fetch(modes[1])
				pos := int(p.value(modes[2]))

				p.Debug("(%d) %d+%d->%d", inst, a, b, pos)

				p.set(a+b, pos)
			case MULT:
				a, b := p.fetch(modes[0]), p.fetch(modes[1])
				pos := int(p.value(modes[2]))

				p.Debug("(%d) %d*%d->%d", inst, a, b, pos)

				p.set(a*b, pos)
			case READ:
				// Notify
				input <- 0

				in := <-input
				pos := int(p.value(modes[0]))

				p.Debug("(%d) Read %d->%d", inst, in, pos)

				p.set(in, pos)
			case WRITE:
				last = p.fetch(modes[0])

				p.Debug("(%d) Write: %d", inst, last)

				output <- last
			case JPTRUE:
				cmp := p.fetch(modes[0])
				pos := int(p.fetch(modes[1]))

				p.Debug("(%d) true %d->%d", inst, cmp, pos)

				if cmp != 0 {
					p.jump(pos)
				}
			case JPFALSE:
				cmp := p.fetch(modes[0])
				pos := int(p.fetch(modes[1]))

				p.Debug("(%d) false %d->%d", inst, cmp, pos)

				if cmp == 0 {
					p.jump(pos)
				}
			case LESS:
				a, b := p.fetch(modes[0]), p.fetch(modes[1])
				pos := int(p.value(modes[2]))

				p.Debug("(%d) %d<%d->%d", inst, a, b, pos)

				if a < b {
					p.set(1, pos)
				} else {
					p.set(0, pos)
				}
			case EQU:
				a, b := p.fetch(modes[0]), p.fetch(modes[1])
				pos := int(p.value(modes[2]))

				p.Debug("(%d) %d==%d->%d", inst, a, b, pos)

				if a == b {
					p.set(1, pos)
				} else {
					p.set(0, pos)
				}
			case BASE:
				n := p.fetch(modes[0])
				p.Debug("(%d) Base %d+%d", inst, p.base, n)
				p.base += int(n)
			case END:
				close(output)
				done <- true
				return
			default:
				panic(fmt.Errorf("Invalid opcode: %d", op))
			}
		}
	}()

	return done
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

type mappy map[int64]map[int64]rune

func newMappy() mappy {
	return make(mappy)
}

func (m mappy) Plot(x, y int64, value rune) {
	if _, ok := m[y]; !ok {
		m[y] = make(map[int64]rune)
	}

	m[y][x] = value
}

func (m mappy) Get(x, y int64) rune {
	if row, ok := m[y]; ok {
		return row[x]
	} else {
		return 0
	}
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

		orig.set(n, int(i))
	}

	in := make(chan int64)
	out := make(chan int64)

	done := orig.run(in, out)

	var robX, robY int64
	var baseX, baseY int64
	var baseChanged bool
	var dir int64
	world := newMappy()
	visits := newMappy()

	// Hide the cursor
	fmt.Print("\033[?25l")

	count := 0

	for {
		baseChanged = false

		select {
		case <-done:
			return
		case <-in:
			var min *rune

			for try := int64(0); try < 4; try++ {
				newX, newY := robX, robY
				switch try {
				case 0:
					newY--
				case 1:
					newY++
				case 2:
					newX--
				case 3:
					newX++
				}

				if world.Get(newX, newY) != '#' {
					if min == nil || visits.Get(newX, newY) < *min {
						v := visits.Get(newX, newY)
						min = &v
						dir = try
					}
				}
			}

			in <- dir + 1
		case c := <-out:
			newX, newY := robX, robY
			switch dir {
			case 0:
				newY--
			case 1:
				newY++
			case 2:
				newX--
			case 3:
				newX++
			}

			switch c {
			case 0:
				world.Plot(newX, newY, '#')
			case 1:
				visits.Plot(newX, newY, visits.Get(newX, newY)+1)
				world.Plot(newX, newY, '.')
				robX = newX
				robY = newY
			case 2:
				visits.Plot(newX, newY, visits.Get(newX, newY)+1)
				world.Plot(newX, newY, 'O')
				robX = newX
				robY = newY
			default:
				panic(fmt.Errorf("Invalid robot output: %d", c))
			}

			if robX < baseX {
				baseChanged = true
				baseX = robX
			}

			if robY < baseY {
				baseChanged = true
				baseY = robY
			}
		}

		// Clear the screen
		if baseChanged {
			fmt.Print("\033[1;1H\033[2J")
		}

		// Draw the map
		for y, row := range world {
			for x, c := range row {
				fmt.Printf("\033[%d;%dH", 2+y-baseY, 2+x-baseX)
				fmt.Print(string(c))
			}
		}

		// Draw the start
		fmt.Printf("\033[%d;%dH", 2-baseY, 2-baseX)
		fmt.Print("S")

		// Draw the robot
		fmt.Printf("\033[%d;%dH", 2+robY-baseY, 2+robX-baseX)
		fmt.Print("D")

		//time.Sleep(1 * time.Millisecond)

		count++
		if count == 10000 {
			break
		}
	}

	fmt.Println("Now time to find the map")
}
