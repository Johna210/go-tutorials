package main

import (
	"fmt"
	"sync"
	"time"
)

// useful for - SYNCHRONIZATION, COORDINATION, RESOURCE MANAGEMENT

// CONSTRUCTION EXAMPLE
type Worker struct {
	ID   int
	Task string
}

// PerformTask simulates a worker performing a task
func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerID %d finished %s\n", w.ID, w.Task)
}

func main_example_three() {
	var wg sync.WaitGroup

	// Define tasks to be performed by workers
	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{
			ID:   i + 1,
			Task: task,
		}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	// Construction is finished
	fmt.Println("Construction finished")
}

func main_with_wg_and_channels() {
	var wg sync.WaitGroup
	numWorkers := 3
	numJobs := 5
	results := make(chan int, numJobs)
	tasks := make(chan int, numJobs)

	wg.Add(numWorkers)

	for i := range numWorkers {
		go workerWithChannels(i+1, tasks, results, &wg)
	}

	for i := range numJobs {
		tasks <- i + 1
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}
}

// ==== example with channels
func workerWithChannels(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // simulate some time spent on processing the task
	for task := range tasks {
		results <- task * 2
	}
	fmt.Printf("Worker %d finished\n", id)
}

func main_With_just_waitGroup() {
	var wg sync.WaitGroup
	numWorkers := 3

	wg.Add(numWorkers)

	// Launch workers
	for i := range numWorkers {
		go worker_without_channels(i, &wg)
	}

	wg.Wait() // blocking mechanism
	fmt.Println("All workers finished")
}

// ===== basic example without using channels
func worker_without_channels(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // simulate some time spent on processing the task
	fmt.Printf("Worker %d finished\n", id)
}
