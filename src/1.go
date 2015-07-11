package main

import "fmt"

func multiples(c chan int) {
	defer close(c)
	for i := 0; i < 30000000; i++ {
		if i%3 == 0 || i%5 == 0 {
			c <- i
		}
	}
}

func main() {
	nums := make(chan int, 2)
	sum := 0
	go multiples(nums)
	for i := range nums {
		sum += i
	}
	fmt.Println(sum)
}
