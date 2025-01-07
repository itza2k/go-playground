package main

import (
	"fmt"
)

// generic function to find the max val
func Max[T comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// generic function to print a slice
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	// using Max with int
	fmt.Printf("Max of 3 and 5: %d\n", Max(3, 5))

	// using Max with float64
	fmt.Printf("Max of 2.5 and 3.7: %.1f\n", Max(2.5, 3.7))

	// using PrintSlice with int
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Int slice:")
	PrintSlice(intSlice)

	// using PrintSlice with string
	stringSlice := []string{"Go", "is", "awesome"}
	fmt.Println("String slice:")
	PrintSlice(stringSlice)
}
