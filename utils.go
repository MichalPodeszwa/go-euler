package main

import "strconv"

func findPrimes(c chan int, limit int) {
	defer close(c)
	numbers := make(map[int]bool)
	for i := 2; i < limit; i++ {
		if _, prs := numbers[i]; prs {
			continue
		}
		c <- i
		numbers[i] = true
		for n := i * i; n < limit; n += i {
			numbers[n] = false
		}
	}
}

func fib(c chan int) {
	defer close(c)
	prev, cur := 1, 1
	for {
		c <- cur
		prev, cur = cur, prev+cur
		if cur > 4e6 {
			return
		}
	}
}

func isPalindrome(number int) bool {
	strnum := strconv.Itoa(number)
	for i := 0; i < len(strnum)/2; i++ {
		if strnum[i] != strnum[len(strnum)-i-1] {
			return false
		}
	}
	return true
}

func findPalindromesMult(c chan int, limit int) {
	defer close(c)
	for x := limit; x > 0; x-- {
		for y := x; y > 0; y-- {
			number := x * y
			if isPalindrome(number) {
				c <- number
			}
		}
	}
}
