package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x, y int
}

type trailmap map[pos]map[pos]int

func (tm trailmap) route(height int, start pos) int {
	count := 0

	for next, nextHeight := range tm[start] {
		if nextHeight == height+1 {
			if nextHeight == 9 {
				count++
			} else {
				count += tm.route(nextHeight, next)
			}
		}
	}

	return count
}

func main() {
	data, _ := os.ReadFile("data")

	tm := make(trailmap)

	nums := make([][]int, 0)
	for y, row := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		nums = append(nums, make([]int, 0))

		for _, char := range row {
			n, err := strconv.Atoi(string(char))
			if err != nil {
				n = -1
			}

			nums[y] = append(nums[y], n)
		}
	}

	heads := make([]pos, 0)

	for y, row := range nums {
		for x, height := range row {
			p := pos{x, y}

			if height == 0 {
				heads = append(heads, p)
			}

			tm[p] = make(map[pos]int)

			// Up
			if y > 0 {
				n := nums[y-1][x]
				if n != -1 {
					tm[p][pos{x, y - 1}] = n
				}
			}

			// Down
			if y < len(nums)-1 {
				n := nums[y+1][x]
				if n != -1 {
					tm[p][pos{x, y + 1}] = n
				}
			}

			// Left
			if x > 0 {
				n := nums[y][x-1]
				if n != -1 {
					tm[p][pos{x - 1, y}] = n
				}
			}

			// Right
			if x < len(row)-1 {
				n := nums[y][x+1]
				if n != -1 {
					tm[p][pos{x + 1, y}] = n
				}
			}
		}
	}

	fmt.Println(tm)

	fmt.Println(heads)

	total := 0
	for _, head := range heads {
		total += tm.route(0, head)
	}

	fmt.Println(total)
}
