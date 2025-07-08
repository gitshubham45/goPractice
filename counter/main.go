package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Shared Counter Service
type Counter struct {
	Value int `json:"value"`
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.Value++
}

func (c *Counter) GetValue() int {
	c.Lock()
	defer c.Unlock()
	return c.Value
}

// Handlers
func incrementHandler(counter *Counter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		counter.Increment()
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status": "ok", "value": %d}`, counter.GetValue())
	}
}

func getHandler(counter *Counter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := counter.GetValue()
		json.NewEncoder(w).Encode(map[string]int{"value": val})
	}
}

// Client function to call increment endpoint
func nodeClient(nodeID int, wg *sync.WaitGroup) {
	defer wg.Done()
	url := "http://localhost:8080/increment"

	for i := 0; i < 200000; i++ {
		_, err := http.Get(url)
		if err != nil {
			log.Printf("Node %d: Error making request: %v\n", nodeID, err)
			continue
		}
		// body, _ := io.ReadAll(resp.Body)
		// log.Printf("Node %d: Response: %s\n", nodeID, body)
		// time.Sleep(200 * time.Millisecond) // Simulate random delay
	}
}

func main() {
	counter := &Counter{Value: 0}

	http.HandleFunc("/increment", incrementHandler(counter))
	http.HandleFunc("/get", getHandler(counter))

	// Run the server in a goroutine
	go func() {
		fmt.Println("Starting counter service on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Wait for server to start
	time.Sleep(time.Second)

	// Start multiple client goroutines simulating distributed nodes
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go nodeClient(i, &wg)
	}

	wg.Wait()

	// Print final value
	finalVal := counter.GetValue()
	fmt.Printf("\nFinal counter value: %d\n", finalVal)
}
