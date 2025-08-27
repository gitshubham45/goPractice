package main

import "fmt"

func getNums() <-chan int {
	nums := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			nums <- i
		}
		close(nums)
	}()
	return nums
}

func getSquares(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range in {
			out <- val * val
		}
		close(out)
	}()
	return out
}

func getDoubles(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range in {
			out <- 2 * val
		}
		close(out)
	}()
	return out
}

func main() {
	nums := getNums()
	squares := getSquares(nums)
	doubles := getDoubles(squares)

	for val := range doubles {
		fmt.Println(val)
	}
}
