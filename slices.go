package main

import (
	"fmt"
)

func exploreSlices() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original slice: %v\n", numbers)

	// append 
	numbers = append(numbers, 6, 7)
	fmt.Printf("After appending: %v\n", numbers)

	// slice the slice
	subSlice := numbers[2:5]
	fmt.Printf("Sub-slice: %v\n", subSlice)

	numbers[0] = 10
	fmt.Printf("After modifying: %v\n", numbers)

	// iterate 
	fmt.Println("Iterating over slice:")
	for i, num := range numbers {
		fmt.Printf("Index %d: %d\n", i, num)
	}

	// make slice 
	moreNumbers := make([]int, 5, 10)
	fmt.Printf("Slice created with make: %v (length: %d, capacity: %d)\n", moreNumbers, len(moreNumbers), cap(moreNumbers))

	// append more
	moreNumbers = append(moreNumbers, 1, 2, 3, 4, 5, 6)
	fmt.Printf("After appending beyond capacity: %v (length: %d, capacity: %d)\n", moreNumbers, len(moreNumbers), cap(moreNumbers))
}


func main() {
	fmt.Println("=== Exploring Slices ===")
	exploreSlices()
}
