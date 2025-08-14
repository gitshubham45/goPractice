package main

import "fmt"

func main() {
	// Slice of integers
	nums := []int{1, 2, 3}

	// Case 1: Modifying the copy from range
	for _, v := range nums {
		if v == 2 {
			v = 99 // only changes the copy, not the slice
		}
	}
	fmt.Println("After case 1:", nums) // No change

	// Case 2: Modifying the slice element directly via index
	for i := range nums {
		if nums[i] == 2 {
			nums[i] = 99 // modifies the actual element
		}
	}
	fmt.Println("After case 2:", nums) // Change is reflected
}
