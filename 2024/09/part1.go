package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type drive struct {
	id     int
	blocks []int
	head   int
}

func (d drive) String() string {
	out := strings.Builder{}

	for _, id := range d.blocks {
		if id >= 0 {
			out.WriteString(fmt.Sprint(id))
		} else {
			out.WriteString(".")
		}
	}

	return out.String()
}

func (d *drive) defrag() {
	d.head = 0

	for i := len(d.blocks) - 1; i >= 0; i-- {
		n := d.blocks[i]

		if n == -1 {
			continue
		}

		d.findFree()

		if d.head >= i {
			return
		}

		d.blocks[d.head] = n
		d.blocks[i] = -1
	}
}

func (d *drive) findFree() {
	for ; d.head < len(d.blocks); d.head++ {
		if d.blocks[d.head] == -1 {
			return
		}
	}
}

func (d drive) checksum() int {
	total := 0

	for i, id := range d.blocks {
		if id >= 0 {
			total += i * id
		}
	}

	return total
}

func main() {
	data, _ := os.ReadFile("data")

	d := drive{
		id:     0,
		blocks: make([]int, 0),
		head:   0,
	}

	file := true
	for _, char := range strings.TrimSpace(string(data)) {
		n, _ := strconv.Atoi(string(char))

		if file {
			for i := 0; i < n; i++ {
				d.blocks = append(d.blocks, d.id)
			}
		} else {
			for i := 0; i < n; i++ {
				d.blocks = append(d.blocks, -1)
			}
			d.id++
		}

		file = !file
	}

	d.defrag()

	fmt.Println(d.checksum())
}
