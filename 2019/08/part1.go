package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type image struct {
	width, height int
	layers        []layer
}

type layer struct {
	rows []row
}

type row []int

func newImage(w, h int, data []int) image {
	size := len(data) / (w * h)

	out := image{
		width:  w,
		height: h,
		layers: make([]layer, size),
	}

	for l := 0; l < size; l++ {
		out.layers[l].rows = make([]row, h)

		for y := 0; y < h; y++ {
			out.layers[l].rows[y] = make(row, w)

			for x := 0; x < w; x++ {
				out.layers[l].rows[y][x] = data[l*w*h+y*w+x]
			}
		}
	}

	return out
}

func main() {
	raw, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(raw))

	data := make([]int, len(input))
	for i, c := range input {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}

		data[i] = n
	}

	im := newImage(25, 6, data)

	var fewest *int
	var selected layer

	for _, l := range im.layers {
		count := 0
		for _, row := range l.rows {
			for _, pix := range row {
				if pix == 0 {
					count++
				}
			}
		}

		if fewest == nil || count < *fewest {
			fewest = &count
			selected = l
		}
	}

	ones := 0
	twos := 0
	for _, row := range selected.rows {
		for _, pix := range row {
			if pix == 1 {
				ones++
			} else if pix == 2 {
				twos++
			}
		}
	}

	fmt.Println(ones * twos)
}
