package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type grid [][]int

func (g grid) String() string {
	out := strings.Builder{}
	for _, row := range g {
		for _, v := range row {
			out.WriteString(fmt.Sprint(v))
		}
		out.WriteString("\n")
	}
	return out.String()
}

func (g grid) visMap() grid {
	out := make(grid, len(g))
	for i, row := range g {
		outRow := make([]int, len(row))
		for j := range row {
			if g.visLeft(i, j) || g.visRight(i, j) || g.visTop(i, j) || g.visBottom(i, j) {
				outRow[j] = 1
			} else {
				outRow[j] = 0
			}
		}
		out[i] = outRow
	}
	return out
}

func (g grid) visLeft(x, y int) bool {
	if x == 0 {
		return true
	}

	for i := 0; i < x; i++ {
		if g[y][i] >= g[y][x] {
			return false
		}
	}

	return true
}

func (g grid) visRight(x, y int) bool {
	if x == len(g[y])-1 {
		return true
	}

	for i := x + 1; i < len(g[y]); i++ {
		if g[y][i] >= g[y][x] {
			return false
		}
	}

	return true
}

func (g grid) visTop(x, y int) bool {
	if y == 0 {
		return true
	}

	for i := 0; i < y; i++ {
		if g[i][x] >= g[y][x] {
			return false
		}
	}

	return true
}

func (g grid) visBottom(x, y int) bool {
	if y == len(g)-1 {
		return true
	}

	for i := y + 1; i < len(g); i++ {
		if g[i][x] >= g[y][x] {
			return false
		}
	}

	return true
}

func main() {
	input, _ := os.ReadFile("input")

	trees := make(grid, 0)
	for _, line := range strings.Split(string(input), "\n") {
		row := make([]int, 0)
		for _, c := range line {
			v, _ := strconv.Atoi(string(c))
			row = append(row, v)
		}
		trees = append(trees, row)
	}

	fmt.Println(trees)

	vis := trees.visMap()

	fmt.Println(vis)

	count := 0
	for _, row := range vis {
		for _, v := range row {
			count += int(v)
		}
	}

	fmt.Println(count)
}
