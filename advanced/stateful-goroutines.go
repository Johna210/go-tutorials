package main

import "time"

type StatefulWorker struct {
	count int
	ch    chan int
}

func (sw *StatefulWorker) Start() {
	go func() {
		for {
			select {
			case value := <-sw.ch:
				sw.count += value
				println("Current count:", sw.count)
			}
		}
	}()
}

func (sw *StatefulWorker) Send(value int) {
	sw.ch <- value
}

func main_stateful_goroutines() {
	worker := &StatefulWorker{
		ch: make(chan int),
	}
	worker.Start()

	for i := range 5 {
		worker.Send(i)
		time.Sleep(500 * time.Millisecond)
	}
}
