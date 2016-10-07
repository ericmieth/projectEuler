package main

import (
	"fmt"
)

type TriangleNumber struct {
	n int // n-th triagle number
	v int // value of triangle number
}

func (t TriangleNumber) nextTriangleNumber() TriangleNumber {
	t.n++
	t.v += t.n
	return t
}

func (t TriangleNumber) countDivisors() int {

	divisors := 1

	for i := 1; i <= int(t.v/2); i++ {
		if t.v%i == 0 {
			divisors++
		}
	}

	return divisors
}

func main() {

	limit := 500
	t := TriangleNumber{1, 1}

	var divisors int

	for {
		divisors = t.countDivisors()
		if divisors > limit {
			break
		}
		t = t.nextTriangleNumber()
	}

	fmt.Println(t.v)
}
