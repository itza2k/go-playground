package main

import (
	"fmt"
	"sync"
	"time"
)

// mutex usage
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock() //  defer - safety
	c.count++
}

func main() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	fmt.Println("Starting counter operations...")

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
				time.Sleep(time.Microsecond) // delay for concurrency
			}
		}()
	}

	var once sync.Once
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			once.Do(func() {
				fmt.Printf(" prints only once (goroutine %d)\n", n)
			})
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d (expected: 5000)\n", counter.count)
}
