package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	x, y int
}

func newRule(in string) rule {
	parts := strings.Split(in, "|")

	r := rule{}

	r.x, _ = strconv.Atoi(parts[0])
	r.y, _ = strconv.Atoi(parts[1])

	return r
}

func (r rule) ok(u update) bool {
	xpos, ok := u.order[r.x]
	if !ok {
		return true
	}

	ypos, ok := u.order[r.y]
	if !ok {
		return true
	}

	return xpos < ypos
}

type ruleset []rule

func (rs ruleset) ok(u update) bool {
	for _, rule := range rs {
		if !rule.ok(u) {
			return false
		}
	}
	return true
}

type update struct {
	order map[int]int
	list  []int
}

func newUpdate(in string) update {
	parts := strings.Split(in, ",")

	u := update{
		order: make(map[int]int),
		list:  make([]int, len(parts)),
	}

	for i, part := range parts {
		p, _ := strconv.Atoi(part)

		u.order[p] = i
		u.list[i] = p
	}

	return u
}

func main() {
	// Parse inputs
	data, _ := os.ReadFile("data")

	parts := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	ruleInputs := strings.Split(parts[0], "\n")
	updateInputs := strings.Split(parts[1], "\n")

	rules := make(ruleset, 0)
	for _, in := range ruleInputs {
		rules = append(rules, newRule(in))
	}

	updates := make([]update, 0)
	for _, in := range updateInputs {
		updates = append(updates, newUpdate(in))
	}

	// Check rules and collect good updates
	good := make([]update, 0)
	for _, update := range updates {
		if rules.ok(update) {
			good = append(good, update)
		}
	}

	// Add up the middles
	total := 0
	for _, u := range good {
		n := u.list[len(u.list)/2]
		fmt.Println(n)
		total += n
	}
	fmt.Println(total)
}
