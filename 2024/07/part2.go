package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	test int
	nums []int
}

func cat(a, b int) int {
	n, _ := strconv.Atoi(fmt.Sprintf("%d%d", a, b))
	return n
}

func solve(e equation) int {
	if len(e.nums) == 1 {
		if e.nums[0] == e.test {
			return e.test
		}
		return 0
	}

	a, b := e.nums[0], e.nums[1]

	if solve(equation{test: e.test, nums: append([]int{a + b}, e.nums[2:]...)}) == e.test {
		return e.test
	}

	if solve(equation{test: e.test, nums: append([]int{a * b}, e.nums[2:]...)}) == e.test {
		return e.test
	}

	if solve(equation{test: e.test, nums: append([]int{cat(a, b)}, e.nums[2:]...)}) == e.test {
		return e.test
	}

	return 0
}

func main() {
	// Parse
	data, _ := os.ReadFile("data")

	parts := make([]equation, 0)

	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		lineParts := strings.Split(line, ": ")
		left := lineParts[0]
		right := strings.Split(lineParts[1], " ")

		l, _ := strconv.Atoi(left)

		part := equation{
			test: l,
			nums: make([]int, len(right)),
		}

		for i, num := range right {
			n, _ := strconv.Atoi(num)
			part.nums[i] = n
		}

		parts = append(parts, part)
	}

	// Solve
	total := 0
	for _, part := range parts {
		total += solve(part)
	}
	fmt.Println(total)
}
