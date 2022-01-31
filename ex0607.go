package main

import "fmt"

func main() {

	f1 := func() {
		fmt.Println("função atribuída a variável sem parâmetro")
	}

	f1()

	f2 := func(x int) {
		fmt.Println("função atribuída a variável com parâmetro", x)
	}

	f2(7)

}
