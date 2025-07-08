package main

import (
	"fmt"
)

func task(done chan bool) {
	defer func() { done <- true }()

	fmt.Println("processing...")
}

// unbuffered channels are blocking unlike buffered

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		// time.Sleep(time.Second * 1)
		fmt.Println("Sending emai to ", email)
	}
}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)
	done := make(chan bool)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	go func() {
		done <- true
	}()

	for {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Value recieved form chan1 ", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Value recieved from chan2 ", chan2Val)
		}

	}

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// // emailChan <- "abc@example.com"
	// // emailChan <- "pqr@example.com"

	// go emailSender(emailChan, done)

	// for i := 0; i < 100; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("done sending.")
	// close(emailChan)
	// <-done

	// done := make(chan bool)
	// go task(done)

	// <-done
}
