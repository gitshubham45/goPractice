package counterShard

import (
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	numShards    = 6
	maxQueueSize = 100
)

var IncrementChannel = make(chan bool, maxQueueSize)

// Shard is one part of the distributed counter
type Shard struct {
	value int64
}

// ShardedCounter groups multiple shards
type ShardedCounter struct {
	shards [numShards]Shard
}

// Increment adds 1 to a random shard
func (c *ShardedCounter) Increment() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	shardID := r.Intn(numShards)

	atomic.AddInt64(&c.shards[shardID].value, 1)
}

// Value returns the total value across all shards
func (c *ShardedCounter) Value() int64 {
	var total int64
	for i := 0; i < 4; i++ {
		total += atomic.LoadInt64(&c.shards[i].value)
	}
	return total
}

func (c *ShardedCounter) Reset() error {
	for i := 0 ; i < 4 ; i++ {
		c.shards[i].value = 0
	}

	return nil
}

// worker simulates concurrent increments
func Worker(counter *ShardedCounter, wg *sync.WaitGroup) {
	defer wg.Done()
	for range IncrementChannel {
		counter.Increment()
	}
}

func IncrementHandler(c *gin.Context) {
	select {
	case IncrementChannel <- true:
		c.JSON(http.StatusOK, gin.H{"message": "counter incremented"})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Too many requests - queue full"})
	}
}
