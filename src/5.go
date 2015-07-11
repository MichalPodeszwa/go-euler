package main

import "fmt"

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

func isDone(numbers []int) bool {
	for _, i := range numbers {
		if i != 1 {
			return false
		}
	}
	return true
}

func main() {
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
