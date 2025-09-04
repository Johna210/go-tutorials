package main

import "fmt"

// Channels enable safe and efficient communication between concurrent
// Goroutines

func channels() {
	// variable := make(chan type)
	greeting := make(chan string)
	greetString := "Hello"

	go func() {
		// blocking because its continuously trying to receive values
		// its ready to receive continuous flow of data.
		greeting <- greetString
		greeting <- "World"
	}()

	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)
}
