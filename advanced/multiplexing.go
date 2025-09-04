package main

import (
	"fmt"
	"time"
)

// The process of handling multiple channel operations simultaneously
// Uses Select statement
//

func simple_time() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		close(ch)
	}()

	// for {
	// 	select {
	// 	case msg, ok := <-ch:
	// 		if !ok {
	// 			fmt.Println("Channel closed")
	// 			// clean up activities
	// 			return
	// 		}
	// 		fmt.Println("Received:", msg)
	// 	}
	// }

	select {
	case msg, ok := <-ch:
		if !ok {
			fmt.Println("Channel closed")
			return
		}
		fmt.Println("Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout.")
	}
}

func simple_select() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		// time.Sleep(time.Second)
		ch1 <- 2
	}()

	time.Sleep(2 * time.Second)
	for range 2 {
		select {
		case msg := <-ch1:
			fmt.Println("Received from ch1: ", msg)
		case msg := <-ch2:
			fmt.Println("Received from ch2: ", msg)
		default:
			fmt.Println("No channels ready.")
		}
	}

	fmt.Println("End of program")

}
