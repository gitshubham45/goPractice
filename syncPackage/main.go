package main

import (
	"fmt"
	"sync"
)

var balance int
var mutex sync.Mutex // mutex from sync package

func deposit(amount int, wg *sync.WaitGroup){
	mutex.Lock()
	defer wg.Done() // signals done
	defer mutex.Unlock()
	balance += amount
	fmt.Printf("Deposited %d. New Balance : %d\n" , amount , balance)
}

func withdraw(amount int , wg *sync.WaitGroup){
	mutex.Lock()
	defer wg.Done() //
	defer mutex.Unlock()
	balance -= amount
	fmt.Printf("Withdrawn %d. New Balance : %d\n" , amount , balance)
}

func main(){

	var wg sync.WaitGroup // waitGrou from sync package

	balance = 1000

	wg.Add(3)
	go deposit(200,&wg)
	go withdraw(300,&wg)
	go deposit(500,&wg)

	wg.Wait() // waiting for go routines to complete
}


// The sync package in Go helps manage parts of your program that run concurrently. When multiple goroutines access shared resources, there's a risk of race conditions or inconsistent behavior.

// To prevent this, sync provides tools like mutexes, wait groups, and condition variables. These help goroutines coordinate properly, ensuring safe and efficient execution in multi-threaded environments.