package main

import (
	"fmt"
	"os"
	"strings"
)

type machine struct {
	stacks []stack
}

func newMachine(state string) machine {
	m := machine{
		stacks: []stack{},
	}

	lines := strings.Split(state, "\n")
	for n := len(lines) - 2; n >= 0; n-- {
		line := lines[n]

		fmt.Println(line)
		for i := 0; i < 1+len(line)/4; i++ {
			c := line[4*i+1]

			if c != ' ' {
				m.push(i, c)
			}
		}
	}

	return m
}

func (m *machine) push(pos int, value byte) {
	for i := len(m.stacks); i < pos+1; i++ {
		m.stacks = append(m.stacks, stack{})
	}

	m.stacks[pos].push(value)
}

func (m *machine) peek() string {
	out := strings.Builder{}

	for _, stack := range m.stacks {
		out.WriteByte(stack.peek())
	}

	return out.String()
}

func (m *machine) exec(input string) {
	for _, line := range strings.Split(input, "\n") {
		var n, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &n, &from, &to)

		s := stack{}

		for i := 0; i < n; i++ {
			s.push(m.stacks[from-1].pop())
		}

		for i := 0; i < n; i++ {
			m.stacks[to-1].push(s.pop())
		}
	}
}

type stack []byte

func (s *stack) push(value byte) {
	*s = append(*s, value)
}

func (s *stack) pop() byte {
	out := s.peek()
	*s = (*s)[:len(*s)-1]
	return out
}

func (s *stack) peek() byte {
	return (*s)[len(*s)-1]
}

func main() {
	f, _ := os.ReadFile("input")

	parts := strings.Split(string(f), "\n\n")

	m := newMachine(parts[0])

	fmt.Println(m.peek())

	m.exec(parts[1])

	fmt.Println(m.peek())
}
