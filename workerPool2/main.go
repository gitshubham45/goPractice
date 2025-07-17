package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
	}
}

func main() {
	jobs := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	time.Sleep(2 * time.Second) // wait for workers to finish
}
