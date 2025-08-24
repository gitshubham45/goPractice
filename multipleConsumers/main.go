package main

import "fmt"

func produce(n int, data chan<- int) {
	for i := 1; i <= n; i++ {
		data <- i
	}
	close(data)
}

func consume(data <-chan int, done chan<- bool, chanID int) {
	for val := range data {
		fmt.Printf("value : %d Printed from chan : %d \n", val, chanID)
	}
	done<-true
}

func main() {

	data := make(chan int,5)
	done := make(chan bool)

	n := 100005

	go produce(n,data)

	for i := 1 ;i <=3 ;i++{
		go consume(data,done,i)
	}

	// for i := 1 ;i <=3 ;i++{
		
	// }
	<-done
}
