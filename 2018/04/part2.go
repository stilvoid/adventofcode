package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

var stampRe *regexp.Regexp
var beginRe *regexp.Regexp
var sleepRe *regexp.Regexp
var awakeRe *regexp.Regexp

func init() {
	stampRe = regexp.MustCompile(`^\[\d\d\d\d-\d\d-\d\d \d\d:(?P<minute>\d\d)\] (?P<message>.+)$`)
	beginRe = regexp.MustCompile(`^Guard #(\d+) begins shift`)
	sleepRe = regexp.MustCompile(`^falls asleep`)
	awakeRe = regexp.MustCompile(`^wakes up`)
}

type guard struct {
	totalSleep     int
	sleepPerMinute map[int]int
	lastAction     int
}

func newGuard() *guard {
	guard := &guard{
		sleepPerMinute: make(map[int]int),
	}

	return guard
}

func main() {
	data, _ := ioutil.ReadFile("input")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	sort.Strings(lines)

	guards := make(map[string]*guard)
	var currentGuard string

	for _, line := range lines {
		var minute int

		// Parse out the things
		submatches := stampRe.FindStringSubmatch(line)
		fmt.Sscanf(submatches[1], "%d", &minute)
		message := submatches[2]

		switch {
		case beginRe.MatchString(message):
			currentGuard = beginRe.FindStringSubmatch(message)[1]
			if _, ok := guards[currentGuard]; !ok {
				guards[currentGuard] = newGuard()
			}
		case sleepRe.MatchString(message):
			guards[currentGuard].lastAction = minute
		case awakeRe.MatchString(message):
			for i := guards[currentGuard].lastAction; i < minute; i++ {
				guards[currentGuard].sleepPerMinute[i]++
				guards[currentGuard].totalSleep++
			}
		}
	}

	// Find sleepiest guard
	sleepiestCount := 0
	sleepiestGuard := ""
	sleepiestMinute := 0
	for id, guard := range guards {
		for minute := 0; minute < 60; minute++ {
			count := guard.sleepPerMinute[minute]

			if count > sleepiestCount {
				sleepiestCount = count
				sleepiestGuard = id
				sleepiestMinute = minute
			}
		}
	}

	fmt.Println(sleepiestGuard, sleepiestMinute)

	var n int
	fmt.Sscanf(sleepiestGuard, "%d", &n)
	fmt.Println(n * sleepiestMinute)
}
