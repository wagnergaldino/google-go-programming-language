package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func changeMe(p *person) {
	p.firstName = "Wagner"
	p.lastName = "Galdino"
	p.age = 55
}

func main() {
	a := person{
		firstName: "WAGNER",
		lastName:  "GALDINO",
		age:       55,
	}
	fmt.Println(a)
	changeMe(&a)
	fmt.Println(a)

}
