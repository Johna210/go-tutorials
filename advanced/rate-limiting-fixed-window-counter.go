package main

import (
	"sync"
	"time"
)

type RateLimiterFixedWindow struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewRateLimiterFixedWindow(limit int, window time.Duration) *RateLimiterFixedWindow {
	return &RateLimiterFixedWindow{
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiterFixedWindow) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func main_for_fixed_window() {
	var wg sync.WaitGroup
	rateLimiter := NewRateLimiterFixedWindow(5, 1*time.Second)

	for range 10 {
		wg.Add(1)
		go func() {
			if rateLimiter.Allow() {
				println("Request allowed")
			} else {
				println("Request denied")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
