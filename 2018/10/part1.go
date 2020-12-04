package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var count = 0

var smallest = 1000000

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>`)
}

type point struct {
	x, y int
}

type node struct {
	pos point
	vel point
}

func (n *node) move() {
	n.pos.x += n.vel.x
	n.pos.y += n.vel.y
}

func display(nodes []*node) {
	// Find min and max
	var minx, miny, maxx, maxy *int

	count++

	for _, n := range nodes {
		// Move the node :)
		n.move()

		if minx == nil || n.pos.x < *minx {
			x := n.pos.x
			minx = &x
		}
		if miny == nil || n.pos.y < *miny {
			y := n.pos.y
			miny = &y
		}
		if maxx == nil || n.pos.x > *maxx {
			x := n.pos.x
			maxx = &x
		}
		if maxy == nil || n.pos.y > *maxy {
			y := n.pos.y
			maxy = &y
		}
	}

	size := (*maxx - *maxx) + (*maxy - *miny)

	if size < smallest {
		smallest = size
	}

	if size < 100 && size == smallest {
		w := *maxx - *minx + 1
		h := *maxy - *miny + 1

		fmt.Printf("%d x %d: %d\n", w, h, count)

		message := make([][]bool, h)
		for i, _ := range message {
			message[i] = make([]bool, w)
		}

		for _, n := range nodes {
			y := n.pos.y - *miny
			x := n.pos.x - *minx

			message[y][x] = true
		}

		for _, line := range message {
			for _, s := range line {
				if s {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println("---")
	}
}

func main() {
	data, _ := ioutil.ReadFile("input")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	nodes := make([]*node, 0)

	// Parse the nodes
	for _, line := range lines {
		var n node

		match := re.FindStringSubmatch(line)

		n.pos.x, _ = strconv.Atoi(match[1])
		n.pos.y, _ = strconv.Atoi(match[2])
		n.vel.x, _ = strconv.Atoi(match[3])
		n.vel.y, _ = strconv.Atoi(match[4])

		nodes = append(nodes, &n)
	}

	for {
		display(nodes)
	}
}
