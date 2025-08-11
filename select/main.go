package main

import (
	"fmt"
	"time"
)

func main(){
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Println(msg)
	default :
		fmt.Println("No message recieved")
	}

	go func(){
		time.Sleep(2*time.Second)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(1*time.Second):
		fmt.Println("timeout")
	}

}