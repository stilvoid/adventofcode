package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type arg struct {
	stone, count int
}

var memo = make(map[arg]int)

func iter(stone int, count int) int {
	r := arg{stone, count}

	if result, ok := memo[r]; ok {
		return result
	}

	if count == 0 {
		return 1
	}

	str := fmt.Sprint(stone)

	switch {
	case stone == 0:
		memo[r] = iter(1, count-1)
	case len(str)%2 == 0:
		size := len(str) / 2
		a, _ := strconv.Atoi(str[:size])
		b, _ := strconv.Atoi(str[size:])
		memo[r] = iter(a, count-1) + iter(b, count-1)
	default:
		memo[r] = iter(stone*2024, count-1)
	}

	return memo[r]
}

func main() {
	data, _ := os.ReadFile("data")

	parts := strings.Split(strings.TrimSpace(string(data)), " ")

	stones := make([]int, len(parts))

	for i, part := range parts {
		stones[i], _ = strconv.Atoi(part)
	}

	fmt.Println(stones)

	total := 0
	for _, stone := range stones {
		total += iter(stone, 75)
	}
	fmt.Println(total)
}
