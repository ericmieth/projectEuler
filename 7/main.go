package main

import (
	"fmt"
)

func isPrime(i int) bool {

	j := 2

	for j <= i/2 {
		if i%j == 0 {
			return false
		}
		j++
	}
	return true
}

func main() {

	limit := 10001
	counter := 2
	primeCounter := 0

	for {
		if isPrime(counter) {
			primeCounter++
		}
		if primeCounter == limit {
			break
		}
		counter++
	}

	fmt.Println(counter)
}
