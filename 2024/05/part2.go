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

func (r rule) apply(u *update) bool {
	xpos, ok := u.order[r.x]
	if !ok {
		return false
	}

	ypos, ok := u.order[r.y]
	if !ok {
		return false
	}

	if xpos < ypos {
		return false
	}

	u.list[xpos], u.list[ypos] = u.list[ypos], u.list[xpos]
	u.order[r.x] = ypos
	u.order[r.y] = xpos
	return true
}

type ruleset []rule

func (rs ruleset) apply(u *update) bool {
	changed := false

	for _, rule := range rs {
		changed = changed || rule.apply(u)
	}

	return changed
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

	// Apply rules to updates
	changed := make([]update, 0)
	for _, update := range updates {
		take := false
		for rules.apply(&update) {
			take = true
		}
		if take {
			changed = append(changed, update)
		}
	}

	// Add up the middles
	total := 0
	for _, u := range changed {
		n := u.list[len(u.list)/2]
		fmt.Println(n)
		total += n
	}
	fmt.Println(total)
}
