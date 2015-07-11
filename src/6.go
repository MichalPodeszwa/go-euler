package main

import "fmt"

func main() {
	sumSquares := 0
	squaresSum := 0
	for i := 0; i <= 100; i++ {
		sumSquares += i * i
		squaresSum += i
	}
	squaresSum *= squaresSum
	fmt.Println(squaresSum - sumSquares)
}
