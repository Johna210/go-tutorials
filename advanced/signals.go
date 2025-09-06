package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main_signals() {
	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Notify channel on interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGHUP)

	go func() {
		sig := <-sigs
		fmt.Println("Received signal:", sig)
		done <- true
	}()

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Stopping work due to signal")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(time.Second)
			}
		}

		// sig := <-sigs
		// for sig := range sigs {
		// 	switch sig {
		// 	case syscall.SIGINT:
		// 		fmt.Println("Received SIGINT (Ctrl+C) (Interrupt)")
		// 	case syscall.SIGTERM:
		// 		fmt.Println("Received SIGTERM (Terminate)")
		// 	case syscall.SIGHUP:
		// 		fmt.Println("Received SIGHUP (Hangup)")
		// 	case syscall.SIGUSR1:
		// 		fmt.Println("Received SIGUSR1 (User-defined signal 1)")
		// 		fmt.Println("User defined function is executed")
		// 		continue
		// 	default:
		// 		fmt.Println("Received other signal")
		// 	}
		// 	fmt.Println("Graceful exit")
		// 	os.Exit(0)
		// }
	}()

	// simulate work
	// fmt.Println("Working...")
	for {
		time.Sleep(time.Second)
	}
}

func simple_signal_example() {

	pid := os.Getpid()
	fmt.Println("Process ID:", pid)
	sigs := make(chan os.Signal, 1)

	// Notify channel on interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println("Received signal:", sig)
		fmt.Println("Graceful exit")
		os.Exit(0)
	}()

	// simulate work
	fmt.Println("Working...")
	for {
		time.Sleep(time.Second)
	}
}

// tasklist - List of all processes on Windows
// taskkill /PID <pid> /F - Forcefully kill a process by PID on Windows
// Stop-Process - Id
