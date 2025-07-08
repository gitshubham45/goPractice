package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {
	now := time.Now()
	id := getUserByName("John")
	println(id)

	wg := &sync.WaitGroup{}
	ch := make(chan *Message )

	wg.Add(2)

	go getUserChats(id, ch , wg)
	go getUserFriends(id, ch , wg)

	wg.Wait()
	close(ch)


	// log.Println(chats)

	// log.Println(friends)

	// for msg := range ch {
	// 	log.Println(msg)
	// }

	select{
	case msg := <-ch:
		log.Println(msg)
	default : 
		log.Println("done")
	}

	



	log.Println(time.Since(now))
}

func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)
	defer wg.Done()

	ch <- &Message{
		friends: []string{
			"john",
			"Jane",
			"tiago",
			"Sam",
		},
	}


}

func getUserChats(id string, ch chan<- *Message ,wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	defer wg.Done()

	ch <- &Message{
		chats: []string{
			"John",
			"Jane",
			"Peter",
			"Ram",
		},
	}
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)

	return fmt.Sprintf("%s-2", name)
}
