package main

import (
	"fmt"
	"time"
)

// Buffered channel - a channel with storage
// useful for - asynchronous communication, load balancing, flow control

// ========= BLOCKING ON SEND ONLY IF THE BUFFER IS FULL
func FIRST_CASE() {
	// make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("receiving from buffer")
	go func() {
		fmt.Println("Goroutine 2 second timer started.")
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch)
	}()
	fmt.Println("Blocking starts here")
	ch <- 3
	fmt.Println("Blocking ends here")
	// fmt.Println("Received:", <-ch)
	// fmt.Println("Received:", <-ch)
	// fmt.Println("Buffered Channels")
}

// ========= BLOCKING ON RECEIVE ONLY CHANNEL IF THE BUFFER IS EMPTY
func SECOND_CASE() {
	ch := make(chan int, 2)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		ch <- 1
	}()

	fmt.Println("Value: ", <-ch)
	fmt.Println("Value: ", <-ch)
	fmt.Println("End of program")
}
