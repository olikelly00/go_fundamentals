// Write a function tripleFibonacci that
// concurrently calculates the fibonacci result
// for three numbers at once, and then returns
// them as a slice of ints.

package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := make([]int, 0)
	tripleFibonacci(10, 20, 30, &numbers)
	time.Sleep(50 * time.Microsecond)
	fmt.Println(numbers)

}

func fibonacci(n int, results *[]int) {
	var fibSlice = []int{0, 1}
	for i := len(fibSlice); i <= n; i++ {
		num := fibSlice[len(fibSlice)-1] + fibSlice[len(fibSlice)-2]
		fibSlice = append(fibSlice, num)
	}
	*results = append(*results, fibSlice[len(fibSlice)-1])
}

func tripleFibonacci(a int, b int, c int, numbers *[]int) {
	go fibonacci(a, numbers)
	go fibonacci(b, numbers)
	go fibonacci(c, numbers)
}
