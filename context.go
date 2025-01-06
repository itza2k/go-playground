package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// context timeout 2 sec
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// goroutine long-running task
	go func() {
		if err := longRunningTask(ctx); err != nil {
			fmt.Printf("Task error: %v\n", err)
		} else {
			fmt.Println("Task completed successfully")
		}
	}()

	// wait for timeout
	select {
	case <-ctx.Done():
		fmt.Println("Main function: context timeout or cancellation")
	case <-time.After(3 * time.Second):
		fmt.Println("Main function: task completed")
	}
}

func longRunningTask(ctx context.Context) error {
	// 7 sec task
	select {
	case <-time.After(7 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
