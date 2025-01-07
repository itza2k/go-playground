package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// worker processing jobs with context
func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
	for {
		select {
		case j, ok := <-jobs:
			if (!ok) {
				return
			}
			fmt.Printf("Worker %d started job %d\n", id, j)
			select {
			case <-time.After(time.Second): // simulate work
				fmt.Printf("Worker %d finished job %d\n", id, j)
				results <- j * 2 // send result to channel
			case <-ctx.Done():
				fmt.Printf("Worker %d canceled job %d\n", id, j)
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

// sets up the jobs and results channels with context
func mainWorker() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// start 3 worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(ctx, w, jobs, results)
	}

	// 5 jobs to jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// collecting results channel
	for a := 1; a <= numJobs; a++ {
		select {
		case result := <-results:
			log.Printf("Result received: %d\n", result)
		case <-ctx.Done():
			log.Println("Main function: context timeout or cancellation")
			return
		}
	}
}

// error handling
func simulateErrorHandling() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from error: %v\n", r)
		}
	}()

	panic("simulated error")
}

// main func starts mainWorker func
func main() {
	fmt.Println("Starting mainWorker")
	mainWorker()

	fmt.Println("Simulating error handling")
	simulateErrorHandling()
}
