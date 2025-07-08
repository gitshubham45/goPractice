package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "Every 500 ms from ch 1"
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for {
			ch2 <- "Every 2 second from ch 2"
			time.Sleep(2000 * time.Millisecond)
		}
	}()

	for {
		select{
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

}
