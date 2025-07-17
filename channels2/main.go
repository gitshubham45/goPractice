package main

import "fmt"

func task(ch chan int){
	for i := 1 ;i <= 5 ;i ++{
		ch <- i
	}
	close(ch) // closing the channel to prevent deadlock
}

func main(){
	ch := make(chan int,5) // integer channel of size 5

	go task(ch)

	// loopin over channle
	for val := range ch {
		fmt.Println(val)
	}
}