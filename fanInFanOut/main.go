package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker processes jobs and send results
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n" ,id , j )
		time.Sleep(300 * time.Millisecond) // simulate processing
		results <- j * 3
	}
}

func main() {
	const numJobs = 6
	const numWorkers = 2

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Fan-out : Start Workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// send jobs
	for j := 1; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// wait for all workers to fisnish'
	wg.Wait()
	close(results)

	// Fan-in : Collect results
	for res := range results {
		fmt.Println("Result:", res)
	}
}
