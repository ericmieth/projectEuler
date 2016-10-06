package main

import (
	"fmt"
)

func quad(i int) int {
	return i * i
}

func main() {

	limit := 1000

	var result int
	exit := false

	a := 1
	for a < limit {
		b := 1
		for b < limit {
			c := 1
			for c < limit {
				if quad(a)+quad(b) == quad(c) && a+b+c == limit {
					result = a * b * c
					exit = true
				}
				if exit {
					break
				}
				c++
			}
			if exit {
				break
			}
			b++
		}
		if exit {
			break
		}
		a++
	}

	fmt.Println(result)
}
