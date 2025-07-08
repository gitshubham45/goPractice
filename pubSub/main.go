package main

import (
	"fmt"
	"sync"
	"time"
)

type Message any

func PublishMessages(publisherID int, messages chan<- Message, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		msg := fmt.Sprintf("Publisher %d: Message %d", publisherID, i+1)
		messages <- msg
		fmt.Printf("Published: %s\n", msg)
		time.Sleep(time.Millisecond * 150)
	}
}

func SubscribeToMessages(subscriberID int, messages <-chan Message) {
	for msg := range messages {
		fmt.Printf("Subscriber %d: Received '%v'\n", subscriberID, msg)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Printf("Subscriber %d: Channel closed. Exiting.\n", subscriberID)
}

func main() {
	var wg sync.WaitGroup
	messageChannel := make(chan Message, 5)

	go SubscribeToMessages(1, messageChannel)

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go PublishMessages(i+1, messageChannel, &wg)
	}

	wg.Wait()

	close(messageChannel)

	time.Sleep(time.Second)
	fmt.Println("Main: All done!")
}
