package main

import "fmt"

// for any types
// func printItems[T any](items []T){
// 	for _ , item := range items{
// 		fmt.Println(item)
// 	}
// }

// for desired types
// func printItems[T int | string | bool](items []T){
// 	for _ , item := range items{
// 		fmt.Println(item)
// 	}
// }

// compareables

func printItems[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// multiple generic
func printMultipleItems[T comparable, V string](items []T, items2 []V) {
	for _, item := range items {
		fmt.Println(item)
	}

	for _, item := range items2 {
		fmt.Println(item)
	}
}

type Stack[T any] struct{
	elements []T
}

func main() {

	nums := []int{1, 2, 3, 4, 5}

	names := []string{"Ram", "Shyam", "Sam", "Peter"}

	// printItems(nums)
	// printItems(names)

	printMultipleItems(nums, names)

	myStack := Stack[int]{
		elements: []int{1,2,3,4},
	}

	myStringStack := Stack[string]{
		elements: []string{"hello" , "From" , "Go"},
	}

	fmt.Println(myStack)

	fmt.Println(myStringStack)
}
