package main

import "fmt"

func main(){
	jobs := make(chan int , 50)
	results := make(chan int, 50)

	for w := 1 ; w <= 4 ; w++ {
		go worker(jobs,results)
	}

	for i := 0 ;i < 50 ; i++ {
		jobs <- i
	}
	close(jobs)

	for res := range results{
		fmt.Println(res)
	}
	
}

func worker(jobs <-chan int , results chan<- int ){
	for n := range jobs {
		results <- fib(n)
	}

	close(results)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}