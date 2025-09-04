package main

import "fmt"

// Specify the intended operations on channels
// Enhance type safety by clearly defining the channel's purpose
// improves code clarity and maintainability

func channel_direction() {
	ch := make(chan int)
	producer(ch)
	consumer(ch)
}

// Send only channel
func producer(ch chan<- int) {
	go func(ch chan<- int) {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}(ch)
}

// Receive only channel
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received: ", value)
	}
}
