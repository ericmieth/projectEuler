package main

import (
	"fmt"
	"math"
)

func isPrime(i int) bool {

	j := 2

	for j <= int(math.Sqrt(float64(i)))+1 {
		if i%j == 0 {
			return false
		}
		j++
	}
	return true
}

func main() {

	limit := 2000000
	counter := 2

	sum := 2

	for counter < limit {
		if isPrime(counter) {
			sum += counter
		}
		counter++
	}

	fmt.Println(sum)
}
