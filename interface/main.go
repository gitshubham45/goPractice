package main

import "fmt"

// Define an interface
type Greeter interface {
    Greet() string
}

// Define a struct
type Person struct {
    Name string
}

// Person implicitly implements Greeter because it has a Greet() method
func (p Person) Greet() string {
    return "Hello, my name is " + p.Name
}

// Another struct
type Robot struct {
    Model string
}

// Robot also implicitly implements Greeter
func (r Robot) Greet() string {
    return "BEEP BOOP, I am model " + r.Model
}
 
func sayHello(g Greeter) {
    fmt.Println(g.Greet())
}

type Phone struct{
	Type string
}

func (s *Phone) Greet() string{
	return "Hello from phone type " + s.Type
}

func main() {
    p := Person{Name: "Alice"}
    r := Robot{Model: "RX-800"}

	s := Phone{Type : "Android"}

    sayHello(p) // Works because Person implements Greeter
    sayHello(r) // Works because Robot implements Greeter
	sayHello(&s)
}