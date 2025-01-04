package main

import (
	"errors"
	"fmt"
)

type DivisionError struct {
	dividend float64
	divisor  float64
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("cannot divide %.2f by zero", e.dividend)
}

// returns error
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, &DivisionError{dividend: dividend, divisor: divisor}
	}
	return dividend / divisor, nil
}

//  multiple error handling
func calculate(x, y float64) error {
	// for simple errors
	if x < 0 || y < 0 {
		return errors.New("negative numbers not allowed")
	}

	result, err := divide(x, y)
	if err != nil {
		return err
	}

	fmt.Printf("Result of %.2f / %.2f = %.2f\n", x, y, result)
	return nil
}

func main() {
	// Test cases
	testCases := []struct{ x, y float64 }{
		{10, 2},    // valid div
		{69, 0},    // div by zero
		{-5, 1},    // neg num
	}

	for _, tc := range testCases {
		fmt.Printf("\nCalculating %.2f / %.2f:\n", tc.x, tc.y)
		if err := calculate(tc.x, tc.y); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
