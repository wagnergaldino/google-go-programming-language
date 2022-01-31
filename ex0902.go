package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("person speak")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "Wagner",
		age:  55,
	}
	saySomething(&p)

	//se não passar ponteiro pra p não funciona
	//saySomething(p)

	p.speak()
}
