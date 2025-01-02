package main

import (
	"fmt"
	"time"
)

// worker processing jobs
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // time taken
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2 // send result to channel
	}
}

// sets up the jobs and results channels
func mainWorker() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start 3 worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//  5 jobs to jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// collecting results channel
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

// main func starts mainWorker func
func main() {
	mainWorker()
}
