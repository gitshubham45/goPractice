package main

import "fmt"

func main() {
	a := 5
	b := float64(a) // b is now 5.0
	fmt.Println(b)  // 5.0

	f := 3.14
	i := int(f)    // i is 3
	fmt.Println(i) // 3

	s := "hello"
	d := []byte(s) // d = [104 101 108 108 111]
	fmt.Println(d) // [104 101 108 108 111]

	t := []byte{72, 73}
	r := string(t) // r = "HI"
	fmt.Println(r) // "HI"
}
