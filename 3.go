package main

import (
	"fmt"
	"math"
)

func Third() {
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
