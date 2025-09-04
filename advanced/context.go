package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// USED FOR - cancellation, timeouts, values
//

func context_main() {
	rootCtx := context.Background()
	// ctx, cancel := context.WithTimeout(rootCtx, 2*time.Second)
	// defer cancel()

	ctx, cancel := context.WithCancel(rootCtx)
	go func() {
		time.Sleep(2 * time.Second) // simulating a heavy task.
		cancel()                    // manually canceling after the task is finished.
	}()

	ctx = context.WithValue(ctx, "requestID", "dhasddsah456")

	go doWork2(ctx)

	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("No request ID found.")
	}

	logWithContext(ctx, "This is a test log message")
}

func logWithContext(ctx context.Context, message string) {
	requestIDVal := ctx.Value("requestID")
	log.Printf("RequestID: %v - %v", requestIDVal, message)
}

func doWork2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work Cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func simple_timeout_example() {
	ctx := context.TODO()

	result := checkEvenOdd(ctx, 5)
	fmt.Println("Result with context.TODO():", result)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result = checkEvenOdd(ctx, 10)
	fmt.Println("Result from timeout context:", result)

	time.Sleep(2 * time.Second)
	result = checkEvenOdd(ctx, 15)
	fmt.Println("Result from timeout context:", result)
}

func checkEvenOdd(ctx context.Context, num int) string {
	select {
	case <-ctx.Done():
		return "Operation cancelled"
	default:
		if num%2 == 0 {
			return fmt.Sprintf("%d is even", num)
		} else {
			return fmt.Sprintf("%d is odd", num)
		}
	}
}

// ==== DIFFERENCE BETWEEN CONTEXT.TODO AND CONTEXT.BACKGROUND
func different_context() {
	// mainly used as a placeholder (by go convention) - TODO
	todoContext := context.TODO()
	contextBkg := context.Background()

	ctx := context.WithValue(todoContext, "name", "John")
	fmt.Println(ctx)
	fmt.Println(ctx.Value("name"))

	ctx2 := context.WithValue(contextBkg, "age", "30")
	fmt.Println(ctx2)
	fmt.Println(ctx2.Value("age"))
}
