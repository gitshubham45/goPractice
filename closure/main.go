package main

import "fmt"

// counter returns a function that increments an int and returns its new value each time it's called.
func counter() func() int {
	i := 0 // This variable 'i' is in the outer scope of the returned function.
	return func() int { // This is the inner function (the closure).
		i++ // It 'closes over' and modifies 'i'.
		return i
	}
}

func main() {
	// Create the first counter. 'c1' is a closure that has its own 'i'.
	c1 := counter()

	fmt.Println("c1 first call:", c1()) // Output: c1 first call: 1
	fmt.Println("c1 second call:", c1()) // Output: c1 second call: 2

	// Create a second counter. 'c2' is a *new* closure with its *own* separate 'i'.
	c2 := counter()
	fmt.Println("c2 first call:", c2()) // Output: c2 first call: 1
	fmt.Println("c1 third call:", c1()) // Output: c1 third call: 3 (c1's 'i' is separate)
}