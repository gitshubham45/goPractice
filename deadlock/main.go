package main

import "fmt"

func publish(deadCh chan<- int ){
	for i:= 1; i<= 10 ;i++{
		deadCh <- i
	}
	close(deadCh)
}

func consume(deadCh <-chan int, done chan<- bool){
	for val := range deadCh{
		fmt.Println(val)
	}
	done <- true
}

func main(){
	
	deadCh := make(chan int)
	done := make(chan bool)

	go publish(deadCh)
	go consume(deadCh,done)

	<-done


}