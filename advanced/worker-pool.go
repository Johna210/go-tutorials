package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personID   int
	numTickets int
	cost       int
}

func main_example_two() {
	numRequests := 5
	price := 5
	ticketRequests := make(chan ticketRequest, numRequests)
	ticketResults := make(chan int)

	// start the ticket processor(worker)
	for range 3 {
		go ticketProcessor(ticketRequests, ticketResults)
	}

	// send ticket requests
	for i := range numRequests {
		ticketRequests <- ticketRequest{
			personID:   i + 1,
			numTickets: (i + 1) * 2,
			cost:       (i + 1) * price,
		}
	}
	close(ticketRequests)

	for range numRequests {
		fmt.Printf("Ticket for personID %d processed successfully!\n", <-ticketResults)
	}
}

// simulate processing of ticket requests
func ticketProcessor(requests <-chan ticketRequest, results chan<- int) {
	for req := range requests {
		fmt.Printf("Processing %d ticket(s) of personID %d with total cost %d\n",
			req.numTickets, req.personID, req.cost)
		// simulate processing time
		time.Sleep(time.Second)
		results <- req.personID
	}
}

// ==== Basic worker pool pattern
func main_example_one() {
	numWorkers := 3
	numJobs := 10

	tasks := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create worker
	for i := range numWorkers {
		go worker(i, tasks, results)
	}

	// Send values to the tasks channel
	for i := range numJobs {
		tasks <- i
	}

	close(tasks)

	// Collect the results
	for range numJobs {
		result := <-results
		fmt.Println("Result:", result)
	}
}

func worker(id int, tasks <-chan int, results chan<- int) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		// Simulate some work
		time.Sleep(time.Second)
		results <- task * 2
	}
}
