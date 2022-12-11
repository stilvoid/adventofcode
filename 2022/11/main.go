package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/alecthomas/participle/v2"
)

type Program struct {
	Monkeys []*Monkey `@@*`

	div int64
}

type Monkey struct {
	ID          int        `"Monkey" @Int ":"`
	Items       []int64    `"Starting" "items" ":" @Int ("," @Int)*`
	Operation   *Operation `"Operation" ":" "new" "=" "old" @@`
	Test        int64      `"Test" ":" "divisible" "by" @Int`
	TrueTarget  int        `"If" "true" ":" "throw" "to" "monkey" @Int`
	FalseTarget int        `"If" "false" ":" "throw" "to" "monkey" @Int`

	inspections int
}

type Operation struct {
	Operator string `@("+" | "*")`
	Number   int64  `(@Int`
	Old      bool   `| @"old")`
}

func (p *Program) Run(n int) {
	// Set up the divisor
	p.div = p.Monkeys[0].Test
	for i := 1; i < len(p.Monkeys); i++ {
		p.div *= p.Monkeys[i].Test
	}

	for i := 0; i < n; i++ {
		for _, m := range p.Monkeys {
			m.Run(p)
		}
	}
}

func (m *Monkey) Run(p *Program) {
	for _, worry := range m.Items {
		worry = m.Operation.Run(worry)

		//worry /= 3

		var other *Monkey

		worry %= p.div

		if worry%m.Test == 0 {
			other = p.Monkeys[m.TrueTarget]
		} else {
			other = p.Monkeys[m.FalseTarget]
		}

		other.Items = append(other.Items, worry)

		m.inspections++
	}

	m.Items = []int64{}
}

func (o *Operation) Run(worry int64) int64 {
	var v int64
	if o.Old {
		v = worry
	} else {
		v = o.Number
	}

	if o.Operator == "+" {
		return worry + v
	}

	return worry * v
}

func (m *Program) Len() int {
	return len(m.Monkeys)
}

func (m *Program) Less(i, j int) bool {
	return m.Monkeys[i].inspections < m.Monkeys[j].inspections
}

func (m *Program) Swap(i, j int) {
	m.Monkeys[i], m.Monkeys[j] = m.Monkeys[j], m.Monkeys[i]
}

func main() {
	nyp := participle.MustBuild[Program]()

	f, _ := os.Open("input")

	// Parse the program
	prog, err := nyp.Parse("input", f)
	if err != nil {
		panic(err)
	}

	prog.Run(10000)

	sort.Sort(sort.Reverse(prog))

	for _, m := range prog.Monkeys {
		fmt.Println(m.inspections)
	}

	fmt.Println(prog.Monkeys[0].inspections * prog.Monkeys[1].inspections)
}
