package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type node struct {
	name string
	done bool
	deps map[string]bool
}

var nodes map[string]*node

func init() {
	nodes = make(map[string]*node)
}

func newNode(name string) *node {
	n := node{
		name: name,
		done: false,
		deps: make(map[string]bool),
	}

	nodes[name] = &n

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

func (n *node) complete() {
	n.done = true

	for _, node := range nodes {
		node.deps[n.name] = true
	}
}

func main() {
	data, _ := ioutil.ReadFile("input")
	content := strings.TrimSpace(string(data))

	names := make([]string, 0)

	// Make the nodes
	for _, line := range strings.Split(content, "\n") {
		var dep, name string

		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &dep, &name)

		if _, ok := nodes[name]; !ok {
			newNode(name)
			names = append(names, name)
		}

		if _, ok := nodes[dep]; !ok {
			newNode(dep)
			names = append(names, dep)
		}

		nodes[name].deps[dep] = false
	}

	// Get a sorted list of names
	sort.Strings(names)

	// Start working
	for {
		var start *node = nil

		// Find the first node
		for _, name := range names {
			if n := nodes[name]; n.isReady() {
				start = n
				break
			}
		}

		// Couldn't find a node, we've finished
		if start == nil {
			break
		}

		fmt.Print(start.name)

		start.complete()
	}

	fmt.Println()
}
