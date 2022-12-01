package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type elf []int

func newElf(in string) elf {
	out := make(elf, 0)

	for _, line := range strings.Split(in, "\n") {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		out = append(out, num)
	}

	return out
}

func (e elf) sum() int {
	sum := 0
	for _, cals := range e {
		sum += cals
	}

	return sum
}

type elves []elf

func (e elves) Len() int {
	return len(e)
}

func (e elves) Less(i, j int) bool {
	return e[i].sum() < e[j].sum()
}

func (e elves) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func newElves(in string) elves {
	out := make(elves, 0)

	for _, line := range strings.Split(in, "\n\n") {
		out = append(out, newElf(line))
	}

	return out
}

func main() {
	f, _ := os.ReadFile("input")

	dudes := newElves(string(f))

	// Sort the elves
	sort.Sort(sort.Reverse(dudes))

	// Top 1
	fmt.Println(dudes[0].sum())

	// Top 3
	fmt.Println(dudes[0].sum() + dudes[1].sum() + dudes[2].sum())
}
