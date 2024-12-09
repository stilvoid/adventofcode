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
	sizes  map[int]int
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
	for i := len(d.blocks) - 1; i >= 0; i-- {
		n := d.blocks[i]

		if n == -1 {
			continue
		}

		size := d.sizes[n]

		start := d.findFree(size)

		if start != -1 && start < i-size {
			for j := 0; j < size; j++ {
				d.blocks[start+j] = d.blocks[i-j]
				d.blocks[i-j] = -1
			}
		}

		//i -= size - 1
	}
}

func (d *drive) findFree(size int) int {
	found := 0
	start := 0

	for i := 0; i < len(d.blocks)-size-1; i++ {
		if d.blocks[i] == -1 {
			if found == 0 {
				start = i
			}
			found++

			if found == size {
				return start
			}
		} else {
			found = 0
		}
	}

	return -1
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
		sizes:  make(map[int]int),
	}

	file := true
	for _, char := range strings.TrimSpace(string(data)) {
		n, _ := strconv.Atoi(string(char))

		if file {
			for i := 0; i < n; i++ {
				d.blocks = append(d.blocks, d.id)
			}
			d.sizes[d.id] = n
		} else {
			for i := 0; i < n; i++ {
				d.blocks = append(d.blocks, -1)
			}
			d.id++
		}

		file = !file
	}

	fmt.Println(d)
	d.defrag()
	fmt.Println(d)

	fmt.Println(d.checksum())
}
