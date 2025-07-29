package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	// var name string
	// fmt.Println("Enter your name:")
	// fmt.Scanln(&name) // reads input untill a newline (space breaks input)
	// fmt.Println("Hello, " , name)

	// multiple inputs
	var name string 
	var age int

	fmt.Println("Enter your name and age : ")
	fmt.Scanln(&name,&age)

	fmt.Printf("Name: %s , Age: %d\n", name , age)
		// 	Enter your name and age : 
		// shubham 23
		// Name: shubham , Age: 23

	fmt.Println("Enter a sentence : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // reads input including space
	fmt.Println("You entered : ", input)

	// output
	// Enter a sentence : 
	// Hi Sam How are you
	// You entered :  Hi Sam How are you
}