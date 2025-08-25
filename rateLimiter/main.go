package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 100)

	go func() {
		i := 1
		for {
			requests <- i
			fmt.Println("Produced request:", i)
			i++
			time.Sleep(200 * time.Millisecond)
		}
	}()

	limiter := time.NewTicker(1 * time.Second)
	defer limiter.Stop()

	go func() {
		for req := range requests {
			select {
			case <-limiter.C:
				fmt.Println("Processed request : ", req, time.Now())
			default:
				fmt.Println("Dropped request : ", req)
			}
		}
	}()

	time.Sleep(10 * time.Second)
}
