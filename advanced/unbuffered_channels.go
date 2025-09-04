package main

import (
	"fmt"
	"time"
)

func UnbufferedChannles() {
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(2 * time.Second)
		fmt.Println("2 Second goroutine finished")
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("3 Second goroutine finished")
	}()

	receiver := <-ch
	fmt.Println(receiver)
	fmt.Println("End of program")
}
