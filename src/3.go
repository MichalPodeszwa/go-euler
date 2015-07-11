package main

import (
	"fmt"
	"math"
)

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

func main() {
	in_number := 600851475143
	numbers := make(chan int)
	limit := int(math.Sqrt(float64(in_number)))
	go findPrimes(numbers, limit)
	for i := range numbers {
		if in_number%i == 0 {
			fmt.Println(i)
		}
	}
}
