package main

import (
	"fmt"
)

func Fourth() {
	max := 0
	numbers := make(chan int)
	go findPalindromesMult(numbers, 999)
	for i := range numbers {
		if i > max {
			max = i
		}
	}
	fmt.Println(max)
}
