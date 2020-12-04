package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type point struct {
	x, y, z int
}

func (p point) String() string {
	return fmt.Sprintf("<x=%d, y=%d, z=%d>", p.x, p.y, p.z)
}

type moon struct {
	pos point
	vel point
}

func (m moon) String() string {
	return fmt.Sprintf("pos=%s, vel=%s", m.pos, m.vel)
}

func findPeriod(init ...int) int64 {
	pos := make([]*int, len(init))
	vel := make([]*int, len(init))
	initVel := make([]*int, len(init))

	for i, p := range init {
		newP := p
		newVel := 0
		newInitVel := 0

		pos[i] = &newP
		vel[i] = &newVel
		initVel[i] = &newInitVel
	}

	var count int64

loop:
	for {
		// Calculate velocities
		for i, ip := range pos {
			for j, jp := range pos {
				if i == j {
					continue
				}

				dx := *ip - *jp
				if dx > 0 {
					*(vel[i])--
				} else if dx < 0 {
					*(vel[i])++
				}
			}
		}

		// Apply velocities
		for i, ip := range pos {
			*ip += *(vel[i])
		}

		count++

		for i, ip := range pos {
			if *ip != init[i] {
				continue loop
			}

			if *(vel[i]) != *(initVel[i]) {
				continue loop
			}
		}

		return count
	}
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	moons := make([]*moon, 0)

	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		var x, y, z int

		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &x, &y, &z)

		moons = append(moons, &moon{
			pos: point{x, y, z},
		})
	}

	xs := make([]int, len(moons))
	ys := make([]int, len(moons))
	zs := make([]int, len(moons))

	for i, m := range moons {
		xs[i] = m.pos.x
		ys[i] = m.pos.y
		zs[i] = m.pos.z
	}

	xp := findPeriod(xs...)
	yp := findPeriod(ys...)
	zp := findPeriod(zs...)

	fmt.Println(xp, yp, zp)

	// Find the common period
	p := xp * yp * zp
	d := int64(2)
	for p%d == 0 {
		d += 2
	}
	fmt.Println(p / (d + 2))
}
