package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type chemical struct {
	name   string
	stock  int
	used   int
	output int
	recipe map[string]int
}

func (c *chemical) Make() {
	for ingredient, amount := range c.recipe {
		for cupboard[ingredient].stock < amount {
			cupboard[ingredient].Make()
		}

		cupboard[ingredient].stock -= amount
		cupboard[ingredient].used += amount
	}

	cupboard[c.name].stock += c.output
}

func (c chemical) String() string {
	out := strings.Builder{}

	out.WriteString(fmt.Sprintf("%s: %d / %d {\n", c.name, c.stock, c.used))
	for ingredient, amount := range c.recipe {
		out.WriteString(fmt.Sprintf("    %s: %d\n", ingredient, amount))
	}
	out.WriteString(fmt.Sprintf("} -> %d\n", c.output))

	return out.String()
}

type stock map[string]*chemical

func (s stock) String() string {
	out := strings.Builder{}

	for _, c := range s {
		out.WriteString(c.String())
		out.WriteString("\n")
	}

	return out.String()
}

var cupboard = make(stock)

func init() {
	cupboard["ORE"] = &chemical{
		name:   "ORE",
		output: 1,
		recipe: make(map[string]int),
	}
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		parts := strings.Split(line, "=>")
		head, tail := parts[0], parts[1]

		c := chemical{
			recipe: make(map[string]int),
		}

		fmt.Sscanf(tail, "%d %s", &c.output, &c.name)
		cupboard[c.name] = &c

		for _, part := range strings.Split(head, ", ") {
			var ingredient string
			var count int
			fmt.Sscanf(part, "%d %s", &count, &ingredient)
			c.recipe[ingredient] = count
		}
	}

	cupboard["FUEL"].Make()
	fmt.Println(cupboard["ORE"].String())
}
