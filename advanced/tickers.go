package main

import (
	"fmt"
	"time"
)

func ticker_main() {
	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)

	defer ticker.Stop()

	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at:", tick)
		case <-stop:
			fmt.Println("Stopping ticker.")
			return
		}
	}
}

// ==== SCHEDULED LOGGING, PERIODIC TASK, POLLING FOR UPDATES
func simple_periodic() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			periodicTask()
		}
	}
}

func periodicTask() {
	fmt.Println("Performing periodic task at:", time.Now())
}

func simple_ticker() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// for tick := range ticker.C {
	// 	fmt.Println("Tick at:", tick)
	// }

	// i := 1
	// for range 5 {
	// 	i *= 2
	// 	fmt.Println(i)
	// }

	i := 1
	for range 5 {
		i *= 2
		fmt.Println(i)
	}
}
