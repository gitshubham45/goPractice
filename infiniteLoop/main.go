package main

import (
	"fmt"
	"time"
)

func main() {
	// First goroutine: runs a ticker every 1 second
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for t := range ticker.C {
			go func(t time.Time) {
				fmt.Println("Goroutine started at:", t)
				time.Sleep(5 * time.Second) // simulate long task
				fmt.Println("Goroutine finished at:", time.Now())
			}(t)
		}
	}()

	// Second goroutine: simple loop every second
	go func() {
		for {
			fmt.Println("Running every second")
			time.Sleep(1 * time.Second)
		}
	}()

	// Block main from exiting
	select {}
}
