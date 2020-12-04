package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input")
	boxes := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, a := range boxes {
		for _, b := range boxes {
			diffCount := 0

			same := ""

			for i, achar := range a {
				bchar := rune(b[i])

				if achar != bchar {
					diffCount++
				} else {
					same = same + string(achar)
				}
			}

			if diffCount == 1 {
				fmt.Println(same)
				return
			}
		}
	}
}
