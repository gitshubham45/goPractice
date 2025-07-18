package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("test.txt")
	defer file.Close() // Ensures file is closed before main returns
	fmt.Fprintln(file, "Hello, defer!")
	fmt.Println("File written.")
}