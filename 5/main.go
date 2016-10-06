package main

import (
	"fmt"
)

func div(i int, d int) bool {
	if i%d == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	divisor := 20
	counter := 20

	var exit bool
	var res int

	for {
		i := divisor
		for i > 1 {
			if !div(counter, i) {
				break
			}
			if i == 2 {
				res = counter
				exit = true
			}
			i--
		}
		if exit {
			break
		}
		counter++
	}

	fmt.Println(res)
}
