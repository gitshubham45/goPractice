package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	social  string `json:"social"`
	Ignored string `json:"-"`
}

func main() {
	person := Person{Name: "Shubham Kumar", Age: 23, Address: "123 Main St", social: "wwkfnasmf" , Ignored: "new"}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}

	// convet to string to make it readable
	fmt.Println(string(jsonData)) // Output : {"name":"Shubham Kumar","age":23,"address":"123 Main St"} // social ignored

	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println(newPerson)
	}
	fmt.Printf("%+v\n", newPerson) // Output: {Name:John Doe Age:30 Address:123 Main St}
}
