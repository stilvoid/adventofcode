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

func (p *intcode) run(input <-chan int64, output chan<- int64) <-chan int64 {
	done := make(chan int64, 100)
	var last int64

	go func() {
		var inst int64

		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Died at %d: %d\n", p.pc, inst)
				panic(r)
			}
		}()

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
				done <- last
				close(output)
				close(done)
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

	in := make(chan int64, 100)
	out := make(chan int64, 100)

	done := orig.run(in, out)

	world := struct {
		dir  int64
		pos  struct{ x, y int }
		grid map[int]map[int][]int64
	}{
		grid: map[int]map[int][]int64{
			0: map[int][]int64{
				0: make([]int64, 0),
			},
		},
	}

	// Currently on black
	in <- 0

	for paint := range out {
		turn := <-out

		fmt.Printf("Paint: %d\n", paint)
		fmt.Printf("Turn: %d\n", turn)

		world.grid[world.pos.y][world.pos.x] = append(world.grid[world.pos.y][world.pos.x], paint)

		if turn == 0 {
			world.dir += 3
		} else if turn == 1 {
			world.dir += 1
		} else {
			panic(fmt.Errorf("Unexpected turn: %d", turn))
		}

		world.dir %= 4

		switch world.dir {
		case 0:
			world.pos.y--
		case 1:
			world.pos.x++
		case 2:
			world.pos.y++
		case 3:
			world.pos.x--
		default:
			panic(fmt.Errorf("Unexpected dir: %d", world.dir))
		}

		_, ok := world.grid[world.pos.y]
		if !ok {
			world.grid[world.pos.y] = make(map[int][]int64)
		}

		_, ok = world.grid[world.pos.y][world.pos.x]
		if !ok {
			world.grid[world.pos.y][world.pos.x] = make([]int64, 0)
		}

		point := world.grid[world.pos.y][world.pos.x]

		if len(point) == 0 {
			in <- 0
		} else {
			in <- point[len(point)-1]
		}
	}

	fmt.Println("Exit:", <-done)

	var minX, maxX, minY, maxY *int

	for y, row := range world.grid {
		if minY == nil || y < *minY {
			newY := y
			minY = &newY
		}

		if maxY == nil || y > *maxY {
			newY := y
			maxY = &newY
		}

		for x, _ := range row {
			if minX == nil || x < *minX {
				newX := x
				minX = &newX
			}

			if maxX == nil || x > *maxX {
				newX := x
				maxX = &newX
			}
		}
	}

	fmt.Println(*minX, *maxX, *minY, *maxY)
	count := 0

	for y := *minY; y <= *maxY; y++ {
		row, ok := world.grid[y]
		if !ok {
			row = make(map[int][]int64)
		}

		for x := *minX; x <= *maxX; x++ {
			if point, ok := row[x]; ok {
				if len(point) > 0 {
					count++

					fmt.Print(len(point) % 10)
					continue
				}
			}
			fmt.Print(".")
		}
		fmt.Println()
	}

	fmt.Println(count)
}
