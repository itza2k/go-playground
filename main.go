package main

import (
	"fmt"
	"time"
)

func main() {
	// basic var stuff
	name := "Akkshay"
	age := 20
	language := "Go"
	isLearning := true

	// printing w formatting
	fmt.Println("=== My First Go Program ===")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Learning: %s\n", language)
	fmt.Printf("Currently Learning: %v\n\n", isLearning)

	// concurrent function call
	go greetConcurrently("Hello from a goroutine!")

	// thread work - main
	fmt.Println("Main function is doing other tasks...")

	// channels
	resultChan := make(chan int)
	go sumNumbers(3, 5, resultChan)

	// goroutine to complete
	time.Sleep(1 * time.Second)
	fmt.Println("Program is ending.")

	// channel result and print
	sumResult := <-resultChan
	fmt.Printf("Received sum from channel: %d\n", sumResult)
}

// sleep and print
func greetConcurrently(message string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println(message)
}

// sleep and sum to channel
func sumNumbers(a, b int, c chan int) {
	time.Sleep(200 * time.Millisecond)
	c <- a + b
}
