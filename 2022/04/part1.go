package main

import (
	"fmt"
	"io"
	"os"
)

type section struct {
	from, to int
}

func (s section) contains(other section) bool {
	return s.from <= other.from && other.to <= s.to
}

func (s section) overlaps(other section) bool {
	return s.from <= other.to && other.to <= s.to
}

type pair struct {
	left, right section
}

func (p pair) fullyContained() bool {
	return p.left.contains(p.right) || p.right.contains(p.left)
}

func (p pair) overlaps() bool {
	return p.left.overlaps(p.right) || p.right.overlaps(p.left)
}

func getPairs(r io.Reader) []pair {
	pairs := make([]pair, 0)

	for {
		var a, b, c, d int
		_, err := fmt.Fscanf(r, "%d-%d,%d-%d", &a, &b, &c, &d)
		if err == io.EOF {
			fmt.Println("EOF")
			return pairs
		} else if err != nil {
			panic(err)
		}

		fmt.Println("Adding pair", a, b, c, d)
		pairs = append(pairs, pair{
			left:  section{a, b},
			right: section{c, d},
		})
	}
}

func main() {
	f, _ := os.Open("input")
	pairs := getPairs(f)
	fmt.Println(pairs)

	containCount := 0
	overlapCount := 0
	for _, p := range pairs {
		if p.overlaps() {
			overlapCount++

			if p.fullyContained() {
				containCount++
			}
		}
	}

	fmt.Println("Contained:", containCount)
	fmt.Println("Overlapped:", overlapCount)
}
