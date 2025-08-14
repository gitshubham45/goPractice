package main

import (
	"fmt"
	"time"
)

func safeGoRoutoine(id int) {
	// panic recovery for this goroutine
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[Goroutine %d] recovered from panic %v\n", id, r)
		}
	}()

	if id == 2 {
		panic("something went wrong in goroutine 2")
	}

	fmt.Printf("[Goroutine %d] finished successfully\n", id)
}

func main() {
	for i := 1; i <= 3; i++ {
		go safeGoRoutoine(i)
	}

	time.Sleep(time.Second) // wait for goroutines
}
