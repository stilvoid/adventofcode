package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type node struct {
	name        string
	in_progress bool
	done        bool
	deps        map[string]bool
	time        int
}

var names []string
var nodes map[string]*node

func init() {
	names = make([]string, 0)
	nodes = make(map[string]*node)
}

func newNode(name string) *node {
	n := node{
		name:        name,
		in_progress: false,
		done:        false,
		deps:        make(map[string]bool),
		time:        int(name[0]-'A') + 1 + 60,
	}

	nodes[name] = &n
	names = append(names, name)
	sort.Strings(names)

	return &n
}

func (n *node) isReady() bool {
	if n.done {
		return false
	}

	for _, done := range n.deps {
		if !done {
			return false
		}
	}

	return true
}

func (n *node) work() {
	n.time--

	if n.time == 0 {
		n.done = true

		for _, node := range nodes {
			node.deps[n.name] = true
		}
	}
}

func main() {
	data, _ := ioutil.ReadFile("input")
	content := strings.TrimSpace(string(data))

	// Make the nodes
	for _, line := range strings.Split(content, "\n") {
		var dep, name string

		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &dep, &name)

		if _, ok := nodes[name]; !ok {
			newNode(name)
		}

		if _, ok := nodes[dep]; !ok {
			newNode(dep)
		}

		nodes[name].deps[dep] = false
	}

	for _, name := range names {
		fmt.Printf("%s: %#v\n", name, nodes[name])
	}

	sec := 0
	workers := make([]string, 5)

	// Start working
	for {
		fmt.Printf("%5d ", sec)

		for i, name := range workers {
			var start *node
			var err error

			if name != "" && !nodes[name].done {
				start = nodes[name]
			} else {
				start, err = nextNode()
				if err != nil {
					// We've finished!
					fmt.Println()
					return
				}
			}

			// Nothing to work on
			if start == nil {
				workers[i] = ""
				fmt.Print(". ")
				continue
			}

			// Do the node
			workers[i] = start.name
			fmt.Printf("%s ", start.name)
			start.in_progress = true
		}

		for _, name := range workers {
			if name != "" {
				nodes[name].work()
			}
		}

		sec++
		fmt.Println()
	}
}

func nextNode() (*node, error) {
	doneCount := 0

	// Find the first available node
	for _, name := range names {
		n := nodes[name]

		if !n.in_progress && n.isReady() {
			return n, nil
		}

		if n.done {
			doneCount++
		}
	}

	if doneCount == len(nodes) {
		return nil, fmt.Errorf("All nodes done")
	}

	return nil, nil
}
