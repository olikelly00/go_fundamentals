// package main

// import (
// // 	"fmt"
// 	"time"
// // )

// // This function creates a copy of an array when used
// 	func doubleArray(numbers []int) {
// 		for i := 0; i < len(numbers); i++ {
// 		numbers[i] = numbers[i] * 2
// 		}
// 	}

// 	// This function takes a pointer to an array as an argument
// 	func doubleArrayP(numbers *[]int) {
// 		for i := 0; i < len(*numbers); i++ {
// 			(*numbers)[i] = (*numbers)[i] * 2
// 		}
// 	}

// 	func main() {
// 		// Declare an array
// 		ary1 := []int{1, 2, 3}
// 		ary2 := []int{4, 5, 6}

// 		// Display the original array
// 		fmt.Println("Original ary1:", ary1)
// 		fmt.Println("Original ary2:", ary2)

// 		// Modify ary1 passing ary1 directly
// 		doubleArray(ary1)
// 		// Modify ary2 passing a pointer to ary2
// 		doubleArrayP(&ary2)

// 		// Display the modified array
// 		fmt.Println("Modified array:", ary1)
// 		fmt.Println("Modified array:", ary2)
// 	}

package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := make([]int, 0)
	tripleFibonacci(10, 20, 30, &numbers)
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
	//time.Sleep(2 * time.Nanosecond)
}

// func sumOfSquares(n int, results *[]int) {
// 	var total int = 0
// 	for i := 1; i <= n; i++ {
// 		total += (i * i)
// 	}
// 	*results = append(*results, total)
// }

// func fillwithFives(sl *[]int) {
// 	for i := range *sl {
// 		(*sl)[i] = 5
// 	}
// }

// func fillwithFives() []int {
// 	var sl []int
// 	slp := &sl
// 	*slp = append(*slp, 5)
// 	return sl
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	jsonString := `["John", "Hafsah", "Sandy"]`
// 	jsonBytes := []byte(jsonString)

// 	var namesList []string
// 	err := json.Unmarshal(jsonBytes, namesList)

// 	if err != nil {
// 	fmt.Println(err)
// 	}

// 	fmt.Println(namesList)
// }

// x := 4
// fmt.Println(x)

// go sumOfSquares(30, &numbers)
// go fibonacci(30, &numbers)

// for len(numbers) < 2 {
// }
// total := numbers[0] + numbers[1]
// fmt.Println("The combined result is", total)
// fmt.Println(fillwithFives())
