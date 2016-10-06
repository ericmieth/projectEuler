package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(i int) bool {
	str := strconv.Itoa(i)
	if str[0] == str[len(str)-1] &&
		str[1] == str[len(str)-2] &&
		str[2] == str[len(str)-3] {
		return true
	} else {
		return false
	}
}

func main() {

	a := 100

	largestPalindrome := 0
	var number int

	for a < 1000 {
		b := 100
		for b < 1000 {
			number = a * b
			if isPalindrome(number) && number > largestPalindrome {
				largestPalindrome = number
			}
			b++
		}
		a++
	}
	fmt.Println(largestPalindrome)
}
