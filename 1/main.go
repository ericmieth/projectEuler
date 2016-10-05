package main

import (
	"fmt"
)

func multiple(i int) bool {
	if i%3 == 0 || i%5 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	limit := 1000
	count := 1

	sum := 0

	for count < limit {
		if multiple(count) {
			sum += count
		}

		count++
	}

	fmt.Println(sum)
}
