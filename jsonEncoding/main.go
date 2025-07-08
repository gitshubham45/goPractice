package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	data := []byte(`{"name": "Alice", "age": 30}`)
	var u1 User
	var u2 *User = &User{}

	json.Unmarshal(data, &u1) // Value semantic
	json.Unmarshal(data, u2)  // Pointer semantic

	fmt.Printf("u1: %+v\n", u1)
	fmt.Printf("u2: %+v\n", *u2)
}
