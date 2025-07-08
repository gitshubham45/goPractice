package main

import "fmt"

func main(){

	// var mp map[interface{}]interface{}

	// var map2 map[string]int


	map2 := map[string]int{
		"ram" :1,
	}

	map3 := make(map[string]int )

	map3["nil"] = 5

	fmt.Println(map2)
	fmt.Println(map3)

	// mp[2] = 3

	// mp["ram"] = "shyam"

	map2["ram"] = 2

	// fmt.Println(mp)
	fmt.Println(map2)


}