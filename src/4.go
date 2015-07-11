package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(number int) bool {
	strnum := strconv.Itoa(number)
	for i := 0; i < len(strnum)/2; i++ {
		if strnum[i] != strnum[len(strnum)-i-1] {
			return false
		}
	}
	return true
}

func findPalindromes(c chan int) {
	defer close(c)
	for x := 999; x > 0; x-- {
		for y := x; y > 0; y-- {
			number := x * y
			if isPalindrome(number) {
				c <- number
			}
		}
	}
}

func main() {
	max := 0
	numbers := make(chan int)
	go findPalindromes(numbers)
	for i := range numbers {
		if i > max {
			max = i
		}
	}
	fmt.Println(max)
}
