//Develop a Go program that launches 10 concurrent goroutines.
//Each goroutine should generate and send 10 unique numbers
//through a shared channel. Implement a mechanism to retrieve
//these numbers from the channel and print them out.

//This exercise involves creating a channel, spawning multiple
//goroutines that populate the channel with data, and consuming
//that data from the channel in the main goroutine.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// numbers := make(chan int)
	// for i := 0; i < 10; i++ {
	// 	go AddNumbersToChannel(numbers)
	// }
	// ListenToNumbersFromChannel(numbers)
	//time.Sleep(1 * time.Second)
	myTime := time.Now()
	formatted_time := myTime.Format("15:04 - 02, January, 2006")
	fmt.Printf("The time is %s\n", formatted_time)


}

func AddNumbersToChannel(numbers chan<- int) {
	randomSeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	seed := randomSeed.Int()
	for i := 1; i <= 10; i++ {
		numbers <- seed
	}
}

func ListenToNumbersFromChannel(numbers <-chan int) {
	for num := range numbers {
		fmt.Println("Number: ", num)
	}
}


