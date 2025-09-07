package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwmu     sync.RWMutex
	counter2 int
)

func main_rw() {
	var wg sync.WaitGroup

	for range 5 {
		wg.Add(1)
		go readCounter(&wg)
	}
	wg.Add(1)
	time.Sleep(time.Second)
	go writeCounter(&wg, 18)

	wg.Wait()
}

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	rwmu.RLock()
	fmt.Println("Read counter:", counter2)
	rwmu.RUnlock()
}

func writeCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	rwmu.Lock()
	counter2 = value
	fmt.Println("Wrote counter:", counter2)
	rwmu.Unlock()
}
