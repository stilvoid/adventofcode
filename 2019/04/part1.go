package main

import "fmt"

func main() {
	start := 248345
	end := 746315

	var a string
	var i int

	var valid = struct {
		double   bool
		increase bool
	}{}

	var count = 0

	for n := start; n <= end; n++ {
		a = fmt.Sprint(n)
		valid.double = false
		valid.increase = true

		// Apply rules
		for i = 0; i < 5; i++ {
			// Two adjacent digits are the same
			if !valid.double {
				if a[i] == a[i+1] {
					valid.double = true
				}
			}

			if a[i+1] < a[i] {
				valid.increase = false
				break
			}
		}

		if valid.double && valid.increase {
			count++
		}
	}

	fmt.Println(count)
}
