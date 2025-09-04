package main

import (
	"fmt"
	"time"
)

// Ensures that data is properly exchanged between Goroutines
//
//
//

// ======== SYNCHRONIZING DATA EXCHANGE
func synchronizing_data_exchange() {
	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "hello " + string('0'+i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()
	// close(data) // Channel closed before Goroutine could send a value to the channel

	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	}
}

// ========== SYNCHRONIZING MULTIPLE GOROUTINES
func synchronizing_signals() {
	numGoroutines := 3
	done := make(chan int, 3)

	for i := range numGoroutines {
		go func(id int) {
			fmt.Printf("Goroutine %d working...\n", id)
			time.Sleep(1 * time.Second)
			done <- id // SENDING SIGNAL OF COMPLETION
		}(i)
	}

	for range numGoroutines {
		<-done // Wait for each goroutine to finish, WAIT FOR ALL TO SIGNAL COMPLETION
	}

	fmt.Println("All Goroutines are complete")
}

func simple_function() {
	done := make(chan struct{})

	// Send channel
	go func() {
		fmt.Println("Working...")
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	<-done
	fmt.Println("Finished.")
}

func simple_function2() {
	ch := make(chan int)

	go func() {
		ch <- 9 // Blocking until the value is received
		time.Sleep(1 * time.Second)
		fmt.Println("Sent value")
	}()

	value := <-ch // Blocking until a value is sent
	fmt.Println(value)
}
