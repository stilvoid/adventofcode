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
			outRow[j] = g.visLeft(j, i) * g.visRight(j, i) * g.visTop(j, i) * g.visBottom(j, i)
		}
		out[i] = outRow
	}
	return out
}

func (g grid) visLeft(x, y int) int {
	if x == 0 {
		return 0
	}

	dist := 0
	for i := x - 1; i >= 0; i-- {
		dist++
		if g[y][i] >= g[y][x] {
			return dist
		}
	}

	return dist
}

func (g grid) visTop(x, y int) int {
	if y == 0 {
		return 0
	}

	dist := 0
	for i := y - 1; i >= 0; i-- {
		dist++
		if g[i][x] >= g[y][x] {
			return dist
		}
	}

	return dist
}

func (g grid) visRight(x, y int) int {
	if x == len(g[y])-1 {
		return 0
	}

	dist := 0
	for i := x + 1; i < len(g[y]); i++ {
		dist++
		if g[y][i] >= g[y][x] {
			return dist
		}
	}

	return dist
}

func (g grid) visBottom(x, y int) int {
	if y == len(g)-1 {
		return 0
	}

	dist := 0
	for i := y + 1; i < len(g); i++ {
		dist++
		if g[i][x] >= g[y][x] {
			return dist
		}
	}

	return dist
}

func (g grid) Max() int {
	max := 0
	for _, row := range g {
		for _, v := range row {
			if v > max {
				max = v
			}
		}
	}
	return max
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

	fmt.Println(vis.Max())
}
