package main

import (
	"fmt"
)

func quad(i int) int {
	return i * i
}

func main() {

	counter := 1

	sum1 := 0
	sum2 := 0

	for counter <= 100 {
		sum1 += quad(counter)
		sum2 += counter
		counter++
	}

	res := quad(sum2) - sum1
	fmt.Println(res)
}
