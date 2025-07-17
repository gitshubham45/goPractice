package main

import (
	"fmt"
	"time"
)

func fetchData(source string) {
	time.Sleep(2 * time.Second)
	fmt.Println("Data fetched from", source) 
}

func main() {
	go fetchData("API 1") // Data fetched from API 2 - after 2 sec
	go fetchData("API 2") // Data fetched from API 1 - after 2 sec

	time.Sleep(3 * time.Second)
	fmt.Println("Done") // Done - after 3 sec
}
