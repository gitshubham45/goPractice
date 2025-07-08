package main

import "fmt"

type Good struct {
	Name string
	Qty  int
}

type Order struct {
	ID    int
	Items []string
	Goods []Good
}

func main() {
	g1 := Good{Name : "Bat" , Qty : 5}
	
	o1 := Order{ID: 1, Items: []string{"book", "pen"} , Goods : []Good{g1}}


	o2 := o1


	// o2.Items[0] = "notebook"
	// o2.Items[1] = "llm"

	

	g2 := g1

	g2.Name = "Ball"

	// o1.Goods[0] = g1
	o1.Goods[0] = g2



	// fmt.Println(o1.Items) // Output: [notebook pen]

	fmt.Println(o1.Goods)
	fmt.Println(o2.Goods)

}
