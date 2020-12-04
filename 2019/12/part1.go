package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	for t := 1; t <= 1000; t++ {
		for i, im := range moons {
			for j, jm := range moons {
				if i == j {
					continue
				}

				dx := im.pos.x - jm.pos.x
				if dx < 0 {
					im.vel.x++
				} else if dx > 0 {
					im.vel.x--
				}

				dy := im.pos.y - jm.pos.y
				if dy < 0 {
					im.vel.y++
				} else if dy > 0 {
					im.vel.y--
				}

				dz := im.pos.z - jm.pos.z
				if dz < 0 {
					im.vel.z++
				} else if dz > 0 {
					im.vel.z--
				}
			}
		}

		for _, m := range moons {
			m.pos.x += m.vel.x
			m.pos.y += m.vel.y
			m.pos.z += m.vel.z
		}

		fmt.Printf("After %d steps:\n", t)

		for _, m := range moons {
			fmt.Println(m.String())
		}

		fmt.Println()
	}

	// Calculate energy
	total := float64(0)
	for _, m := range moons {
		p := math.Abs(float64(m.pos.x)) + math.Abs(float64(m.pos.y)) + math.Abs(float64(m.pos.z))
		k := math.Abs(float64(m.vel.x)) + math.Abs(float64(m.vel.y)) + math.Abs(float64(m.vel.z))

		total += p * k
	}

	fmt.Println(total)
}
