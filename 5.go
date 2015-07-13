package main

import "fmt"

func isDone(numbers []int) bool {
	for _, i := range numbers {
		if i != 1 {
			return false
		}
	}
	return true
}

func Fifth() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	primes := make(chan int)
	go findPrimes(primes, len(numbers))

	output := 1

	for prime := range primes {
		for !isDone(numbers) {
			wasUsed := false
			for i, n := range numbers {
				if n%prime == 0 {
					if !wasUsed {
						output *= prime
					}
					numbers[i] /= prime
					wasUsed = true
				}
			}
			if wasUsed {
				continue
			} else {
				break
			}
		}
	}

	fmt.Println(output)
}
