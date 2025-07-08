package main

import (
    "fmt"
    "sync"
)

const (
    numShards = 4
    totalIterations = 50_000_000 // Each worker does this many increments
)

// Shard is one part of the distributed counter
type Shard struct {
    value int
    mu    sync.Mutex
}

// ShardedCounter groups multiple shards
type ShardedCounter struct {
    shards [numShards]Shard
}

// Increment adds 1 to a random shard
func (c *ShardedCounter) Increment() {
    shardID := hashFast(totalIterations) % numShards // deterministic or random
    c.shards[shardID].mu.Lock()
    c.shards[shardID].value++
    c.shards[shardID].mu.Unlock()
}

// hashFast is a fast pseudo-random-ish function to choose a shard
func hashFast(i int) int {
    return ((i * 74779648) >> 16) % numShards
}

// Value returns the total value across all shards
func (c *ShardedCounter) Value() int {
    total := 0
    for i := 0; i < numShards; i++ {
        c.shards[i].mu.Lock()
        total += c.shards[i].value
        c.shards[i].mu.Unlock()
    }
    return total
}

// worker simulates concurrent increments
func worker(c *ShardedCounter, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < totalIterations; i++ {
        c.Increment()
    }
}

func main() {
    var wg sync.WaitGroup
    counter := ShardedCounter{}

    fmt.Println("Starting counter value:", counter.Value())

    wg.Add(2)
    go worker(&counter, &wg)
    go worker(&counter, &wg)

    wg.Wait()

    fmt.Println("Final counter value:", counter.Value())
}