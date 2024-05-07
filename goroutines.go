package main

import (
	"fmt"
	"time"
)

func main() {
	x := 30
	y := 30
	fmt.Println(calculateFibonnaci(&x))
	fmt.Println(sumOfSquares(&y))
	// fmt.Println(difficultComputation(&x, &y))
	start := time.Now()
	result := difficultComputation(&x, &y)
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("The difficult result is %v and it took %v nanoseconds.\n", result, duration.Nanoseconds())
}

func calculateFibonnaci(n *int) int {
	var fibSlice = []int{0, 1}
	for i := len(fibSlice); i <= *n; i++ {
		num := fibSlice[len(fibSlice)-1] + fibSlice[len(fibSlice)-2]
		fibSlice = append(fibSlice, num)
	}
	return fibSlice[len(fibSlice)-1]
}

func sumOfSquares(n *int) int {
	var total int = 0
	for i := 1; i <= *n; i++ {
		total += (i * i)
	}
	return total
}

func difficultComputation(x *int, y *int) int {
	return sumOfSquares(x) + calculateFibonnaci(y)
}
