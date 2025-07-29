package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/sony/gobreaker"
)

func riskyService() (string, error) {
	if rand.Float32() < 0.5 {
		return "", errors.New("simulted failure")
	}
	return "success", nil
}

func main() {
	settings := gobreaker.Settings{
		Name:        "MyCircuitBreaker",
		MaxRequests: 3,                // Allowed requests in half-open state
		Interval:    time.Duration(0), // Reset internal counts every X duration ( 0 = never)
		Timeout:     5 * time.Second,  // Time to wait before trying after open
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// open the circuit when more than 3 failures in 5 attempts
			return counts.ConsecutiveFailures > 3
		},
	}

	cb := gobreaker.NewCircuitBreaker(settings)

	for i := 0; i < 10; i++ {
		result, err := cb.Execute(func() (interface{}, error) {
			return riskyService()
		})

		if err != nil {
			fmt.Println("Request failed:", err)
		} else {
			fmt.Println("Request success:", result)
		}
		time.Sleep(2 * time.Second)
	}
}
