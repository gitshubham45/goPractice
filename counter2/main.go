package main

import (
	"fmt"
	"sync"
)

// SharedCounter is a thread-safe counter
type SharedCounter struct {
	value int
	mu    sync.Mutex
}

// Increment increases the counter by 1
func (c *SharedCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	fmt.Println(c.value)
}

// Value returns the current counter value
func (c *SharedCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// worker simulates one incrementor from 1 to 50,000
func worker(c *SharedCounter, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 50000_000; i++ {
		c.Increment()
	}
	// fmt.Println("Worker done:", i)
}

func main() {
	var wg sync.WaitGroup
	counter := SharedCounter{}

	fmt.Println("Starting counter value:", counter.Value())

	wg.Add(2)
	go worker(&counter, &wg)
	go worker(&counter, &wg)

	wg.Wait()

	fmt.Println("Final counter value:", counter.Value())
}
