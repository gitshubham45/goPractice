package main

import "fmt"

func main() {
    name := "Shubham"
    age := 30
    pi := 3.14159
    flag := true

	// string and int
    fmt.Printf("Name: %s, Age: %d\n", name, age)
	// float upto 2 digit
    fmt.Printf("Pi approx: %.2f\n", pi)
    fmt.Printf("Boolean value: %t\n", flag)
    fmt.Printf("Hex of age: %x\n", age)
	// type of pi
    fmt.Printf("Type of pi: %T\n", pi)
    fmt.Printf("Pointer value: %p\n", &age)
    fmt.Printf("Show struct: %+v\n", struct{ Name string }{Name: name})
    fmt.Printf("Literal percent sign: %%\n")
}
