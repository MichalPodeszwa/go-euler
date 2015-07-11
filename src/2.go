package main

import "fmt"

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

func main() {
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
