package main

import "fmt"

func process(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		processed := job * 10
		results <- processed
	}
}

func main() {
	jobs := make(chan int, 3)
	results := make(chan int, 3)

	go process(jobs, results)

	jobs <- 1
	jobs <- 2
	jobs <- 3
	close(jobs)

	for i := 0; i < 3; i++ {
		fmt.Println("Result:", <-results)
	}
}
