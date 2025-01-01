package main

import (
	"fmt"
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
	fmt.Printf("Currently Learning: %v\n", isLearning)
}
