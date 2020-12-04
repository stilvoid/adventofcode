package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type direction int
type section string

const (
	up direction = iota
	down
	left
	right
	straight
)

const (
	horizontal    section = "-"
	vertical              = "|"
	clockwise             = "/"
	anticlockwise         = `\`
	intersection          = "+"
)

type truck struct {
	dir  direction
	next direction
}

type point struct {
	x, y int
}

var track map[point]section
var trucks map[point]truck

var w, h = 0, 0

func newTruck(dir direction) truck {
	return truck{
		dir:  dir,
		next: left,
	}
}

func turn_clockwise(dir direction) direction {
	switch dir {
	case up:
		return right
	case right:
		return up
	case down:
		return left
	case left:
		return down
	}

	panic("Not a direction we can rotate: " + string(dir))
}

func turn_anticlockwise(dir direction) direction {
	switch dir {
	case up:
		return left
	case left:
		return up
	case down:
		return right
	case right:
		return down
	}

	panic("Not a direction we can rotate: " + string(dir))
}

func init() {
	track = make(map[point]section)
	trucks = make(map[point]truck)

	data, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(data), "\n")

	for y, line := range lines {
		h = y
		parts := strings.Split(line, "")

		for x, part := range parts {
			w = x
			pos := point{x, y}

			switch part {
			case ">":
				track[pos] = horizontal
				trucks[pos] = newTruck(right)
			case "<":
				track[pos] = horizontal
				trucks[pos] = newTruck(left)
			case "^":
				track[pos] = vertical
				trucks[pos] = newTruck(up)
			case "v":
				track[pos] = vertical
				trucks[pos] = newTruck(down)
			case "-":
				track[pos] = horizontal
			case "|":
				track[pos] = vertical
			case "+":
				track[pos] = intersection
			case "/":
				track[pos] = clockwise
			case `\`:
				track[pos] = anticlockwise
			default:
				continue
			}
		}
	}
}

func render() {
	for y := 0; y <= h; y++ {
		for x := 0; x <= w; x++ {
			pos := point{x, y}

			if truck, ok := trucks[pos]; ok {
				switch truck.dir {
				case up:
					fmt.Print("^")
				case down:
					fmt.Print("v")
				case left:
					fmt.Print("<")
				case right:
					fmt.Print(">")
				}

				continue
			}

			if sec, ok := track[pos]; ok {
				fmt.Print(sec)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	nextTrucks := make(map[point]truck)

	for {
		for y := 0; y <= h; y++ {
			for x := 0; x <= w; x++ {
				pos := point{x, y}

				if truck, ok := trucks[pos]; ok {
					newPos := pos

					switch truck.dir {
					case up:
						newPos.y--
					case down:
						newPos.y++
					case left:
						newPos.x--
					case right:
						newPos.x++
					}

					// Did we crash?
					if _, ok := nextTrucks[newPos]; ok {
						fmt.Println("CRASH!")

						fmt.Println(newPos)

						return
					}

					sec := track[newPos]

					if sec == clockwise {
						truck.dir = turn_clockwise(truck.dir)
					} else if sec == anticlockwise {
						truck.dir = turn_anticlockwise(truck.dir)
					} else if sec == intersection {
						switch truck.next {
						case left:
							truck.next = straight
							truck.dir = turn_anticlockwise(truck.dir)
						case straight:
							truck.next = right
						case right:
							truck.next = left
							truck.dir = turn_clockwise(truck.dir)
						}
					}

					nextTrucks[newPos] = truck
				}
			}
		}

		trucks = nextTrucks
		nextTrucks = make(map[point]truck)

		render()
		time.Sleep(50 * time.Millisecond)
	}
}
