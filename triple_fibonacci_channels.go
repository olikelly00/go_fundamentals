//Rewrite the tripleFibonacci function from the previous section,
//so that it uses channels instead of a shared pointer to perform
// the calculation. Send your solution to your coach.

package main

import (
	"fmt"
	//"time"
)

func main() {
	// numbers := make([]int, 0)
	fibNumbers := make(chan int)
	tripleFibonacci(fibNumbers)
	num := <-fibNumbers
	fmt.Println(num)
	num = <-fibNumbers
	fmt.Println(num)
	num = <-fibNumbers
	fmt.Println(num)

	// tripleFibonacci(10, 20, 30, &numbers)
	// fmt.Println(numbers)
}

func fibonacci(n int) int {
	var fibSlice = []int{0, 1}
	for i := len(fibSlice); i <= n; i++ {
		num := fibSlice[len(fibSlice)-1] + fibSlice[len(fibSlice)-2]
		fibSlice = append(fibSlice, num)
	}
	return fibSlice[len(fibSlice)-1]
}

func tripleFibonacci(ch chan<- int) {
	go func() { ch <- fibonacci(5) }()
	go func() { ch <- fibonacci(50) }()
	go func() { ch <- fibonacci(100) }()
}

// func tripleFibonacci(a int, b int, c int, numbers *[]int) {
// 	go fibonacci(a, numbers)
// 	go fibonacci(b, numbers)
// 	go fibonacci(c, numbers)
// 	time.Sleep(2 * time.Nanosecond)
// }
