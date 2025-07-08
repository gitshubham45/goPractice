package main

import (
	counterShard "counter/counterShard"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	var wg sync.WaitGroup
	counter := counterShard.ShardedCounter{}

	fmt.Println("Starting counter value:", counter.Value())

	for i := 0; i < 6; i++ {
		wg.Add(1)
		go counterShard.Worker(&counter, &wg)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	router.GET("/increment", counterShard.IncrementHandler)
	router.GET("/value", func(c *gin.Context) {
		value := counter.Value()

		c.JSON(http.StatusOK, gin.H{"currentValue": value})
	})

	router.GET("/reset" , func(c *gin.Context) {
		err := counter.Reset()
		if err != nil {
			c.JSON(http.StatusInternalServerError , gin.H{"error" : err})
		}
		c.JSON(http.StatusOK , gin.H{"message" : "counter reset successfull"})
	})

	go func() {
		router.Run(":8090")
	}()

	wg.Wait()

}
