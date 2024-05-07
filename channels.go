// package main

// import "fmt"

// func main() {
// 	messages := make(chan string, 2) // (The 2 here indicates that the channel can store up to two messages at a time.)

// 	// Sending values into a channel
// 	messages <- "Hey"
// 	messages <- "World"

// 	// Reading values out of a channel.
// 	msg := <-messages
// 	fmt.Println("msg:", msg)
// 	msg = <-messages
// 	fmt.Println("msg:", msg)
// }

package main

import "fmt"

func main() {
	messages := make(chan string)

	go SayHello(messages)

	// Reading values out of a channel.
	msg := <-messages
	fmt.Println("msg:", msg)

	msg = <-messages
	fmt.Println("msg:", msg)
}

func SayHello(ch chan<- string) { // means that ch is a string channel, and that this function will send values into that channel. The opposite would be a function that will read values from a channel. The type for that would be ch <-chan string
	ch <- "Hello"
	ch <- "world"
}
