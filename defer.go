package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Yo, itza2k exploring 'defer'!")
	fmt.Println("=============================================")
	fmt.Println("Created by Akkshay - Build & Go")

	fmt.Println("\n1. Basic defer usage:")
	basicDefer()

	fmt.Println("\n2. Multiple defers (LIFO order):")
	multipleDefers()

	fmt.Println("\n3. Argument evaluation in defer:")
	deferArgumentEval()

	fmt.Println("\n4. Practical example - resource cleanup:")
	resourceCleanup()

	fmt.Println("\n5. Akkshay's builder :")
	akkshaysBuildProject()
}

func basicDefer() {
	// defer pushes a function call onto a stack
	defer fmt.Println("  This line prints at the end (deferred)")

	fmt.Println("  This line prints first")
	fmt.Println("  This line prints second")

}

func multipleDefers() {
	// executed in LIFO order 
	defer fmt.Println("  First defer (prints last)")
	defer fmt.Println("  Second defer (prints second)")
	defer fmt.Println("  Third defer (prints first)")

	fmt.Println("  Regular code executes before any defers")
}

func deferArgumentEval() {
	x := 5
	defer fmt.Printf("  Deferred print: x = %d\n", x)

	x = 10
	fmt.Printf("  Regular print: x = %d\n", x)
}

func resourceCleanup() {

	fmt.Println("  Opening resource...")
	defer func() {
		fmt.Println("  Closing resource... (deferred)")
	}()

	fmt.Println("  Working with resource...")
	fmt.Println("  Done with resource")

}

func akkshaysBuildProject() {

	fmt.Println("  itza2k is starting a new build project!")
	startTime := time.Now()

	// Track how long the project takes using defer
	defer func() {
		duration := time.Since(startTime)
		fmt.Printf("  this project took %v to complete\n", duration)
		fmt.Println("  ✅ Build successful")
	}()

	defer fmt.Println("  🧹 Cleaning up the workspace...")
	defer fmt.Println("  📦 Putting away unused materials...")
	defer fmt.Println("  📝 Saving final project documentation...")

	// steps
	fmt.Println("  📐 Planning the design...")
	time.Sleep(300 * time.Millisecond)

	fmt.Println("  🔨 Building the foundation...")
	time.Sleep(500 * time.Millisecond)

	fmt.Println("  🏗️ Constructing the main structure...")
	time.Sleep(700 * time.Millisecond)

	fmt.Println("  🎨 Adding finishing touches...")
	time.Sleep(200 * time.Millisecond)

	fmt.Println("  🔍 Inspecting the final product...")
	time.Sleep(100 * time.Millisecond)

	// all deferred functions execute in LIFO order
}
