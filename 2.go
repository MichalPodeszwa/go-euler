package main

import "fmt"

func Second() {
	nums := make(chan int)
	go fib(nums)
	sum := 0
	for i := range nums {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
