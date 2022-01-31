package main

import "fmt"

type person struct {
	firstName, lastName string
	iceCreamFlavours    []string
}

func main() {
	p1 := person{
		firstName:        "Wagner",
		lastName:         "Galdino",
		iceCreamFlavours: []string{"peanut", "corn"},
	}
	fmt.Println(p1.firstName, p1.lastName)
	for _, v := range p1.iceCreamFlavours {
		fmt.Println(v)
	}
}
