package main

import (
	"fmt"
	"time"
)

func closing_channel() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer2(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println(val)
	}
}

func producer2(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int) {
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}
	close(out)
}

// panics
func closing_channel_twice() {
	ch := make(chan int)
	go func() {
		close(ch)
		close(ch)
	}()
	time.Sleep(1 * time.Second)

}

func range_over_closed_channel() {
	ch := make(chan int)
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println("Received: ", val)
	}
}

func receiving_from_closed_channel() {
	ch := make(chan int)
	close(ch)

	val, ok := <-ch
	if !ok {
		fmt.Println("Channel is closed")
		return
	}

	fmt.Println(val)
}

func simple_closing_channel() {
	ch := make(chan int)
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}
}
