package main

import (
	"fmt"
)

type sequence struct {
	s int // start value
	v int // actual value
	l int // lenght of chain
}

func (s sequence) next() sequence {
	s.l++
	if s.v%2 == 0 {
		return sequence{s.s, s.v / 2, s.l}
	} else {
		return sequence{s.s, (3 * s.v) + 1, s.l}
	}
}
func main() {

	var longestChain sequence

	for i := 1; i < 1000000; i++ {
		s := sequence{i, i, 1}
		for {
			s = s.next()
			if s.v == 1 {
				if s.l > longestChain.l {
					longestChain = s
				}
				break
			}
		}
	}

	fmt.Println(longestChain.s)
}
