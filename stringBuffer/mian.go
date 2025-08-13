package main

import (
	"bytes"
	"fmt"
)

func main() {
	// String → Buffer
	str := "hello"
	buf := bytes.NewBufferString(str)
	fmt.Println("Buffer from string:", buf) // prints the buffer contents

	// Buffer → String
	strFromBuf := buf.String()
	fmt.Println("String from buffer:", strFromBuf)

	// String → Byte slice
	byteSlice := []byte(str)
	fmt.Println("Byte slice:", byteSlice)

	// Byte slice → String
	strFromBytes := string(byteSlice)
	fmt.Println("String from bytes:", strFromBytes)

	// Create buffer from byte slice
	bufFromBytes := bytes.NewBuffer(byteSlice)
	fmt.Println("Buffer from bytes:", bufFromBytes)
}
