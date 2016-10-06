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
	number := 600851475143
	var largestPrimeFactor int

	i := 2
	for i <= number/2 {
		j := 2
		for {
			if div(number, j) {
				largestPrimeFactor = j
				number = number / j
				break
			}
			j++
		}
		i++
	}

	fmt.Println(largestPrimeFactor)
}
