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

type layer []row

type row []int

func newImage(w, h int, data []int) image {
	size := len(data) / (w * h)

	out := image{
		width:  w,
		height: h,
		layers: make([]layer, size),
	}

	for l := 0; l < size; l++ {
		out.layers[l] = make(layer, h)

		for y := 0; y < h; y++ {
			out.layers[l][y] = make(row, w)

			for x := 0; x < w; x++ {
				out.layers[l][y][x] = data[l*w*h+y*w+x]
			}
		}
	}

	return out
}

// Flatten to a single layer
func (im image) resolve() layer {
	out := make(layer, im.height)

	for y := 0; y < im.height; y++ {
		out[y] = make(row, im.width)
	}

	for i, l := range im.layers {
		for y, r := range l {
			for x, pix := range r {
				if i == 0 || out[y][x] == 2 {
					out[y][x] = pix
				}
			}
		}
	}

	return out
}

func (l layer) String() string {
	out := strings.Builder{}

	for _, row := range l {
		for _, pix := range row {
			if pix == 0 {
				out.WriteString(".")
			} else if pix == 1 {
				out.WriteString("#")
			} else {
				out.WriteString(" ")
			}
		}
		out.WriteString("\n")
	}

	return out.String()
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

	result := im.resolve()

	fmt.Println(result.String())
}
