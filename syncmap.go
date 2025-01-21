package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map

	// store values
	sm.Store("name", "Akkshay")
	sm.Store("age", 20)
	sm.Store("language", "Go")

	// load & print
	if name, ok := sm.Load("name"); ok {
		fmt.Printf("Name: %s\n", name)
	}
	if age, ok := sm.Load("age"); ok {
		fmt.Printf("Age: %d\n", age)
	}
	if language, ok := sm.Load("language"); ok {
		fmt.Printf("Language: %s\n", language)
	}

	// delete a value from the sync.Map
	sm.Delete("age")

	// range sync.Map
	fmt.Println("Contents of sync.Map:")
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %v\n", key, value)
		return true
	})

	// concurrent access to sync.Map
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Store(fmt.Sprintf("key%d", i), i)
		}(i)
	}
	wg.Wait()

	// final contents of sync.Map
	fmt.Println("Final contents of sync.Map:")
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %v\n", key, value)
		return true
	})
}
