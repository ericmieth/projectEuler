package main

import (
	"fmt"
)

func fab(a, b int) (int, int) {
	return b, a + b
}

func main() {
	limit := 4000000

	a := 1
	b := 1

	sum := 0

	for b < limit {
		a, b = fab(a, b)
		if b%2 == 0 {
			sum += b
		}
	}

	fmt.Println(sum)
}
