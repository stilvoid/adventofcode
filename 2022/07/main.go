package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Tree struct {
	entries []*Tree
	size    int
	name    string
	parent  *Tree
}

func (t *Tree) Size() int {
	size := t.size
	for _, entry := range t.entries {
		size += entry.Size()
	}
	return size
}

func (t *Tree) AddEntry(entry *Tree) {
	t.entries = append(t.entries, entry)
}

func (t *Tree) FindDirs() chan *Tree {
	ch := make(chan *Tree)

	go func() {
		defer close(ch)

		for _, entry := range t.entries {
			if len(entry.entries) > 0 {
				ch <- entry

				for dir := range entry.FindDirs() {
					ch <- dir
				}
			}
		}
	}()

	return ch
}

func newFileTree(name string, size int) *Tree {
	return &Tree{
		name: name,
		size: size,
	}
}

func newDirTree(name string, parent *Tree) *Tree {
	return &Tree{
		name:    name,
		entries: []*Tree{},
		parent:  parent,
	}
}

func parseTree(in []string) *Tree {
	root := newDirTree("/", nil)
	t := root

Outer:
	for i := 0; i < len(in); i++ {
		line := in[i]
		parts := strings.Fields(line)

		if parts[0] != "$" {
			panic("parse error")
		}

		if parts[1] == "ls" {
			for {
				i++
				if i >= len(in) {
					break
				}
				next := in[i]
				entry := strings.Fields(next)
				if entry[0] == "$" {
					i--
					continue Outer
				} else if entry[0] == "dir" {
					continue
				}

				size, err := strconv.Atoi(entry[0])
				if err != nil {
					panic(err)
				}
				t.AddEntry(newFileTree(entry[1], size))
			}
		} else if parts[1] == "cd" {
			if parts[2] == "/" {
				t = root
			} else if parts[2] == ".." {
				t = t.parent
			} else {
				n := newDirTree(parts[2], t)
				t.AddEntry(n)
				t = n
			}
		} else {
			panic("parse error")
		}
	}

	return root
}

type Trees []*Tree

func (t Trees) Len() int {
	return len(t)
}

func (t Trees) Less(i, j int) bool {
	return t[i].Size() < t[j].Size()
}

func (t Trees) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func main() {
	input, _ := os.ReadFile("input")
	cmds := strings.Split(string(input), "\n")

	root := parseTree(cmds)

	totalSpace := 70000000
	needed := 30000000
	have := totalSpace - root.Size()
	want := needed - have
	fmt.Println(want)

	part2 := make(Trees, 0)

	total := 0
	for dir := range root.FindDirs() {
		if dir.Size() <= 100000 {
			total += dir.Size()
		}

		if dir.Size() >= want {
			part2 = append(part2, dir)
		}
	}

	fmt.Println("Part 1:", total)

	sort.Sort(part2)
	fmt.Println("Part 2:", part2[0].Size())
}
