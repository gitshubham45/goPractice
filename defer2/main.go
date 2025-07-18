package main

import "fmt"

func main() {
	fmt.Println("Function start")

	defer fmt.Println("First defer: Executed last")
	defer fmt.Println("Second defer: Executed second")
	defer fmt.Println("Third defer: Executed first")

	fmt.Println("Function end")
}