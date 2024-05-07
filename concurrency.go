package main

import (
	"fmt"
	"time"
)

func main() {
	go count("John")
	go count("Hafsah")
	go count("Sandy")
	time.Sleep(2 * time.Millisecond)
	
}

func count(name string) {
	for i := 0; i < 100; i++ {
		fmt.Println(name, "has counted to", i)
	}
}
