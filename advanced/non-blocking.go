package main

import (
	"fmt"
	"time"
)

func non_blocking() {

	// ==== NON BLOCKING RECEIVE OPERATIONS
	// ch := make(chan int)

	// 1.
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received msg", msg)
	// default:
	// 	fmt.Println("No messages available")
	// }

	// ==== NON BLOCKING SEND OPERATION
	// select {
	// case ch <- 1:
	// 	fmt.Println("Sent message.")
	// default:
	// 	fmt.Println("Channel is not ready to receive.")
	// }

	// ==== NON BLOCKING OPERATION IN REAL TIME SYSTEMS - graceful shutdown
	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received:", d)
			case <-quit:
				fmt.Println("Stopping...")
				return
			default:
				fmt.Println("Waiting for data...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(1 * time.Second)
	}

	quit <- true
}
